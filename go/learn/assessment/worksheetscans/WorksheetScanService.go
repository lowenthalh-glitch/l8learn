/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package worksheetscans

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "WkshtScan"
	ServiceArea = byte(40)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	sla := common.NewOrmSLA(ServiceName, ServiceArea, "ScanId", newWorksheetScanServiceCallback(vnic),
		&learn.WorksheetScan{}, &learn.WorksheetScanList{})
	common.ActivateService(sla, creds, dbname, vnic)
}

func WorksheetScans(vnic ifs.IVNic) (ifs.IServiceHandler, bool) {
	return common.ServiceHandler(ServiceName, ServiceArea, vnic)
}
