/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package compliance

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newComplianceServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.StateCompliance{}, vnic).
		BeforeAction(func(elem interface{}, action ifs.Action, vnic ifs.IVNic) error {
			if action == ifs.POST {
				common.GenerateID(&elem.(*learn.StateCompliance).ComplianceId)
			}
			return nil
		}).
		Require(func(v interface{}) string { return v.(*learn.StateCompliance).ComplianceId }, "ComplianceId").
		Require(func(v interface{}) string { return v.(*learn.StateCompliance).FamilyId }, "FamilyId").
		Require(func(v interface{}) string { return v.(*learn.StateCompliance).StateCode }, "StateCode").
		Require(func(v interface{}) string { return v.(*learn.StateCompliance).AcademicYear }, "AcademicYear").
		Build()
}

// NOTE: Daily scheduled job (outside of callback) will:
// 1. Auto-increment instruction_hours_logged from session data
// 2. Auto-track subjects_covered from activities completed
// 3. Alert parent 30 days before deadlines via l8notify
// 4. Auto-generate portfolio PDF when portfolio review is due
