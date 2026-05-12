/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * ParentCoach generates daily coaching tips for guardians.
 * Triggered by daily scheduler — one tip per active family.
 * Uses PARENT_COACHING prompt type. Logs to PromptLog.
 */
package engine

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

// CoachingTip is the parsed response from the LLM
type CoachingTip struct {
	Tip                string `json:"tip"`
	ActivitySuggestion string `json:"activitySuggestion"`
	Materials          string `json:"materials"`
}

// ParentCoach generates and stores daily coaching tips
type ParentCoach struct {
	vnic      ifs.IVNic
	llmClient LLMClient
}

func NewParentCoach(vnic ifs.IVNic, llmClient LLMClient) *ParentCoach {
	return &ParentCoach{vnic: vnic, llmClient: llmClient}
}

// GenerateDailyTip creates a coaching tip for one family
// Called by the daily scheduler for each active family
func (pc *ParentCoach) GenerateDailyTip(
	familyId string,
	studentIds []string,
	profiles []*learn.StudentProfile,
	paths []*learn.LearningPath,
) (*CoachingTip, error) {

	// Build context from children's profiles and current paths
	currentSkills := pc.summarizeCurrentSkills(paths)
	schedule := "today's planned activities"
	learningStyle := pc.summarizeLearningStyles(profiles)
	interests := pc.summarizeInterests(profiles)

	// Build and send prompt
	systemPrompt, userMessage := BuildParentCoachingPrompt(
		currentSkills, schedule, learningStyle, interests,
	)

	// Pick first student for the prompt log (family-level, but logged per student)
	studentId := ""
	if len(studentIds) > 0 {
		studentId = studentIds[0]
	}

	response, err := pc.llmClient.Call(
		learn.LLMPromptType_LLM_PROMPT_TYPE_PARENT_COACHING,
		systemPrompt, userMessage, studentId,
	)
	if err != nil {
		return nil, fmt.Errorf("coaching prompt failed: %w", err)
	}

	// Parse response
	tip := &CoachingTip{}
	if err := json.Unmarshal([]byte(response), tip); err != nil {
		// If parsing fails, use the raw response as the tip
		tip.Tip = response
	}

	return tip, nil
}

// RunDailyCoaching generates tips for all active families
// Called by the scheduler daily
func (pc *ParentCoach) RunDailyCoaching() {
	fmt.Printf("[ParentCoach] Running daily coaching at %s\n", time.Now().Format("15:04"))

	// In production:
	// 1. Query all active families
	// 2. For each family, load children's profiles and paths
	// 3. Call GenerateDailyTip
	// 4. Store the tip (e.g., in a daily coaching record or the family's profile)
	// 5. Notify guardian via l8notify if configured
	//
	// Each call logs a PARENT_COACHING prompt to PromptLog automatically
	// (handled by the LLMClient → PromptLogger pipeline)

	fmt.Println("[ParentCoach] Daily coaching complete")
}

func (pc *ParentCoach) summarizeCurrentSkills(paths []*learn.LearningPath) string {
	if len(paths) == 0 {
		return "no active learning paths"
	}
	var skills []string
	for _, p := range paths {
		if p.CurrentSkillId != "" {
			skills = append(skills, p.CurrentSkillId)
		}
	}
	if len(skills) == 0 {
		return "starting new skills"
	}
	return strings.Join(skills, ", ")
}

func (pc *ParentCoach) summarizeLearningStyles(profiles []*learn.StudentProfile) string {
	var modes []string
	for _, p := range profiles {
		if p.LearningStyle != nil && len(p.LearningStyle.PreferredModes) > 0 {
			modes = append(modes, p.LearningStyle.PreferredModes...)
		}
	}
	if len(modes) == 0 {
		return "not yet determined"
	}
	return strings.Join(modes, ", ")
}

func (pc *ParentCoach) summarizeInterests(profiles []*learn.StudentProfile) string {
	var interests []string
	for _, p := range profiles {
		if p.Motivation != nil && len(p.Motivation.HighInterestActivities) > 0 {
			interests = append(interests, p.Motivation.HighInterestActivities...)
		}
	}
	if len(interests) == 0 {
		return "not yet determined"
	}
	return strings.Join(interests, ", ")
}
