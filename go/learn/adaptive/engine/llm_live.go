/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * LiveClient — calls a real LLM API (Anthropic Claude Messages API).
 * Handles PII masking, cost tracking, daily call limits, and prompt logging.
 */
package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/saichler/l8learn/go/types/learn"
)

const (
	anthropicEndpoint   = "https://api.anthropic.com/v1/messages"
	anthropicAPIVersion = "2023-06-01"
	defaultModel        = "claude-sonnet-4-6-20250514"
	defaultMaxTokens    = 4096
)

// LiveClient calls a real LLM API with PII masking and cost tracking
type LiveClient struct {
	apiKey       string
	model        string
	maxTokens    int
	masker       *PIIMasker
	logger       *PromptLogger
	httpClient   *http.Client
	callsToday   int64
	maxDailyCall int64
	lastResetDay int
}

func NewLiveClient(apiKey string, masker *PIIMasker, logger *PromptLogger) *LiveClient {
	return &LiveClient{
		apiKey:       apiKey,
		model:        defaultModel,
		maxTokens:    defaultMaxTokens,
		masker:       masker,
		logger:       logger,
		httpClient:   &http.Client{Timeout: 60 * time.Second},
		maxDailyCall: 1000,
		lastResetDay: time.Now().YearDay(),
	}
}

// SetModel overrides the default model
func (c *LiveClient) SetModel(model string) { c.model = model }

// SetMaxTokens overrides the default max tokens
func (c *LiveClient) SetMaxTokens(max int) { c.maxTokens = max }

// SetMaxDailyCalls sets the daily call limit
func (c *LiveClient) SetMaxDailyCalls(max int64) { c.maxDailyCall = max }

func (c *LiveClient) GetMode() learn.LLMMode {
	return learn.LLMMode_LLM_MODE_LIVE
}

func (c *LiveClient) Call(promptType learn.LLMPromptType, systemPrompt, userMessage, studentId string) (string, error) {
	// Reset daily counter at midnight
	today := time.Now().YearDay()
	if today != c.lastResetDay {
		atomic.StoreInt64(&c.callsToday, 0)
		c.lastResetDay = today
	}

	// Check daily limit
	current := atomic.AddInt64(&c.callsToday, 1)
	if c.maxDailyCall > 0 && current > c.maxDailyCall {
		return "{}", fmt.Errorf("daily LLM call limit exceeded (%d/%d)", current, c.maxDailyCall)
	}

	start := time.Now()
	tokenMap := NewTokenMap()

	// Scan and mask PII
	piiReport := c.masker.Scan(systemPrompt + userMessage)
	maskedSystem := c.masker.MaskText(systemPrompt, tokenMap)
	maskedUser := c.masker.MaskText(userMessage, tokenMap)

	// Build API request
	reqBody := anthropicRequest{
		Model:     c.model,
		MaxTokens: c.maxTokens,
		System:    maskedSystem,
		Messages: []anthropicMessage{
			{Role: "user", Content: maskedUser},
		},
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "{}", fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", anthropicEndpoint, bytes.NewReader(bodyBytes))
	if err != nil {
		return "{}", fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.apiKey)
	req.Header.Set("anthropic-version", anthropicAPIVersion)

	// Call API
	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Log(promptType, studentId, learn.LLMMode_LLM_MODE_LIVE,
			maskedSystem, maskedUser, "", 0, piiReport, false)
		return "{}", fmt.Errorf("API call failed: %w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "{}", fmt.Errorf("read response: %w", err)
	}

	elapsed := time.Since(start).Milliseconds()

	if resp.StatusCode != 200 {
		c.logger.Log(promptType, studentId, learn.LLMMode_LLM_MODE_LIVE,
			maskedSystem, maskedUser, string(respBytes), elapsed, piiReport, false)
		return "{}", fmt.Errorf("API returned %d: %s", resp.StatusCode, string(respBytes))
	}

	// Parse response
	var apiResp anthropicResponse
	if err := json.Unmarshal(respBytes, &apiResp); err != nil {
		return "{}", fmt.Errorf("parse response: %w", err)
	}

	// Extract text from response content blocks
	responseText := ""
	for _, block := range apiResp.Content {
		if block.Type == "text" {
			responseText += block.Text
		}
	}

	// Unmask PII in response (restore original values)
	responseText = tokenMap.Unmask(responseText)

	// Log success
	c.logger.Log(promptType, studentId, learn.LLMMode_LLM_MODE_LIVE,
		maskedSystem, maskedUser, responseText, elapsed, piiReport, true)

	return responseText, nil
}

// Anthropic Messages API request/response types

type anthropicRequest struct {
	Model     string             `json:"model"`
	MaxTokens int                `json:"max_tokens"`
	System    string             `json:"system,omitempty"`
	Messages  []anthropicMessage `json:"messages"`
}

type anthropicMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type anthropicResponse struct {
	Content []anthropicContentBlock `json:"content"`
	Usage   anthropicUsage          `json:"usage"`
}

type anthropicContentBlock struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type anthropicUsage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}
