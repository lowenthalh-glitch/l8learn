/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package evalimports

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newEvalImportServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.EvalImport{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.EvalImport).ImportId }, "ImportId").
		Require(func(v interface{}) string { return v.(*learn.EvalImport).StudentId }, "StudentId").
		After(onEvalImportChange).
		Build()
}

func onEvalImportChange(elem interface{}, action ifs.Action, notify bool, vnic ifs.IVNic) (interface{}, bool, error) {
	eval := elem.(*learn.EvalImport)
	if action == ifs.POST {
		// PDF uploaded:
		// 1. Read PDF text via Layer8FileUpload
		// 2. Detect document type
		// 3. Send to LLM (or simulator) with masked profile
		// 4. Store extracted findings as PENDING
		// 5. Notify guardian: "New evaluation imported — review needed"
		_ = eval
	}
	if action == ifs.PUT && eval.AppliedToProfile {
		// Guardian reviewed and approved:
		// 1. For each ACCEPTED finding → update StudentProfile field
		// 2. For each EDITED finding → use edited_value
		// 3. Update therapy services if new provider/goals found
		// 4. Log to l8events for audit trail
		// 5. Trigger adaptive engine recalibration
	}
	return nil, true, nil
}
