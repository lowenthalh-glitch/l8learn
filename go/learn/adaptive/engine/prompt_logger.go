/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package engine

import (
	"fmt"
	"time"

	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

// PromptLogger stores every LLM prompt to the PromptLog service for audit
type PromptLogger struct {
	vnic ifs.IVNic
}

func NewPromptLogger(vnic ifs.IVNic) *PromptLogger {
	return &PromptLogger{vnic: vnic}
}

func (l *PromptLogger) Log(
	promptType learn.LLMPromptType,
	studentId string,
	mode learn.LLMMode,
	systemPrompt string,
	userMessage string,
	response string,
	responseTimeMs int64,
	piiReport *PIIReport,
	dataMasked bool,
) {
	log := &learn.LLMPromptLog{
		LogId:               fmt.Sprintf("PL-%d", time.Now().UnixNano()),
		Type:                promptType,
		StudentId:           studentId,
		Mode:                mode,
		SystemPrompt:        systemPrompt,
		UserMessage:         userMessage,
		SystemPromptTokens:  int32(estimateTokens(systemPrompt)),
		UserMessageTokens:   int32(estimateTokens(userMessage)),
		Response:            response,
		ResponseTokens:      int32(estimateTokens(response)),
		ResponseTimeMs:      responseTimeMs,
		ContainsStudentName: piiReport.HasNames,
		ContainsPii:         piiReport.HasPII,
		PiiFieldsFound:      piiReport.Fields,
		DataMasked:          dataMasked,
		Timestamp:           time.Now().Unix(),
		TriggeredBy:         promptType.String(),
	}

	// POST to PromptLog service via VNet
	if l.vnic != nil {
		// TODO: Post log record to PromptLog service
		// For now, just log to console
		fmt.Printf("[PromptLogger] %s | student=%s | pii=%v | tokens=%d+%d | mode=%s\n",
			promptType.String(), studentId, piiReport.HasPII,
			log.SystemPromptTokens, log.UserMessageTokens, mode.String())
	}

	_ = log
}

// estimateTokens provides a rough token count (~4 chars per token)
func estimateTokens(text string) int {
	return len(text) / 4
}
