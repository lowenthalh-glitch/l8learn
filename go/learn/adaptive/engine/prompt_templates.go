/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * Prompt templates — config-only functions that use PromptBuilder.
 * No behavioral duplication: all templates call the shared builder.
 */
package engine

import (
	"fmt"

	"github.com/saichler/l8learn/go/types/learn"
)

func BuildPathDecisionPrompt(profile, mastery, interactions, activities string) (string, string) {
	return NewPromptBuilder(learn.LLMPromptType_LLM_PROMPT_TYPE_PATH_DECISION).
		SetRole("adaptive learning engine").
		AddRule("Never assign activities whose prerequisite skills are not PROFICIENT").
		AddRule("Respect adaptive_settings (max_consecutive_errors, break_frequency)").
		AddRule("Respect AI tutor settings (personality, should_avoid)").
		AddRule("If attention profile shows losing_focus_signs, insert break activities").
		AddRule("Use the student's preferred_modes and effective_activity_types").
		AddRule("Consider error_patterns to choose targeted activities").
		AddContext("student_profile", profile).
		AddContext("current_mastery", mastery).
		AddContext("recent_interactions", interactions).
		AddContext("available_activities", activities).
		SetReturnFormat(`{"nextActivities":[{"activityId":"...","skillId":"...","difficulty":1-5,"reason":"..."}],"reasoning":"..."}`).
		Build()
}

func BuildProfileUpdatePrompt(currentProfile, interactionSummary, masteryDeltas, sessionData string) (string, string) {
	return NewPromptBuilder(learn.LLMPromptType_LLM_PROMPT_TYPE_PROFILE_UPDATE).
		SetRole("student learning profile analyst").
		AddRule("Only update sections where data supports a change").
		AddRule("Flag concerning patterns (regression, disengagement, possible learning difficulties)").
		AddRule("Do not repeat unchanged sections").
		AddContext("current_profile", currentProfile).
		AddContext("last_7_days_interactions", interactionSummary).
		AddContext("mastery_changes", masteryDeltas).
		AddContext("session_patterns", sessionData).
		SetReturnFormat(`{"overallDescription":"...","mainStrengths":[...],"mainChallenges":[...],"updated_sections":{...}}`).
		Build()
}

func BuildRiskAssessmentPrompt(masteryTrend, engagementData, sessionFrequency, scoreTrends, cohortPercentiles string) (string, string) {
	return NewPromptBuilder(learn.LLMPromptType_LLM_PROMPT_TYPE_RISK_ASSESSMENT).
		SetRole("early warning system for student learning").
		AddRule("Be specific about contributing factors and recommended interventions").
		AddRule("Consider peer context for relative performance").
		AddContext("mastery_trajectory_4_weeks", masteryTrend).
		AddContext("engagement_signals", engagementData).
		AddContext("session_frequency", sessionFrequency).
		AddContext("score_trends", scoreTrends).
		AddContext("peer_comparison_anonymous", cohortPercentiles).
		SetReturnFormat(`{"riskLevel":"ON_TRACK|WATCH|AT_RISK|CRITICAL","riskScore":0.0-1.0,"factors":[{"type":"...","description":"...","weight":0.0-1.0}],"recommendation":"..."}`).
		Build()
}

func BuildParentCoachingPrompt(currentSkills, schedule, learningStyle, interests string) (string, string) {
	return NewPromptBuilder(learn.LLMPromptType_LLM_PROMPT_TYPE_PARENT_COACHING).
		SetRole("teaching coach for a homeschool parent").
		AddRule("Generate ONE actionable tip for today").
		AddRule("Use simple language. Be encouraging. Never criticize the parent.").
		AddRule("Connect the tip to what the child is currently working on").
		AddContext("child_current_focus_skills", currentSkills).
		AddContext("today_planned_activities", schedule).
		AddContext("child_learning_style", learningStyle).
		AddContext("child_interests", interests).
		SetReturnFormat(`{"tip":"...","activitySuggestion":"...","materials":"..."}`).
		Build()
}

