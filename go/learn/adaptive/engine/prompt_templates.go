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
		AddRule("Map every finding to the student profile schema using ONLY the valid profile_section and profile_field values listed below").
		AddRule("Only extract what is explicitly stated — do not infer").
		AddRule("Flag findings that contradict the current profile").
		AddRule("For string array fields (marked []), return the new_value as a comma-separated list").
		AddRule("Do NOT invent field names — use ONLY the exact field names from this schema").
		AddRule(validFieldSchema).
		AddContext("pdf_text_content", pdfText).
		AddContext("document_type", documentType).
		AddContext("current_student_profile", currentProfile).
		SetReturnFormat(`{"document_type":"...","professional":"...","findings":[{"profile_section":"...","profile_field":"...","current_value":"...","new_value":"...","source_text":"...","confidence":0.0-1.0}],"contradictions":[{"profile_field":"...","current_value":"...","document_says":"...","ai_recommendation":"..."}]}`).
		Build()
}

// EvalDocument represents one cleaned evaluation document for batch processing.
type EvalDocument struct {
	Text         string
	DocumentType string
	FileName     string
}

// BuildBatchEvalPrompt builds a prompt for analyzing multiple evaluation documents at once.
func BuildBatchEvalPrompt(documents []EvalDocument, currentProfile string) (string, string) {
	// Build concatenated document text
	var docText string
	for i, doc := range documents {
		docType := doc.DocumentType
		if docType == "" || docType == "EVAL_DOCUMENT_TYPE_UNSPECIFIED" {
			docType = "auto-detect from content"
		}
		docText += fmt.Sprintf("\n=== DOCUMENT %d (%s) ===\n%s\n", i+1, docType, doc.Text)
	}

	return NewPromptBuilder(learn.LLMPromptType_LLM_PROMPT_TYPE_EVAL_IMPORT).
		SetRole("professional evaluation report extractor analyzing multiple evaluation documents for one student").
		AddRule("Analyze ALL documents together to build a comprehensive student profile").
		AddRule("Map every finding to the student profile schema using ONLY the valid profile_section and profile_field values listed below").
		AddRule("Only extract what is explicitly stated — do not infer").
		AddRule("Flag findings that contradict the current profile or contradict across documents").
		AddRule("For string array fields (marked []), return the new_value as a comma-separated list").
		AddRule("Do NOT invent field names — use ONLY the exact field names from this schema").
		AddRule(validFieldSchema).
		AddContext("evaluation_documents", docText).
		AddContext("current_student_profile", currentProfile).
		SetReturnFormat(`{"document_type":"batch","professional":"multiple","findings":[{"profile_section":"...","profile_field":"...","current_value":"...","new_value":"...","source_text":"...","confidence":0.0-1.0}],"contradictions":[{"profile_field":"...","current_value":"...","document_says":"...","ai_recommendation":"..."}]}`).
		Build()
}

// validFieldSchema is the complete field reference shared by single and batch prompts.
var validFieldSchema = `VALID FIELD SCHEMA:
scores: overallAcademicReadiness(int), readingReadiness(int), mathReadiness(int), writingFineMotor(int), speechLanguage(int), attentionTaskStamina(int), grossMotor(int), socialMotivation(int), independenceDailyLiving(int), confidenceWithLearning(int)
strengths: socialEmotional[], playAndMotivation[], grossMotor[], academic[], communication[]
challenges: speechLanguage[], attentionExecutiveFunction[], fineMotorGraphomotor[], sensoryMotor[], academicReadiness[]
learningStyle: preferredModes[], bestSessionLengthMinutes(int), bestActivityLengthMinutes(int), maxSeatedWorkMinutes(int), breakFrequencyMinutes(int), bestTimeOfDay, bestLearningFormula[], worksBestWith[], worksPoorlyWith[]
attention: maxBookSittingTime, structuredTaskStamina, needsFrequentBreaks(bool), impulsivityPresent(bool), distractibilityPresent(bool), focusPreferredActivityMinutes(int), focusAcademicTaskMinutes(int), losingFocusSigns[], regulationSupports[]
motivation: highInterestActivities[], rewardPreferences[], avoidAsReward[], avoidedActivities[]
literacy: currentLevel, readingLevel, letterRecognition, phonemicAwareness, sightWords, comprehension, bookStaminaMinutes(int), primaryNeeds[], prioritySequence[], recommendedApproach[], avoid[]
math: level, recommendedMode, counting, numberRecognition, oneToOneCorrespondence, prioritySkills[], recommendedActivities[], errorPatterns[], avoid[]
speech: therapyNeed, receivesSpeechTherapy(bool), clarity, speechSounds[], currentGoals[], helpfulPrompts[]
speech.articulation: notedSoundChallenges[], phonologicalProcessesObserved[], functionalImpact[]
speech.expressiveLanguage: strengths[], needs[]
speech.receptiveLanguage: needs[], bestSupports[]
speech.languageHistory: lateLanguageOnset(bool), singleWordsApproximateAge, sentencesApproximateAge
fineMotor: therapyNeed, receivesOccupationalTherapy(bool), handDominance, pencilGrip, tracing, cutting, observedNeeds[], helpfulWritingSupports[], helpfulTools[], recommendedActivities[], avoid[]
fineMotor.nameWritingStatus: targetName, currentPattern, lettersNeedingSupport[], recommendedPractice[]
grossMotor: overallStatus, energyLevel, strengths[], needs[], learningUse, recommendedMovementBreaks[], movementBreakFrequencyMinutes(int)
sensory: sensoryPattern, sensitivities[], registrationNeeds[], functionalImpact[], helpfulSupports[], avoid[]
socialEmotional: confidence, peerInteraction, turnTaking, emotionNaming, strengths[], needs[], frustrationTriggers[], calmingStrategies[], recommendedEmotionalSupports[]
behavior: avoidanceBehaviors[], redirectStrategies[], successfulSupports[]
dailyLiving: homePracticeIdeas[]
dailyLiving.toileting: urination, bowelHygiene
dailyLiving.hygiene: handWashing, toothBrushing, bathing
dailyLiving.dressing: lowerBody, upperBody, socks, shoes, zippers, buttons, beltBuckle
dailyLiving.feeding: spoon, fork, knife, preference, foodSelectivity
health: medicalConditions[], visionConcerns, hearingConcerns, safetyConcerns[]`

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

