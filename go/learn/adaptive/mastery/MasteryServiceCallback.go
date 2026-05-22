/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package mastery

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newMasteryServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.SkillMastery{}, vnic).
		BeforeAction(func(elem interface{}, action ifs.Action, vnic ifs.IVNic) error {
			if action == ifs.POST {
				common.GenerateID(&elem.(*learn.SkillMastery).MasteryId)
			}
			return nil
		}).
		Require(func(v interface{}) string { return v.(*learn.SkillMastery).MasteryId }, "MasteryId").
		Require(func(v interface{}) string { return v.(*learn.SkillMastery).StudentId }, "StudentId").
		Require(func(v interface{}) string { return v.(*learn.SkillMastery).SkillId }, "SkillId").
		After(onMasteryChange).
		Build()
}

func onMasteryChange(elem interface{}, action ifs.Action, vnic ifs.IVNic) error {
	m := elem.(*learn.SkillMastery)
	if action == ifs.PATCH || action == ifs.PUT {
		// 1. Notify guardian via l8notify if mastery level changed
		// 2. Log mastery change event via l8events
		// 3. Update GrowthRecord for this student/subject
		// 4. Recalculate StandardMastery for all standards mapped to this skill
		// 5. Append MasterySnapshot to history
		// 6. Update StudentProfile readiness + subject sections
		//    ProfileUpdater.OnMasteryChange(m, skill, profile)
		//    (requires loading the skill and profile for this student)
		_ = m
	}
	return nil
}
