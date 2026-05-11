/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package worksheetscans

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newWorksheetScanServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.WorksheetScan{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.WorksheetScan).ScanId }, "ScanId").
		Require(func(v interface{}) string { return v.(*learn.WorksheetScan).WorksheetId }, "WorksheetId").
		Require(func(v interface{}) string { return v.(*learn.WorksheetScan).TeacherId }, "TeacherId").
		Enum(func(v interface{}) int32 { return int32(v.(*learn.WorksheetScan).Status) }, "Status", learn.ScanStatus_name).
		After(onScanChange).
		Build()
}

func onScanChange(elem interface{}, action ifs.Action, notify bool, vnic ifs.IVNic) (interface{}, bool, error) {
	scan := elem.(*learn.WorksheetScan)

	if action == ifs.POST {
		// Image uploaded — trigger AI extraction + grading pipeline:
		// 1. Load original worksheet definition (questions + correct answers)
		// 2. Send scanned image + worksheet def to AI (l8agent LLM)
		// 3. AI extracts student name + answers from handwriting
		// 4. Auto-grade each answer using math equivalence evaluator
		// 5. Flag low-confidence items (confidence < 0.7) for teacher review
		// 6. Update scan status → REVIEW (has flags) or COMPLETE (no flags)
		// 7. If COMPLETE, auto-save scores to WorksheetScore + SkillMastery
		scan.Status = learn.ScanStatus_SCAN_STATUS_PROCESSING
	}

	if action == ifs.PUT && scan.Status == learn.ScanStatus_SCAN_STATUS_COMPLETE {
		// Teacher reviewed flagged items and marked complete:
		// 1. Recalculate score with teacher overrides
		// 2. Create WorksheetScore on parent Worksheet
		// 3. Update SkillMastery for student per skill
		// 4. Trigger LearningPath recalculation
		// 5. Notify guardian if configured (l8notify)
		// 6. Log event (l8events) for audit trail
	}

	return nil, true, nil
}
