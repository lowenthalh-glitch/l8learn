/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package progress

import (
	"github.com/saichler/l8common/go/common"
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newProgressServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.ProgressReport{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.ProgressReport).ReportId }, "ReportId").
		Require(func(v interface{}) string { return v.(*learn.ProgressReport).StudentId }, "StudentId").
		Enum(func(v interface{}) int32 { return int32(v.(*learn.ProgressReport).Subject) }, "Subject", learn.SubjectType_name).
		Enum(func(v interface{}) int32 { return int32(v.(*learn.ProgressReport).Period) }, "Period", learn.ReportPeriod_name).
		Enum(func(v interface{}) int32 { return int32(v.(*learn.ProgressReport).Engagement) }, "Engagement", learn.EngagementLevel_name).
		After(onProgressGenerated).
		Build()
}

func onProgressGenerated(elem interface{}, action ifs.Action, notify bool, vnic ifs.IVNic) (interface{}, bool, error) {
	report := elem.(*learn.ProgressReport)
	if action == ifs.POST {
		// After generating a progress report:
		// 1. If guardian has receives_reports=true, send via l8notify
		// 2. Generate AI narrative summary (report.AiSummary)
		// 3. Generate AI recommendations (report.AiRecommendations)
		_ = report
	}
	return nil, true, nil
}