// BuildSchedulePrompt builds a prompt for generating a weekly learning schedule.
func BuildSchedulePrompt(profile, programSettings, constraints string) (string, string) {
	return NewPromptBuilder(learn.LLMPromptType_LLM_PROMPT_TYPE_SCHEDULE_GENERATION).
		SetRole("adaptive learning schedule planner for homeschool families").
		AddRule("Generate a 5-day weekly schedule (Monday through Friday)").
		AddRule("Each day has 2 sessions (morning + afternoon) per the program settings").
		AddRule("Each session follows: movement warmup → academic task → break → therapy/skill task → creative → cleanup").
		AddRule("Respect attention span: maxSeatedWorkMinutes and breakFrequencyMinutes from the profile").
		AddRule("Alternate subjects across days — do not schedule the same subject every morning").
		AddRule("Include movement breaks between seated activities").
		AddRule("Factor in weather (outdoor vs indoor) and parent energy level if provided").
		AddRule("Each block must specify: blockId, startMinute (minutes from midnight, e.g. 540=9:00AM), durationMinutes, activityType, subject, description, parentRole, requiresParent").
		AddRule("Valid activityType values: movement_warmup, academic, therapy, creative, break, cleanup").
		AddRule("Valid subject values: math, literacy, speech, fine_motor, gross_motor, sensory, social, science, art, music").
		AddRule("Valid parentRole values: teach, guide, supervise, observe, none").
		AddContext("student_profile", profile).
		AddContext("program_settings", programSettings).
		AddContext("constraints", constraints).
		SetReturnFormat(`{"days":[{"dayName":"Monday","blocks":[{"blockId":"mon-1","startMinute":540,"durationMinutes":5,"activityType":"movement_warmup","subject":"gross_motor","description":"...","parentRole":"guide","requiresParent":true}]}]}`).
		Build()
}

// BuildLessonFromSchedulePrompt builds a prompt for generating a lesson for a specific schedule block.
func BuildLessonFromSchedulePrompt(profile, scheduleBlock, mastery string) (string, string) {
	return NewPromptBuilder(learn.LLMPromptType_LLM_PROMPT_TYPE_GENERATE_LESSON).
		SetRole("personalized lesson creator for adaptive learning. You replace the teacher — generate a complete, self-contained lesson").
		AddRule("Generate a lesson for the specific subject and duration described in the schedule block").
		AddRule("Include physical/hands-on step if learning style includes kinesthetic").
		AddRule("Theme the lesson around the student's interests when possible").
		AddRule("Duration must match the schedule block's durationMinutes").
		AddRule("Include progressive hints matched to the student's learning style").
		AddRule("Generate real, mathematically correct problems — never placeholder text").
		AddRule("Respect accommodations (IEP, 504, sensory needs)").
		AddRule("Each question must have exactly one correct answer").
		AddContext("student_profile", profile).
		AddContext("schedule_block", scheduleBlock).
		AddContext("current_mastery", mastery).
		SetReturnFormat(`{"title":"...","objective":"...","theme":"...","estimatedMinutes":N,"materialsNeeded":["..."],"parentInstructions":"...","steps":[{"stepNumber":1,"stepType":"physical|screen|discussion|worksheet|break","title":"...","instructions":"...","durationMinutes":N,"parentRole":"guide|observe|none","materialsInstructions":"...","questions":[{"prompt":"...","questionType":1-4,"correctAnswer":"...","explanation":"...","hints":["..."],"difficulty":1-5,"options":[{"text":"...","isCorrect":true,"feedback":"..."}]}]}],"minCorrectToAdvance":3,"minCorrectToPass":2,"onStruggleStrategy":"scaffold"}`).
		Build()
}
