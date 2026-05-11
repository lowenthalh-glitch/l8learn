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

const (
	ServiceName = "Engage"
	ServiceArea = byte(50)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	common.ActivateService(common.ServiceConfig{
		ServiceName: ServiceName,
		ServiceArea: ServiceArea,
		PrimaryKey:  "MetricId",
		Callback:    newEngagementServiceCallback(vnic),
	}, &learn.EngagementMetric{}, &learn.EngagementMetricList{}, creds, dbname, vnic)
}
