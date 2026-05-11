/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package engine

import (
	"fmt"
	"strings"

	"github.com/saichler/l8learn/go/types/learn"
)

// PromptBuilder constructs LLM prompts with a shared structure.
// Prevents duplication across prompt types (plan-duplication-audit.md).
type PromptBuilder struct {
	promptType   learn.LLMPromptType
	systemRole   string
	rules        []string
	context      map[string]string
	returnFormat string
}

func NewPromptBuilder(promptType learn.LLMPromptType) *PromptBuilder {
	return &PromptBuilder{
		promptType: promptType,
		context:    make(map[string]string),
	}
}

func (b *PromptBuilder) SetRole(role string) *PromptBuilder {
	b.systemRole = role
	return b
}

func (b *PromptBuilder) AddRule(rule string) *PromptBuilder {
	b.rules = append(b.rules, rule)
	return b
}

func (b *PromptBuilder) AddContext(key, value string) *PromptBuilder {
	b.context[key] = value
	return b
}

func (b *PromptBuilder) SetReturnFormat(format string) *PromptBuilder {
	b.returnFormat = format
	return b
}

func (b *PromptBuilder) Build() (systemPrompt, userMessage string) {
	// Build system prompt
	var sys strings.Builder
	sys.WriteString(fmt.Sprintf("You are a %s.\n\n", b.systemRole))

	if len(b.rules) > 0 {
		sys.WriteString("Rules:\n")
		for _, rule := range b.rules {
			sys.WriteString(fmt.Sprintf("- %s\n", rule))
		}
		sys.WriteString("\n")
	}

	if b.returnFormat != "" {
		sys.WriteString(fmt.Sprintf("Return JSON with this structure:\n%s\n", b.returnFormat))
	}

	// Build user message from context
	var usr strings.Builder
	usr.WriteString("Context:\n")
	for key, value := range b.context {
		usr.WriteString(fmt.Sprintf("\n%s:\n%s\n", strings.ToUpper(key), value))
	}

	return sys.String(), usr.String()
}
