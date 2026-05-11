/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package districts

import (
	"github.com/saichler/l8common/go/common"
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newDistrictServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.District{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.District).DistrictId }, "DistrictId").
		Require(func(v interface{}) string { return v.(*learn.District).Name }, "Name").
		Build()
}
