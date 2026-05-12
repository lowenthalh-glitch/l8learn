/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 * Pattern adapted from l8agent/go/llm/client.go
 */
package engine

import (
	"github.com/saichler/l8learn/go/types/learn"
)

// LLMClient is the interface for calling the LLM (real or simulated)
type LLMClient interface {
	Call(promptType learn.LLMPromptType, systemPrompt, userMessage, studentId string) (string, error)
	GetMode() learn.LLMMode
}

// NewLLMClient creates the appropriate client based on mode
func NewLLMClient(mode learn.LLMMode, apiKey string, masker *PIIMasker, logger *PromptLogger) LLMClient {
	switch mode {
	case learn.LLMMode_LLM_MODE_SIMULATE:
		return NewLLMSimulator(masker, logger)
	case learn.LLMMode_LLM_MODE_LOG_ONLY:
		return NewLogOnlyClient(masker, logger)
	case learn.LLMMode_LLM_MODE_LIVE:
		return NewLiveClient(apiKey, masker, logger)
	default:
		return NewLLMSimulator(masker, logger)
	}
}

// LogOnlyClient logs the prompt but returns empty response
type LogOnlyClient struct {
	masker *PIIMasker
	logger *PromptLogger
}

func NewLogOnlyClient(masker *PIIMasker, logger *PromptLogger) *LogOnlyClient {
	return &LogOnlyClient{masker: masker, logger: logger}
}

func (c *LogOnlyClient) Call(promptType learn.LLMPromptType, systemPrompt, userMessage, studentId string) (string, error) {
	tokenMap := NewTokenMap()
	piiReport := c.masker.Scan(systemPrompt + userMessage)
	maskedSystem := c.masker.MaskText(systemPrompt, tokenMap)
	maskedUser := c.masker.MaskText(userMessage, tokenMap)

	c.logger.Log(promptType, studentId, learn.LLMMode_LLM_MODE_LOG_ONLY,
		maskedSystem, maskedUser, "", 0, piiReport, true)

	return "{}", nil
}

func (c *LogOnlyClient) GetMode() learn.LLMMode {
	return learn.LLMMode_LLM_MODE_LOG_ONLY
}
