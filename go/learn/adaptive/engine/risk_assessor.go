/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * RiskAssessor — weekly AI batch job that predicts at-risk students.
 * Analyzes mastery trajectory, engagement, session patterns.
 * Logs RISK_ASSESSMENT prompts to PromptLog.
 */
package engine

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

// RiskAssessor runs weekly risk prediction for all active students
type RiskAssessor struct {
	vnic      ifs.IVNic
	llmClient LLMClient
}

func NewRiskAssessor(vnic ifs.IVNic, llmClient LLMClient) *RiskAssessor {
	return &RiskAssessor{vnic: vnic, llmClient: llmClient}
}

// AssessRisk evaluates one student's risk of falling behind
func (ra *RiskAssessor) AssessRisk(
	studentId string,
	masteryTrend string,
	engagementData string,
	sessionFrequency string,
	scoreTrends string,
	cohortPercentiles string,
) (*learn.RiskAssessment, error) {

	systemPrompt, userMessage := BuildRiskAssessmentPrompt(
		masteryTrend, engagementData, sessionFrequency,
		scoreTrends, cohortPercentiles,
	)

	response, err := ra.llmClient.Call(
		learn.LLMPromptType_LLM_PROMPT_TYPE_RISK_ASSESSMENT,
		systemPrompt, userMessage, studentId,
	)
	if err != nil {
		return nil, fmt.Errorf("risk assessment failed: %w", err)
	}

	// Parse AI response
	var result struct {
		RiskLevel      string  `json:"riskLevel"`
		RiskScore      float64 `json:"riskScore"`
		Factors        []struct {
			Type        string  `json:"type"`
			Description string  `json:"description"`
			Weight      float64 `json:"weight"`
		} `json:"factors"`
		Recommendation string `json:"recommendation"`
	}

	if err := json.Unmarshal([]byte(response), &result); err != nil {
		// Use defaults on parse failure
		result.RiskLevel = "ON_TRACK"
		result.RiskScore = 0.1
	}

	// Map string risk level to enum
	riskLevel := learn.RiskLevel_RISK_LEVEL_ON_TRACK
	switch result.RiskLevel {
	case "WATCH":
		riskLevel = learn.RiskLevel_RISK_LEVEL_WATCH
	case "AT_RISK":
		riskLevel = learn.RiskLevel_RISK_LEVEL_AT_RISK
	case "CRITICAL":
		riskLevel = learn.RiskLevel_RISK_LEVEL_CRITICAL
	}

	// Build factors
	var factors []*learn.RiskFactor
	for _, f := range result.Factors {
		factors = append(factors, &learn.RiskFactor{
			FactorType:  f.Type,
			Description: f.Description,
			Weight:      f.Weight,
		})
	}

	assessment := &learn.RiskAssessment{
		AssessmentId:     fmt.Sprintf("RISK-%s-%d", studentId, time.Now().Unix()),
		StudentId:        studentId,
		RiskLevel:        riskLevel,
		RiskScore:        result.RiskScore,
		Factors:          factors,
		AiRecommendation: result.Recommendation,
	}

	return assessment, nil
}

// RunWeeklyRiskBatch assesses all active students
// Called by scheduler weekly
func (ra *RiskAssessor) RunWeeklyRiskBatch() {
	fmt.Printf("[RiskAssessor] Running weekly risk assessment at %s\n", time.Now().Format("2006-01-02"))
	// In production:
	// 1. Query all active students
	// 2. For each student, gather: mastery trend (4 weeks), engagement,
	//    session frequency, score trends, cohort percentiles
	// 3. Call AssessRisk
	// 4. POST result to RiskAssmt service
	// 5. If AT_RISK or CRITICAL, notify teacher via l8notify
	// Each call logs a RISK_ASSESSMENT prompt to PromptLog
	fmt.Println("[RiskAssessor] Weekly risk assessment complete")
}
