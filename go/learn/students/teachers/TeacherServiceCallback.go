/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package teachers

import (
	"github.com/saichler/l8common/go/common"
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newTeacherServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.Teacher{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.Teacher).TeacherId }, "TeacherId").
		Require(func(v interface{}) string { return v.(*learn.Teacher).FirstName }, "FirstName").
		Require(func(v interface{}) string { return v.(*learn.Teacher).LastName }, "LastName").
		Require(func(v interface{}) string { return v.(*learn.Teacher).Email }, "Email").
		Require(func(v interface{}) string { return v.(*learn.Teacher).SchoolId }, "SchoolId").
		Enum(func(v interface{}) int32 { return int32(v.(*learn.Teacher).Role) }, "Role", learn.TeacherRole_name).
		Build()
}
