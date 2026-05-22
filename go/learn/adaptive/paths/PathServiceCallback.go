/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package paths

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newPathServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.LearningPath{}, vnic).
		BeforeAction(func(elem interface{}, action ifs.Action, vnic ifs.IVNic) error {
			if action == ifs.POST {
				common.GenerateID(&elem.(*learn.LearningPath).PathId)
			}
			return nil
		}).
		Require(func(v interface{}) string { return v.(*learn.LearningPath).PathId }, "PathId").
		Require(func(v interface{}) string { return v.(*learn.LearningPath).StudentId }, "StudentId").
		After(onPathChange).
		Build()
}

func onPathChange(elem interface{}, action ifs.Action, vnic ifs.IVNic) error {
	path := elem.(*learn.LearningPath)
	if action == ifs.POST {
		// 1. Schedule initial diagnostic benchmark for this subject
		// 2. Populate upcoming_queue with first activities (after diagnostic)
		_ = path
	}
	return nil
}
