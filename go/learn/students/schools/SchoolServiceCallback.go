/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package schools

import (
	"github.com/saichler/l8common/go/common"
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newSchoolServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.School{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.School).SchoolId }, "SchoolId").
		Require(func(v interface{}) string { return v.(*learn.School).Name }, "Name").
		Require(func(v interface{}) string { return v.(*learn.School).DistrictId }, "DistrictId").
		Build()
}
