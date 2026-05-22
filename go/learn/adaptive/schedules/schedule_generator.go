/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * Schedule generator — async goroutine that:
 *   1. Calls LLM to generate weekly schedule blocks
 *   2. Saves schedule with blocks
 *   3. For each academic block, calls LLM to generate a lesson (serial loop)
 */
package schedules

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/learn/adaptive/engine"
	"github.com/saichler/l8learn/go/learn/content/genlessons"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

type scheduleHandler struct {
	llmClient engine.LLMClient
	masker    *engine.PIIMasker
}

var sHandler = &scheduleHandler{}

func SetLLMClient(client engine.LLMClient, masker *engine.PIIMasker) {
	sHandler.llmClient = client
	sHandler.masker = masker
	fmt.Printf("[Schedule] LLM client injected: %v\n", client != nil)
}

func (h *scheduleHandler) generateSchedule(sched *learn.DailySchedule, vnic ifs.IVNic) {
	log := func(msg string, args ...interface{}) {
		fmt.Printf("[ScheduleGen %s] %s\n", sched.ScheduleId, fmt.Sprintf(msg, args...))
	}

	log("Starting schedule generation")

	// 1. Load student profile (find student via family or use first student)
	// For now, we need a studentId — store it in schedule's custom_fields
	studentId := ""
	if sched.CustomFields != nil {
		studentId = sched.CustomFields["studentId"]
	}
	if studentId == "" {
		log("FAILED: No studentId in schedule custom_fields")
		return
	}

	query := "select * from StudentProfile where studentId=" + studentId
	profiles, err := common.GetEntitiesByQuery("Profile", byte(20), query, vnic)
	if err != nil || len(profiles) == 0 {
		log("FAILED: Could not load profile for student %s: %v", studentId, err)
		return
	}
	studentProfile := profiles[0].(*learn.StudentProfile)

	// 2. Check profile completeness
	ready, msg := CheckProfileReady(studentProfile)
	if !ready {
		log("FAILED: %s", msg)
		return
	}

	// 3. Mask PII in profile before sending to LLM
	knownNames := buildScheduleMaskingContext(studentId, vnic)
	log("Masking context: %d known names: %v", len(knownNames), knownNames)
	tokenMap := engine.NewTokenMap()

	profileJSON, _ := json.Marshal(studentProfile)
	maskedProfile := h.masker.MaskTextWithNames(string(profileJSON), tokenMap, knownNames)
	log("Profile masked: %d chars (original %d)", len(maskedProfile), len(profileJSON))

	programJSON := "{}"
	if studentProfile.ProgramSettings != nil {
		programJSON2, _ := json.Marshal(studentProfile.ProgramSettings)
		programJSON = string(programJSON2)
	}
	constraints := fmt.Sprintf("weather=%s, parentEnergy=%s, appointments=%v",
		sched.Weather, sched.ParentEnergy, sched.Appointments)

	// Token limit check
	totalChars := len(profileJSON) + len(programJSON) + len(constraints)
	log("Total prompt size: ~%d chars (~%d tokens)", totalChars, totalChars/4)
	if totalChars > 400000 {
		log("FAILED: Prompt too large: %d chars", totalChars)
		return
	}

	systemPrompt, userMessage := engine.BuildSchedulePrompt(
		maskedProfile, programJSON, constraints)

	// 4. Call LLM for schedule
	log("Calling LLM for schedule blocks...")
	response, err := h.llmClient.Call(
		learn.LLMPromptType_LLM_PROMPT_TYPE_SCHEDULE_GENERATION,
		systemPrompt, userMessage, studentId)
	if err != nil {
		log("FAILED: LLM call: %s", err.Error())
		return
	}
	log("LLM response: %d chars", len(response))
	// Unmask PII in response
	response = tokenMap.Unmask(response)
	log("--- SCHEDULE RESPONSE (first 3000 chars) ---\n%s\n--- END ---", truncate(response, 3000))

	// 5. Parse schedule response
	blocks, academicCount := parseScheduleResponse(response)
	if len(blocks) == 0 {
		log("FAILED: No blocks parsed from LLM response")
		return
	}
	log("Parsed %d blocks (%d academic)", len(blocks), academicCount)

	// 6. Save schedule with blocks
	sched.Blocks = blocks
	sched.LessonsTotal = academicCount
	sched.LessonsGenerated = 0
	common.PutEntity(ServiceName, ServiceArea, sched, vnic)
	log("Schedule saved with %d blocks", len(blocks))

	// 7. Serial lesson generation for each academic block (l8agent tool-loop pattern)
	lessonNum := int32(0)
	for _, block := range blocks {
		if block.ActivityType != "academic" && block.ActivityType != "therapy" {
			continue
		}

		lessonNum++
		log("Generating lesson %d/%d for block %s (%s)...",
			lessonNum, academicCount, block.BlockId, block.Description)

		blockJSON, _ := json.Marshal(block)
		sysPrompt, usrMsg := engine.BuildLessonFromSchedulePrompt(
			maskedProfile, string(blockJSON), "{}")

		// Token limit per lesson
		if len(sysPrompt)+len(usrMsg) > 400000 {
			log("WARNING: Lesson prompt too large, skipping block %s", block.BlockId)
			continue
		}

		lessonResp, err := h.llmClient.Call(
			learn.LLMPromptType_LLM_PROMPT_TYPE_GENERATE_LESSON,
			sysPrompt, usrMsg, studentId)
		if err != nil {
			log("WARNING: Lesson generation failed for block %s: %s", block.BlockId, err.Error())
			continue
		}
		// Unmask PII in lesson response
		lessonResp = tokenMap.Unmask(lessonResp)

		lesson, parseErr := genlessons.ParseLessonResponse(lessonResp)
		if parseErr != nil {
			log("WARNING: Lesson parse failed for block %s: %s", block.BlockId, parseErr.Error())
			continue
		}

		// Set metadata and POST
		common.GenerateID(&lesson.GeneratedLessonId)
		lesson.StudentId = studentId
		lesson.ScheduleId = sched.ScheduleId
		lesson.BlockId = block.BlockId
		lesson.Status = learn.GeneratedLessonStatus_GENERATED_LESSON_STATUS_READY
		lesson.GeneratedAt = time.Now().Unix()

		_, postErr := common.PostEntity("GenLesson", byte(10), lesson, vnic)
		if postErr != nil {
			log("WARNING: Failed to save lesson for block %s: %s", block.BlockId, postErr.Error())
			continue
		}

		// Update progress
		sched.LessonsGenerated = lessonNum
		common.PutEntity(ServiceName, ServiceArea, sched, vnic)
		log("Lesson %d/%d saved: %s", lessonNum, academicCount, lesson.Title)
	}

	log("Schedule generation COMPLETE — %d/%d lessons generated", lessonNum, academicCount)
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "... [TRUNCATED]"
}

// buildScheduleMaskingContext loads student + family names for PII masking.
func buildScheduleMaskingContext(studentId string, vnic ifs.IVNic) []string {
	var names []string
	result, err := common.GetEntity("Student", byte(20),
		&learn.Student{StudentId: studentId}, vnic)
	if err == nil && result != nil {
		s := result.(*learn.Student)
		addName(&names, s.FirstName)
		addName(&names, s.LastName)
		addName(&names, s.PreferredName)
	}
	// Load guardians
	guardians, gErr := common.GetEntities("Guardian", byte(20),
		&learn.Guardian{StudentIds: []string{studentId}}, vnic)
	if gErr == nil {
		for _, g := range guardians {
			guardian := g.(*learn.Guardian)
			addName(&names, guardian.FirstName)
			addName(&names, guardian.LastName)
		}
	}
	return names
}

func addName(names *[]string, name string) {
	if name != "" && len(name) > 1 {
		*names = append(*names, name)
	}
}
