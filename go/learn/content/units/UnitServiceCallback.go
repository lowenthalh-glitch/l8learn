/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package units

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newUnitServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.Unit{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.Unit).UnitId }, "UnitId").
		Require(func(v interface{}) string { return v.(*learn.Unit).CourseId }, "CourseId").
		Require(func(v interface{}) string { return v.(*learn.Unit).Name }, "Name").
		Enum(func(v interface{}) int32 { return int32(v.(*learn.Unit).Status) }, "Status", learn.ContentStatus_name).
		Build()
}
