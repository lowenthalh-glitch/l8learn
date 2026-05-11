/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package students

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newStudentServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.Student{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.Student).StudentId }, "StudentId").
		Require(func(v interface{}) string { return v.(*learn.Student).FirstName }, "FirstName").
		Require(func(v interface{}) string { return v.(*learn.Student).LastName }, "LastName").
		Require(func(v interface{}) string { return v.(*learn.Student).SchoolId }, "SchoolId").
		Enum(func(v interface{}) int32 { return int32(v.(*learn.Student).GradeLevel) }, "GradeLevel", learn.GradeLevel_name).
		Enum(func(v interface{}) int32 { return int32(v.(*learn.Student).Status) }, "Status", learn.StudentStatus_name).
		Build()
}