func BuildWorksheetScanPrompt(profile, answers, handwriting, workPatterns string) (string, string) {
	return NewPromptBuilder(learn.LLMPromptType_LLM_PROMPT_TYPE_WORKSHEET_SCAN).
		SetRole("worksheet analysis expert for learning insights").
		AddRule("Beyond scoring right/wrong, analyze what the handwriting and work patterns reveal").
		AddRule("Distinguish motor issues from knowledge gaps").
		AddRule("Detect systematic errors vs random errors").
		AddContext("student_profile", profile).
		AddContext("extracted_answers", answers).
		AddContext("handwriting_analysis", handwriting).
		AddContext("work_pattern_analysis", workPatterns).
		SetReturnFormat(`{"fine_motor_updates":{},"math_updates":{},"attention_updates":{},"behavior_updates":{},"learning_style_updates":{},"insights":"...","recommendations":"..."}`).
		Build()
}

func BuildEvalImportPrompt(pdfText, documentType, currentProfile string) (string, string) {
	return NewPromptBuilder(learn.LLMPromptType_LLM_PROMPT_TYPE_EVAL_IMPORT).
		SetRole("professional evaluation report extractor").
		AddRule("Map every finding to the student profile schema").
		AddRule("Only extract what is explicitly stated — do not infer").
		AddRule("Flag findings that contradict the current profile").
		AddContext("pdf_text_content", pdfText).
		AddContext("document_type", documentType).
		AddContext("current_student_profile", currentProfile).
		SetReturnFormat(`{"document_type":"...","professional":"...","evaluation_date":"...","findings":[{"profile_section":"...","profile_field":"...","current_value":"...","new_value":"...","source_text":"...","confidence":0.0-1.0}],"contradictions":[...],"new_therapy_info":{...}}`).
		Build()
}

func BuildGenerateLessonPrompt(profile, mastery, skillGraph, interests, accommodations string, maxMinutes int, difficulty string) (string, string) {
	return NewPromptBuilder(learn.LLMPromptType_LLM_PROMPT_TYPE_GENERATE_LESSON).
		SetRole("personalized lesson creator for adaptive learning. You replace the teacher — generate a complete, self-contained lesson").
		AddRule("Include physical/hands-on step if learning style includes kinesthetic").
		AddRule("Theme the lesson around the student's interests when possible").
		AddRule("Duration must not exceed the student's attention profile").
		AddRule("Include progressive hints matched to the student's hint_style").
		AddRule("Generate real, mathematically correct problems — never placeholder text").
		AddRule("Respect accommodations (IEP, 504, sensory needs)").
		AddRule("Each question must have exactly one correct answer").
		AddContext("student_profile", profile).
		AddContext("current_mastery", mastery).
		AddContext("skill_prerequisites", skillGraph).
		AddContext("interests", interests).
		AddContext("accommodations", accommodations).
		AddContext("max_duration_minutes", fmt.Sprintf("%d", maxMinutes)).
		AddContext("target_difficulty", difficulty).
		SetReturnFormat(`{"title":"...","objective":"...","theme":"...","estimatedMinutes":N,"materialsNeeded":["..."],"parentInstructions":"...","steps":[{"stepNumber":1,"stepType":"physical|screen|discussion|worksheet|break","title":"...","instructions":"...","durationMinutes":N,"parentRole":"guide|observe|none","materialsInstructions":"...","questions":[{"prompt":"...","questionType":1-4,"correctAnswer":"...","explanation":"...","hints":["..."],"difficulty":1-5,"options":[{"text":"...","isCorrect":true,"feedback":"..."}]}]}],"worksheetContent":{"title":"...","problems":[{"prompt":"...","answer":"..."}]},"minCorrectToAdvance":3,"minCorrectToPass":2,"onStruggleStrategy":"scaffold"}`).
		Build()
}
