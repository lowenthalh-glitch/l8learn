/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package engagement

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newEngagementServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.EngagementMetric{}, vnic).
		BeforeAction(func(elem interface{}, action ifs.Action, vnic ifs.IVNic) error {
			if action == ifs.POST {
				common.GenerateID(&elem.(*learn.EngagementMetric).MetricId)
			}
			return nil
		}).
		Require(func(v interface{}) string { return v.(*learn.EngagementMetric).MetricId }, "MetricId").
		Require(func(v interface{}) string { return v.(*learn.EngagementMetric).StudentId }, "StudentId").
		Build()
}
