/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package mocks

import (
	"fmt"
	"github.com/saichler/l8learn/go/types/learn"
	"math/rand"
	"time"
)

func generatePromptLogs(store *MockDataStore) []*learn.LLMPromptLog {
	var list []*learn.LLMPromptLog
	now := time.Now().Unix()

	types := []learn.LLMPromptType{
		learn.LLMPromptType_LLM_PROMPT_TYPE_PATH_DECISION,
		learn.LLMPromptType_LLM_PROMPT_TYPE_PROFILE_UPDATE,
		learn.LLMPromptType_LLM_PROMPT_TYPE_RISK_ASSESSMENT,
		learn.LLMPromptType_LLM_PROMPT_TYPE_PARENT_COACHING,
		learn.LLMPromptType_LLM_PROMPT_TYPE_WORKSHEET_SCAN,
	}

	for i := 0; i < 50; i++ {
		promptType := types[i%len(types)]
		hasPII := i%10 == 0 // 10% have PII detected

		var piiFields []string
		if hasPII {
			piiFields = []string{"student_name"}
		}

		list = append(list, &learn.LLMPromptLog{
			LogId:               fmt.Sprintf("PL-%04d", i+1),
			Type:                promptType,
			StudentId:           pickRef(store.StudentIDs, i),
			Mode:                learn.LLMMode_LLM_MODE_SIMULATE,
			SystemPrompt:        fmt.Sprintf("You are a %s for adaptive learning.", promptType.String()),
			UserMessage:         fmt.Sprintf("Context for student %s: mastery data, interaction patterns.", pickRef(store.StudentIDs, i)),
			SystemPromptTokens:  int32(200 + rand.Intn(800)),
			UserMessageTokens:   int32(300 + rand.Intn(1200)),
			Response:            `{"simulated":true}`,
			ResponseTokens:      int32(50 + rand.Intn(200)),
			ResponseTimeMs:      int64(10 + rand.Intn(50)),
			ContainsStudentName: hasPII,
			ContainsPii:         hasPII,
			PiiFieldsFound:      piiFields,
			DataMasked:          true,
			Timestamp:           now - int64(rand.Intn(7*24*3600)),
			TriggeredBy:         promptType.String(),
		})
	}
	return list
}

func generateLLMConfig() *learn.LLMConfig {
	return &learn.LLMConfig{
		ConfigId:             "LLMCFG-001",
		Mode:                 learn.LLMMode_LLM_MODE_SIMULATE,
		ApiProvider:          "anthropic",
		ModelName:            "claude-sonnet-4-6",
		MaxTokens:            4096,
		Temperature:          0.7,
		PiiMaskingEnabled:    true,
		PromptLoggingEnabled: true,
		MaxDailyCalls:        1000,
		CallsToday:           47,
	}
}
