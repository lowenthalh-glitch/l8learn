/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package engine

import (
	"github.com/saichler/l8learn/go/types/learn"
	"time"
)

// LLMSimulator returns deterministic responses for testing
type LLMSimulator struct {
	masker *PIIMasker
	logger *PromptLogger
}

func NewLLMSimulator(masker *PIIMasker, logger *PromptLogger) *LLMSimulator {
	return &LLMSimulator{masker: masker, logger: logger}
}

func (s *LLMSimulator) Call(promptType learn.LLMPromptType, systemPrompt, userMessage, studentId string) (string, error) {
	start := time.Now()
	tokenMap := NewTokenMap()

	// Scan for PII
	piiReport := s.masker.Scan(systemPrompt + userMessage)

	// Mask PII before logging
	maskedSystem := s.masker.MaskText(systemPrompt, tokenMap)
	maskedUser := s.masker.MaskText(userMessage, tokenMap)

	// Generate deterministic response
	response := s.generateResponse(promptType)
	elapsed := time.Since(start).Milliseconds()

	// Log the prompt
	s.logger.Log(promptType, studentId, learn.LLMMode_LLM_MODE_SIMULATE,
		maskedSystem, maskedUser, response, elapsed, piiReport, true)

	return response, nil
}

func (s *LLMSimulator) GetMode() learn.LLMMode {
	return learn.LLMMode_LLM_MODE_SIMULATE
}

func (s *LLMSimulator) generateResponse(promptType learn.LLMPromptType) string {
	switch promptType {
	case learn.LLMPromptType_LLM_PROMPT_TYPE_PATH_DECISION:
		return `{"nextActivities":[{"activityId":"ACT-0001","skillId":"SKL-001","difficulty":3,"reason":"Simulated: targeting weakest skill"},{"activityId":"ACT-0002","skillId":"SKL-002","difficulty":2,"reason":"Simulated: confidence builder"}],"reasoning":"Simulated response — review prompt in PromptLog"}`

	case learn.LLMPromptType_LLM_PROMPT_TYPE_PROFILE_UPDATE:
		return `{"overallDescription":"Simulated profile update — review prompt in PromptLog","mainStrengths":["simulated_strength"],"mainChallenges":["simulated_challenge"]}`

	case learn.LLMPromptType_LLM_PROMPT_TYPE_RISK_ASSESSMENT:
		return `{"riskLevel":"ON_TRACK","riskScore":0.15,"factors":[{"factorType":"simulated","description":"Simulated risk — review prompt in PromptLog","weight":0.1}],"recommendation":"Simulated recommendation"}`

	case learn.LLMPromptType_LLM_PROMPT_TYPE_PARENT_COACHING:
		return `{"tip":"Simulated coaching tip — review prompt in PromptLog","activitySuggestion":"Try counting objects during grocery shopping","materials":"none needed"}`

	case learn.LLMPromptType_LLM_PROMPT_TYPE_WORKSHEET_SCAN:
		return `{"fine_motor_updates":{"pencil_grip":"simulated"},"math_updates":{"error_patterns":["simulated"]},"insights":"Simulated scan analysis — review prompt in PromptLog"}`

	case learn.LLMPromptType_LLM_PROMPT_TYPE_EVAL_IMPORT:
		return `{"document_type":"simulated","findings":[{"profile_section":"speech","profile_field":"clarity","current_value":"unknown","new_value":"developing","source_text":"Simulated finding","confidence":0.9}],"contradictions":[]}`

	case learn.LLMPromptType_LLM_PROMPT_TYPE_SCHEDULE_GENERATION:
		return `{"blocks":[{"startMinute":540,"durationMinutes":20,"subject":"math","description":"Simulated schedule block"}]}`

	case learn.LLMPromptType_LLM_PROMPT_TYPE_MODERATION:
		return `{"action":"approved","reason":"Simulated moderation"}`

	case learn.LLMPromptType_LLM_PROMPT_TYPE_GENERATE_LESSON:
		return `{"title":"Simulated Lesson — review prompt in PromptLog","objective":"Student will practice the target skill","theme":"general","estimatedMinutes":10,"materialsNeeded":[],"parentInstructions":"Observe and encourage","steps":[{"stepNumber":1,"stepType":"screen","title":"Practice Problems","instructions":"Answer these questions","durationMinutes":10,"parentRole":"none","questions":[{"prompt":"What is 7 x 8?","questionType":3,"correctAnswer":"56","explanation":"7 groups of 8 = 56","hints":["Try counting by 7s"],"difficulty":3},{"prompt":"What is 6 x 9?","questionType":3,"correctAnswer":"54","explanation":"6 groups of 9 = 54","hints":["Try 6 x 10 minus 6"],"difficulty":3}]}],"minCorrectToAdvance":2,"minCorrectToPass":1,"onStruggleStrategy":"scaffold"}`

	default:
		return `{"simulated":true,"message":"Review the prompt in PromptLog service"}`
	}
}
