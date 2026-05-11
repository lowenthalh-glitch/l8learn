/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package classrooms

import (
	"github.com/saichler/l8common/go/common"
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newClassroomServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.Classroom{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.Classroom).ClassroomId }, "ClassroomId").
		Require(func(v interface{}) string { return v.(*learn.Classroom).Name }, "Name").
		Require(func(v interface{}) string { return v.(*learn.Classroom).PrimaryTeacherId }, "PrimaryTeacherId").
		Require(func(v interface{}) string { return v.(*learn.Classroom).SchoolId }, "SchoolId").
		Enum(func(v interface{}) int32 { return int32(v.(*learn.Classroom).GradeLevel) }, "GradeLevel", learn.GradeLevel_name).
		Build()
}
