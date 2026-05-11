/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package families

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newFamilyServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.Family{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.Family).FamilyId }, "FamilyId").
		Require(func(v interface{}) string { return v.(*learn.Family).Name }, "Name").
		Require(func(v interface{}) string { return v.(*learn.Family).PrimaryGuardianId }, "PrimaryGuardianId").
		After(onFamilyCreated).
		Build()
}

func onFamilyCreated(elem interface{}, action ifs.Action, notify bool, vnic ifs.IVNic) (interface{}, bool, error) {
	family := elem.(*learn.Family)
	if action == ifs.POST {
		// 1. Link existing guardians and students by ID
		// 2. Set up state compliance record for the family's state
		_ = family
	}
	return nil, true, nil
}
