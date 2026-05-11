/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package activities

import (
	"github.com/saichler/l8common/go/common"
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "Activity"
	ServiceArea = byte(10)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	common.ActivateService(common.ServiceConfig{
		ServiceName: ServiceName,
		ServiceArea: ServiceArea,
		PrimaryKey:  "ActivityId",
		Callback:    newActivityServiceCallback(vnic),
	}, &learn.Activity{}, &learn.ActivityList{}, creds, dbname, vnic)
}

func Activities(vnic ifs.IVNic) (ifs.IServiceHandler, bool) {
	return common.ServiceHandler(ServiceName, ServiceArea, vnic)
}

func Activity(activityId string, vnic ifs.IVNic) (*learn.Activity, error) {
	result, err := common.GetEntity(ServiceName, ServiceArea, &learn.Activity{ActivityId: activityId}, vnic)
	if err != nil || result == nil {
		return nil, err
	}
	return result.(*learn.Activity), nil
}
