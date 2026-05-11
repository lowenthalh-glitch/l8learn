/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package worksheets

import (
	"github.com/saichler/l8common/go/common"
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newWorksheetServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.Worksheet{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.Worksheet).WorksheetId }, "WorksheetId").
		Require(func(v interface{}) string { return v.(*learn.Worksheet).Name }, "Name").
		Require(func(v interface{}) string { return v.(*learn.Worksheet).TeacherId }, "TeacherId").
		Enum(func(v interface{}) int32 { return int32(v.(*learn.Worksheet).Subject) }, "Subject", learn.SubjectType_name).
		Enum(func(v interface{}) int32 { return int32(v.(*learn.Worksheet).Status) }, "Status", learn.WorksheetStatus_name).
		Enum(func(v interface{}) int32 { return int32(v.(*learn.Worksheet).Layout) }, "Layout", learn.WorksheetLayout_name).
		After(onWorksheetChange).
		Build()
}

func onWorksheetChange(elem interface{}, action ifs.Action, notify bool, vnic ifs.IVNic) (interface{}, bool, error) {
	ws := elem.(*learn.Worksheet)

	if action == ifs.POST && ws.Status == learn.WorksheetStatus_WORKSHEET_STATUS_DRAFT {
		// TODO: Generate PDF
		// 1. Fetch questions from skill_ids / activity_ids
		// 2. Apply shuffle, difficulty filter, question_count limit
		// 3. If personalized=true, generate per-student variants based on SkillMastery
		// 4. Render to PDF using Go PDF library (jung-kurt/gofpdf)
		// 5. Generate answer key PDF
		// 6. Store both via Layer8FileUpload
		// 7. Update ws.PdfStoragePath, ws.AnswerKeyPath
		// 8. Update status: DRAFT → GENERATED
	}

	if action == ifs.PUT && ws.Status == learn.WorksheetStatus_WORKSHEET_STATUS_SCORED {
		// TODO: Feed scores to mastery
		// 1. For each WorksheetScore in ws.Scores
		// 2. Update SkillMastery for each student/skill
		// 3. Trigger LearningPath recalculation
	}

	return nil, true, nil
}
