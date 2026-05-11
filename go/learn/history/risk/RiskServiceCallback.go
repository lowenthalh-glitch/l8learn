/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package risk

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newRiskServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.RiskAssessment{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.RiskAssessment).AssessmentId }, "AssessmentId").
		Require(func(v interface{}) string { return v.(*learn.RiskAssessment).StudentId }, "StudentId").
		Enum(func(v interface{}) int32 { return int32(v.(*learn.RiskAssessment).Subject) }, "Subject", learn.SubjectType_name).
		Enum(func(v interface{}) int32 { return int32(v.(*learn.RiskAssessment).RiskLevel) }, "RiskLevel", learn.RiskLevel_name).
		After(onRiskAssessed).
		Build()
}

func onRiskAssessed(elem interface{}, action ifs.Action, notify bool, vnic ifs.IVNic) (interface{}, bool, error) {
	assessment := elem.(*learn.RiskAssessment)
	if action == ifs.POST && assessment.RiskLevel >= learn.RiskLevel_RISK_LEVEL_AT_RISK {
		// Student flagged at AT_RISK or CRITICAL:
		// 1. Notify teacher via l8notify with AI recommendation
		// 2. Log event via l8events for audit trail
		// 3. Add to teacher's dashboard "Needs Attention" list
		_ = assessment
	}
	return nil, true, nil
}

// NOTE: Risk assessments are COMPUTED by weekly AI batch job:
//   1. For each active student:
//   2. Gather inputs: mastery trajectory, engagement signals, session patterns,
//      days since last session, score trends, historical peer data
//   3. Invoke AI (l8agent LLM) with structured output:
//      - risk_score (0.0-1.0)
//      - risk_level (ON_TRACK, WATCH, AT_RISK, CRITICAL)
//      - factors[] (type, description, weight, evidence)
//      - ai_recommendation (narrative intervention suggestion)
//      - recommended_skill_ids (skills needing immediate attention)
//   4. Store RiskAssessment (creates new or updates existing for this student+subject)
