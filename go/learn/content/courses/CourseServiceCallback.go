/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package courses

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newCourseServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.Course{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.Course).CourseId }, "CourseId").
		Require(func(v interface{}) string { return v.(*learn.Course).Name }, "Name").
		Build()
}
