/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package engagement

import (
	"github.com/saichler/l8common/go/common"
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newEngagementServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.EngagementMetric{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.EngagementMetric).MetricId }, "MetricId").
		Require(func(v interface{}) string { return v.(*learn.EngagementMetric).StudentId }, "StudentId").
		Enum(func(v interface{}) int32 { return int32(v.(*learn.EngagementMetric).CurrentLevel) }, "CurrentLevel", learn.EngagementLevel_name).
		Build()
}
