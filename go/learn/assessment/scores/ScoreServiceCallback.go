/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package scores

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newScoreServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.Score{}, vnic).
		BeforeAction(func(elem interface{}, action ifs.Action, vnic ifs.IVNic) error {
			if action == ifs.POST {
				common.GenerateID(&elem.(*learn.Score).ScoreId)
			}
			return nil
		}).
		Require(func(v interface{}) string { return v.(*learn.Score).ScoreId }, "ScoreId").
		Require(func(v interface{}) string { return v.(*learn.Score).StudentId }, "StudentId").
		Build()
}
