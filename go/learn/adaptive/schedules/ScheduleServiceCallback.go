/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package schedules

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newScheduleServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.DailySchedule{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.DailySchedule).ScheduleId }, "ScheduleId").
		Require(func(v interface{}) string { return v.(*learn.DailySchedule).FamilyId }, "FamilyId").
		After(onScheduleCreate).
		Build()
}

func onScheduleCreate(elem interface{}, action ifs.Action, notify bool, vnic ifs.IVNic) (interface{}, bool, error) {
	sched := elem.(*learn.DailySchedule)
	if action == ifs.POST {
		// AI generates schedule blocks considering:
		// 1. All children's learning paths in this family
		// 2. Parent energy level (sched.ParentEnergy)
		// 3. Appointments (sched.Appointments)
		// 4. Weather (sched.Weather) — outdoor vs indoor activities
		// 5. Attention spans by age (shorter blocks for younger children)
		// 6. Insert sibling collaboration blocks where skills overlap
		// 7. Insert breaks between intensive sessions
		_ = sched
	}
	return nil, true, nil
}
