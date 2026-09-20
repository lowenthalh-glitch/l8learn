/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package compliance

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "Comply"
	ServiceArea = byte(20)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	sla := common.NewOrmSLA(ServiceName, ServiceArea, "ComplianceId", newComplianceServiceCallback(vnic),
		&learn.StateCompliance{}, &learn.StateComplianceList{})
	common.ActivateService(sla, creds, dbname, vnic)
}
