/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package guardians

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newGuardianServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.Guardian{}, vnic).
		BeforeAction(func(elem interface{}, action ifs.Action, vnic ifs.IVNic) error {
			if action == ifs.POST {
				common.GenerateID(&elem.(*learn.Guardian).GuardianId)
			}
			return nil
		}).
		Require(func(v interface{}) string { return v.(*learn.Guardian).GuardianId }, "GuardianId").
		Require(func(v interface{}) string { return v.(*learn.Guardian).FirstName }, "FirstName").
		Require(func(v interface{}) string { return v.(*learn.Guardian).LastName }, "LastName").
		Require(func(v interface{}) string { return v.(*learn.Guardian).Email }, "Email").
		Build()
}
