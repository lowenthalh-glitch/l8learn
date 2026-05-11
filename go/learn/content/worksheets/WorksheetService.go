/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package worksheets

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "Worksheet"
	ServiceArea = byte(10)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	common.ActivateService(common.ServiceConfig{
		ServiceName: ServiceName,
		ServiceArea: ServiceArea,
		PrimaryKey:  "WorksheetId",
		Callback:    newWorksheetServiceCallback(vnic),
	}, &learn.Worksheet{}, &learn.WorksheetList{}, creds, dbname, vnic)
}

func Worksheets(vnic ifs.IVNic) (ifs.IServiceHandler, bool) {
	return common.ServiceHandler(ServiceName, ServiceArea, vnic)
}

func Worksheet(worksheetId string, vnic ifs.IVNic) (*learn.Worksheet, error) {
	result, err := common.GetEntity(ServiceName, ServiceArea, &learn.Worksheet{WorksheetId: worksheetId}, vnic)
	if err != nil || result == nil {
		return nil, err
	}
	return result.(*learn.Worksheet), nil
}
