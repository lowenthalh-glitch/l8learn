/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package units

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "Unit"
	ServiceArea = byte(10)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	common.ActivateService(common.ServiceConfig{
		ServiceName: ServiceName,
		ServiceArea: ServiceArea,
		PrimaryKey:  "UnitId",
		Callback:    newUnitServiceCallback(vnic),
	}, &learn.Unit{}, &learn.UnitList{}, creds, dbname, vnic)
}

func Units(vnic ifs.IVNic) (ifs.IServiceHandler, bool) {
	return common.ServiceHandler(ServiceName, ServiceArea, vnic)
}

func Unit(unitId string, vnic ifs.IVNic) (*learn.Unit, error) {
	result, err := common.GetEntity(ServiceName, ServiceArea, &learn.Unit{UnitId: unitId}, vnic)
	if err != nil || result == nil {
		return nil, err
	}
	return result.(*learn.Unit), nil
}
