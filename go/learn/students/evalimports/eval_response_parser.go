/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * Parses the LLM JSON response into EvalFinding and EvalContradiction protos.
 */
package evalimports

import (
	"fmt"
	"github.com/saichler/l8learn/go/learn/adaptive/engine"
	"github.com/saichler/l8learn/go/types/learn"
)

type llmEvalResponse struct {
	DocumentType string             `json:"document_type"`
	Professional string             `json:"professional"`
	Findings     []llmFinding       `json:"findings"`
	Contradicts  []llmContradiction `json:"contradictions"`
}

type llmFinding struct {
	ProfileSection string  `json:"profile_section"`
	ProfileField   string  `json:"profile_field"`
	CurrentValue   string  `json:"current_value"`
	NewValue       string  `json:"new_value"`
	SourceText     string  `json:"source_text"`
	Confidence     float64 `json:"confidence"`
}

type llmContradiction struct {
	ProfileField     string `json:"profile_field"`
	CurrentValue     string `json:"current_value"`
	DocumentSays     string `json:"document_says"`
	AIRecommendation string `json:"ai_recommendation"`
}

func parseEvalResponse(jsonResponse string) ([]*learn.EvalFinding, []*learn.EvalContradiction) {
	var resp llmEvalResponse
	if err := engine.ParseLLMResponse(jsonResponse, &resp); err != nil {
		fmt.Printf("[EvalParser] %s\n", err.Error())
		return nil, nil
	}

	var findings []*learn.EvalFinding
	for i, f := range resp.Findings {
		findings = append(findings, &learn.EvalFinding{
			FindingId:      fmt.Sprintf("F-%04d", i+1),
			ProfileSection: f.ProfileSection,
			ProfileField:   f.ProfileField,
			CurrentValue:   f.CurrentValue,
			NewValue:       f.NewValue,
			SourceText:     f.SourceText,
			Confidence:     f.Confidence,
			Status:         learn.EvalFindingStatus_EVAL_FINDING_STATUS_PENDING,
		})
	}

	var contradictions []*learn.EvalContradiction
	for _, c := range resp.Contradicts {
		contradictions = append(contradictions, &learn.EvalContradiction{
			ProfileField:     c.ProfileField,
			CurrentValue:     c.CurrentValue,
			DocumentSays:     c.DocumentSays,
			AiRecommendation: c.AIRecommendation,
			Resolution:       learn.EvalFindingStatus_EVAL_FINDING_STATUS_PENDING,
		})
	}

	return findings, contradictions
}
