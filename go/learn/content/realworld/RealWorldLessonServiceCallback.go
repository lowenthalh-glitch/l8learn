/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package realworld

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newRealWorldLessonServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.RealWorldLesson{}, vnic).
		BeforeAction(func(elem interface{}, action ifs.Action, vnic ifs.IVNic) error {
			if action == ifs.POST {
				common.GenerateID(&elem.(*learn.RealWorldLesson).LessonId)
			}
			return nil
		}).
		Require(func(v interface{}) string { return v.(*learn.RealWorldLesson).LessonId }, "LessonId").
		Require(func(v interface{}) string { return v.(*learn.RealWorldLesson).Name }, "Name").
		After(onRealWorldComplete).
		Build()
}

func onRealWorldComplete(elem interface{}, action ifs.Action, vnic ifs.IVNic) error {
	rwl := elem.(*learn.RealWorldLesson)
	if action == ifs.PUT {
		// Real-world lesson logged by parent:
		// 1. Log challenges completed as Interaction records
		// 2. Update SkillMastery for participating students
		// 3. Credit time to StateCompliance.InstructionHoursLogged
		// 4. Store parent-uploaded photos in compliance portfolio
		_ = rwl
	}
	return nil
}
