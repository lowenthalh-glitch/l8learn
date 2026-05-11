/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package mastery

import (
	"github.com/saichler/l8common/go/common"
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newMasteryServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.SkillMastery{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.SkillMastery).MasteryId }, "MasteryId").
		Require(func(v interface{}) string { return v.(*learn.SkillMastery).StudentId }, "StudentId").
		Require(func(v interface{}) string { return v.(*learn.SkillMastery).SkillId }, "SkillId").
		Enum(func(v interface{}) int32 { return int32(v.(*learn.SkillMastery).Level) }, "Level", learn.MasteryLevel_name).
		After(onMasteryChange).
		Build()
}

func onMasteryChange(elem interface{}, action ifs.Action, notify bool, vnic ifs.IVNic) (interface{}, bool, error) {
	m := elem.(*learn.SkillMastery)
	if action == ifs.PATCH || action == ifs.PUT {
		// 1. Notify guardian via l8notify if mastery level changed
		// 2. Log mastery change event via l8events
		// 3. Update GrowthRecord for this student/subject
		// 4. Recalculate StandardMastery for all standards mapped to this skill
		// 5. Append MasterySnapshot to history
		_ = m
	}
	return nil, true, nil
}
