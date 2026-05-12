/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * Lesson Feedback Loop — analyzes completed GeneratedLessons,
 * generates AI observations, and feeds results back into the
 * StudentProfile and content effectiveness tracking.
 */
package engine

import (
	"fmt"
	"time"

	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

// LessonFeedback orchestrates the post-lesson analysis pipeline
type LessonFeedback struct {
	vnic      ifs.IVNic
	llmClient LLMClient
	updater   *ProfileUpdater
}

func NewLessonFeedback(vnic ifs.IVNic, llmClient LLMClient, updater *ProfileUpdater) *LessonFeedback {
	return &LessonFeedback{vnic: vnic, llmClient: llmClient, updater: updater}
}

// OnLessonComplete runs the full feedback loop after a GeneratedLesson is completed
// 1. Generate AI observation analyzing student performance
// 2. Update StudentProfile with theme/approach preferences
// 3. Record content effectiveness data
func (lf *LessonFeedback) OnLessonComplete(lesson *learn.GeneratedLesson) {
	if lesson == nil || lesson.StudentId == "" {
		return
	}

	// Generate AI observation
	observation := lf.generateObservation(lesson)
	lesson.AiObservation = observation

	// Update student profile with lesson outcomes
	lf.updateProfileFromLesson(lesson)

	// Record effectiveness for analytics
	lf.recordEffectiveness(lesson)
}

// generateObservation calls the LLM to analyze lesson performance
func (lf *LessonFeedback) generateObservation(lesson *learn.GeneratedLesson) string {
	accuracy := float64(0)
	if lesson.QuestionsTotal > 0 {
		accuracy = float64(lesson.QuestionsCorrect) / float64(lesson.QuestionsTotal)
	}

	passed := lesson.QuestionsCorrect >= lesson.MinCorrectToPass
	advanced := lesson.QuestionsCorrect >= lesson.MinCorrectToAdvance

	lessonSummary := fmt.Sprintf(
		"Title: %s, Topic: %s, Theme: %s, Difficulty: %s, "+
			"Questions: %d/%d correct (%.0f%%), Duration: %d min (estimated %d), "+
			"Passed: %v, Advanced: %v, Strategy: %s",
		lesson.Title, lesson.Topic, lesson.Theme,
		lesson.Difficulty.String(),
		lesson.QuestionsCorrect, lesson.QuestionsTotal, accuracy*100,
		lesson.ActualMinutes, lesson.EstimatedMinutes,
		passed, advanced, lesson.OnStruggleStrategy,
	)

	systemPrompt, userMessage := BuildLessonObservationPrompt(
		lessonSummary,
		lesson.StudentId,
		lesson.Theme,
		fmt.Sprintf("%d", lesson.ActualMinutes),
	)

	response, err := lf.llmClient.Call(
		learn.LLMPromptType_LLM_PROMPT_TYPE_CONTENT_ANALYSIS,
		systemPrompt, userMessage, lesson.StudentId,
	)
	if err != nil {
		fmt.Printf("[LessonFeedback] Error generating observation: %v\n", err)
		return fmt.Sprintf("Lesson completed: %d/%d correct", lesson.QuestionsCorrect, lesson.QuestionsTotal)
	}

	return response
}

// updateProfileFromLesson updates StudentProfile based on lesson outcomes
func (lf *LessonFeedback) updateProfileFromLesson(lesson *learn.GeneratedLesson) {
	// In production: fetch the student's profile, update, and PUT back
	// For now, build the profile update from lesson data
	profile := &learn.StudentProfile{
		ProfileId: lesson.StudentId + "-profile",
		StudentId: lesson.StudentId,
	}

	accuracy := float64(0)
	if lesson.QuestionsTotal > 0 {
		accuracy = float64(lesson.QuestionsCorrect) / float64(lesson.QuestionsTotal)
	}

	// Update learning style: track which themes produce good results
	if profile.LearningStyle == nil {
		profile.LearningStyle = &learn.LearningStyle{}
	}
	if accuracy >= 0.75 && lesson.Theme != "" {
		// Theme worked well — add to preferred interests
		addUniqueString(&profile.LearningStyle.WorksBestWith, "theme:"+lesson.Theme)
	}

	// Update attention from actual vs estimated duration
	if profile.Attention == nil {
		profile.Attention = &learn.AttentionRegulationProfile{}
	}
	if lesson.ActualMinutes > 0 {
		profile.Attention.FocusAcademicTaskMinutes = lesson.ActualMinutes
	}

	// Track struggle patterns in behavior profile
	if accuracy < 0.5 {
		if profile.Behavior == nil {
			profile.Behavior = &learn.BehaviorProfile{}
		}
		note := fmt.Sprintf("struggled_with:%s_at_%s",
			lesson.Topic, lesson.Difficulty.String())
		addUniqueString(&profile.Behavior.AvoidanceBehaviors, note)
	}

	profile.LastUpdated = time.Now().Unix()

	fmt.Printf("[LessonFeedback] Profile updated for student %s: accuracy=%.0f%%, theme=%s\n",
		lesson.StudentId, accuracy*100, lesson.Theme)
}

// recordEffectiveness tracks which themes and approaches produce results
func (lf *LessonFeedback) recordEffectiveness(lesson *learn.GeneratedLesson) {
	accuracy := float64(0)
	if lesson.QuestionsTotal > 0 {
		accuracy = float64(lesson.QuestionsCorrect) / float64(lesson.QuestionsTotal)
	}

	// Compute efficiency: accuracy per minute
	efficiency := float64(0)
	if lesson.ActualMinutes > 0 {
		efficiency = accuracy / float64(lesson.ActualMinutes)
	}

	// In production: POST to CntEffect service
	// For now, log the effectiveness data
	fmt.Printf("[LessonFeedback] Content effectiveness: topic=%s, theme=%s, "+
		"difficulty=%s, accuracy=%.0f%%, duration=%d min, efficiency=%.4f\n",
		lesson.Topic, lesson.Theme, lesson.Difficulty.String(),
		accuracy*100, lesson.ActualMinutes, efficiency)
}

// BuildLessonObservationPrompt creates the AI prompt for analyzing lesson performance
func BuildLessonObservationPrompt(lessonSummary, studentId, theme, duration string) (string, string) {
	return NewPromptBuilder(learn.LLMPromptType_LLM_PROMPT_TYPE_CONTENT_ANALYSIS).
		SetRole("adaptive learning lesson analyst").
		AddRule("Analyze what worked and what didn't in this lesson").
		AddRule("Note if the theme engaged the student (based on time spent vs estimated)").
		AddRule("Identify specific concepts that need reinforcement").
		AddRule("Suggest adjustments for the next lesson").
		AddRule("Keep observations concise — 2-3 sentences max").
		AddContext("lesson_performance", lessonSummary).
		AddContext("student_id", studentId).
		AddContext("lesson_theme", theme).
		AddContext("actual_duration", duration+" minutes").
		SetReturnFormat(`{"observation":"...","themeEffective":true/false,"conceptsToReinforce":["..."],"nextLessonAdjustments":"..."}`).
		Build()
}
