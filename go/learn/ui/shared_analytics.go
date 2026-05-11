/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package ui

import (
	l8c "github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func registerAnalyticsTypes(resources ifs.IResources) {
	
	l8c.RegisterType(resources, &learn.ProgressReport{}, &learn.ProgressReportList{}, "ReportId")
	l8c.RegisterType(resources, &learn.EngagementMetric{}, &learn.EngagementMetricList{}, "MetricId")
}
