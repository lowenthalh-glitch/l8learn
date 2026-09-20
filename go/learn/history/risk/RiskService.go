/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package risk

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "RiskAssmt"
	ServiceArea = byte(60)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	sla := common.NewOrmSLA(ServiceName, ServiceArea, "AssessmentId", newRiskServiceCallback(vnic),
		&learn.RiskAssessment{}, &learn.RiskAssessmentList{})
	common.ActivateService(sla, creds, dbname, vnic)
}
