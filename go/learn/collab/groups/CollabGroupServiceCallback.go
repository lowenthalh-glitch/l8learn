/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package groups

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newCollabGroupServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.CollabGroup{}, vnic).
		BeforeAction(func(elem interface{}, action ifs.Action, vnic ifs.IVNic) error {
			if action == ifs.POST {
				common.GenerateID(&elem.(*learn.CollabGroup).GroupId)
			}
			return nil
		}).
		Require(func(v interface{}) string { return v.(*learn.CollabGroup).GroupId }, "GroupId").
		Require(func(v interface{}) string { return v.(*learn.CollabGroup).Name }, "Name").
		Build()
}

// NOTE: Scheduled daily job (outside of callback) will:
// 1. If team member inactive 2+ days, send gentle nudge via l8notify
// 2. If whole team inactive 3+ days, notify teacher
// 3. Update team streak counter
