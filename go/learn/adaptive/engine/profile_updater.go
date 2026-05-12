/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * ProfileUpdater — automatically updates StudentProfile from observed data.
 * Triggered by: SkillMastery changes, session completions, worksheet scans.
 * Weekly: LLM generates narrative summary via PROFILE_UPDATE prompt.
 */
package engine

import (
	"fmt"
	"time"

	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

// WorksheetInsights holds handwriting and work pattern data from a scan
type WorksheetInsights struct {
	OverallQuality        string
	NumberFormation        string
	EstimatedFocusMinutes int32
	QualityDegrades       bool
	ProblemsSkipped       int32
	ProblemsWithErasures  int32
}

// ProfileUpdater watches for data changes and updates StudentProfile
type ProfileUpdater struct {
	vnic      ifs.IVNic
	llmClient LLMClient
}

func NewProfileUpdater(vnic ifs.IVNic, llmClient LLMClient) *ProfileUpdater {
	return &ProfileUpdater{vnic: vnic, llmClient: llmClient}
}

// OnMasteryChange updates readiness scores and subject-specific sections
// Called from SkillMastery After callback when mastery level changes
func (pu *ProfileUpdater) OnMasteryChange(mastery *learn.SkillMastery, skill *learn.Skill, profile *learn.StudentProfile) {
	if profile == nil || skill == nil {
		return
	}

	// Update subject-specific section based on skill type
	switch skill.Subject {
	case learn.SubjectType_SUBJECT_TYPE_MATH:
		if profile.Math == nil {
			profile.Math = &learn.MathProfile{}
		}
		updateMathProfile(profile.Math, skill, mastery)
		if profile.Scores != nil {
			profile.Scores.MathReadiness = int32(mastery.CurrentAccuracy * 100)
		}

	case learn.SubjectType_SUBJECT_TYPE_READING:
		if profile.Literacy == nil {
			profile.Literacy = &learn.LiteracyProfile{}
		}
		updateLiteracyProfile(profile.Literacy, skill, mastery)
		if profile.Scores != nil {
			profile.Scores.ReadingReadiness = int32(mastery.CurrentAccuracy * 100)
		}
	}

	profile.LastUpdated = time.Now().Unix()
}

// OnSessionComplete updates learning_style, attention, motivation
// Called from LearningSession After callback when session status → COMPLETED
func (pu *ProfileUpdater) OnSessionComplete(session *learn.LearningSession, profile *learn.StudentProfile) {
	if profile == nil || session == nil {
		return
	}

	// Update attention profile from session duration patterns
	if profile.Attention == nil {
		profile.Attention = &learn.AttentionRegulationProfile{}
	}
	sessionMinutes := session.DurationSeconds / 60
	if sessionMinutes > 0 {
		// Running average of session focus time
		current := profile.Attention.FocusAcademicTaskMinutes
		if current == 0 {
			profile.Attention.FocusAcademicTaskMinutes = sessionMinutes
		} else {
			// Weighted average: 70% old, 30% new
			profile.Attention.FocusAcademicTaskMinutes = (current*7 + sessionMinutes*3) / 10
		}
	}

	// Update learning style from activity completion patterns
	if profile.LearningStyle == nil {
		profile.LearningStyle = &learn.LearningStyle{}
	}
	if sessionMinutes > 0 {
		current := profile.LearningStyle.BestSessionLengthMinutes
		if current == 0 {
			profile.LearningStyle.BestSessionLengthMinutes = sessionMinutes
		} else {
			profile.LearningStyle.BestSessionLengthMinutes = (current*7 + sessionMinutes*3) / 10
		}
	}

	// Detect engagement patterns from hints and correctness
	if profile.Motivation == nil {
		profile.Motivation = &learn.MotivationProfile{}
	}
	if session.QuestionsAnswered > 0 && session.QuestionsCorrect > 0 {
		accuracy := float64(session.QuestionsCorrect) / float64(session.QuestionsAnswered)
		if accuracy < 0.3 {
			// Struggling — record for behavior analysis
			addUniqueString(&profile.Motivation.AvoidedActivities, "current_difficulty_too_high")
		}
	}

	profile.LastUpdated = time.Now().Unix()
}

// OnWorksheetScanned updates fine_motor, attention, behavior
// Called from WorksheetScan After callback when scan is processed
func (pu *ProfileUpdater) OnWorksheetScanned(

	scan *learn.WorksheetScan,
	insights *WorksheetInsights,
	
	profile *learn.StudentProfile,
) {
	if profile == nil {
		return
	}

	// Update fine motor from handwriting analysis
	if true {
		if profile.FineMotor == nil {
			profile.FineMotor = &learn.FineMotorOTProfile{}
		}
		if insights.OverallQuality != "" {
			profile.FineMotor.PencilGrip = insights.OverallQuality
		}
		if insights.NumberFormation != "" {
			profile.FineMotor.NameWriting = insights.NumberFormation
		}
	}

	// Update attention from work patterns
	if true {
		if profile.Attention == nil {
			profile.Attention = &learn.AttentionRegulationProfile{}
		}
		if insights.EstimatedFocusMinutes > 0 {
			profile.Attention.FocusAcademicTaskMinutes = insights.EstimatedFocusMinutes
		}
		if insights.QualityDegrades {
			addUniqueString(&profile.Attention.LosingFocusSigns, "handwriting_degradation")
		}

		// Update behavior from patterns
		if profile.Behavior == nil {
			profile.Behavior = &learn.BehaviorProfile{}
		}
		if insights.ProblemsSkipped > 0 {
			addUniqueString(&profile.Behavior.AvoidanceBehaviors, "skips_when_unsure")
		}
		if insights.ProblemsWithErasures > 0 {
			addUniqueString(&profile.Behavior.SuccessfulSupports, "self_correction_via_erasure")
		}
	}

	profile.LastUpdated = time.Now().Unix()
}

// RunWeeklyProfileUpdate generates AI narrative summary for a student
// Called by weekly scheduler
func (pu *ProfileUpdater) RunWeeklyProfileUpdate(studentId string, profile *learn.StudentProfile) {
	if profile == nil {
		return
	}

	// Build prompt context from current profile
	profileSummary := fmt.Sprintf("Readiness: academic=%d, reading=%d, math=%d. "+
		"Learning style: session=%d min, modes=%v. "+
		"Attention: focus=%d min, signs=%v",
		safeReadiness(profile, "academic"),
		safeReadiness(profile, "reading"),
		safeReadiness(profile, "math"),
		safeSessionLength(profile),
		safeModes(profile),
		safeFocus(profile),
		safeFocusSigns(profile))

	systemPrompt, userMessage := BuildProfileUpdatePrompt(
		profileSummary,
		"last 7 days interaction data (summarized)",
		"mastery deltas from this week",
		"session patterns: count, duration, time-of-day",
	)

	// Call LLM (simulator in test mode)
	response, err := pu.llmClient.Call(
		learn.LLMPromptType_LLM_PROMPT_TYPE_PROFILE_UPDATE,
		systemPrompt, userMessage, studentId,
	)
	if err != nil {
		fmt.Printf("[ProfileUpdater] Error calling LLM: %v\n", err)
		return
	}

	// In production: parse response JSON and update profile fields
	// For now, just log that it happened
	_ = response
	profile.LastUpdated = time.Now().Unix()
}

// Helper functions

func updateMathProfile(mp *learn.MathProfile, skill *learn.Skill, mastery *learn.SkillMastery) {
	levelStr := mastery.Level.String()
	domain := skill.Domain
	switch {
	case contains(domain, "Addition"):
		mp.Addition = levelStr
	case contains(domain, "Subtraction"):
		mp.Subtraction = levelStr
	case contains(domain, "Multiplication"):
		mp.Multiplication = levelStr
	case contains(domain, "Division"):
		mp.Division = levelStr
	case contains(domain, "Fraction"):
		mp.Fractions = levelStr
	}
}

func updateLiteracyProfile(lp *learn.LiteracyProfile, skill *learn.Skill, mastery *learn.SkillMastery) {
	levelStr := mastery.Level.String()
	domain := skill.Domain
	switch {
	case contains(domain, "Phonics") || contains(domain, "Phonemic"):
		lp.PhonemicAwareness = levelStr
	case contains(domain, "Comprehension"):
		lp.Comprehension = levelStr
	case contains(domain, "Letter"):
		lp.LetterRecognition = levelStr
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && findSubstring(s, substr))
}

func findSubstring(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

func addUniqueString(slice *[]string, value string) {
	for _, v := range *slice {
		if v == value {
			return
		}
	}
	*slice = append(*slice, value)
}

func safeReadiness(p *learn.StudentProfile, field string) int32 {
	if p.Scores == nil { return 0 }
	switch field {
	case "academic": return p.Scores.OverallAcademicReadiness
	case "reading": return p.Scores.ReadingReadiness
	case "math": return p.Scores.MathReadiness
	}
	return 0
}

func safeSessionLength(p *learn.StudentProfile) int32 {
	if p.LearningStyle == nil { return 0 }
	return p.LearningStyle.BestSessionLengthMinutes
}

func safeModes(p *learn.StudentProfile) []string {
	if p.LearningStyle == nil { return nil }
	return p.LearningStyle.PreferredModes
}

func safeFocus(p *learn.StudentProfile) int32 {
	if p.Attention == nil { return 0 }
	return p.Attention.FocusAcademicTaskMinutes
}

func safeFocusSigns(p *learn.StudentProfile) []string {
	if p.Attention == nil { return nil }
	return p.Attention.LosingFocusSigns
}
