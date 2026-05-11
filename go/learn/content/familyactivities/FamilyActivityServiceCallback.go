/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package familyactivities

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newFamilyActivityServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.FamilyActivity{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.FamilyActivity).FamilyActivityId }, "FamilyActivityId").
		Require(func(v interface{}) string { return v.(*learn.FamilyActivity).Name }, "Name").
		After(onFamilyActivityComplete).
		Build()
}

func onFamilyActivityComplete(elem interface{}, action ifs.Action, notify bool, vnic ifs.IVNic) (interface{}, bool, error) {
	fa := elem.(*learn.FamilyActivity)
	if action == ifs.PUT {
		// Activity marked complete by parent:
		// 1. Create Interaction records for EACH participating student (from fa.Roles)
		// 2. Update SkillMastery for each student's assigned skills
		// 3. Credit parent-logged time to StateCompliance.InstructionHoursLogged
		// 4. Log subjects covered to StateCompliance.SubjectsCovered
		_ = fa
	}
	return nil, true, nil
}
