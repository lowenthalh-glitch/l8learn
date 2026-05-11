/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package skills

import (
	"github.com/saichler/l8common/go/common"
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newSkillServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.Skill{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.Skill).SkillId }, "SkillId").
		Require(func(v interface{}) string { return v.(*learn.Skill).Name }, "Name").
		Enum(func(v interface{}) int32 { return int32(v.(*learn.Skill).Subject) }, "Subject", learn.SubjectType_name).
		Enum(func(v interface{}) int32 { return int32(v.(*learn.Skill).GradeLevel) }, "GradeLevel", learn.GradeLevel_name).
		Build()
}
