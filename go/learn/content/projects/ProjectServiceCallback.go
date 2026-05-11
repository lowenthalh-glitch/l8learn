/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package projects

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newProjectServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.Project{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.Project).ProjectId }, "ProjectId").
		Require(func(v interface{}) string { return v.(*learn.Project).Name }, "Name").
		After(onProjectUpdate).
		Build()
}

func onProjectUpdate(elem interface{}, action ifs.Action, notify bool, vnic ifs.IVNic) (interface{}, bool, error) {
	proj := elem.(*learn.Project)
	if action == ifs.PUT || action == ifs.PATCH {
		// When milestones completed:
		// 1. Update SkillMastery for participating students per milestone skills
		// 2. Store deliverable photos/files for compliance portfolio
		// 3. Credit time toward StateCompliance.InstructionHoursLogged
		_ = proj
	}
	return nil, true, nil
}
