/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package paths

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "LearnPath"
	ServiceArea = byte(30)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	common.ActivateService(common.ServiceConfig{
		ServiceName: ServiceName,
		ServiceArea: ServiceArea,
		PrimaryKey:  "PathId",
		Callback:    newPathServiceCallback(vnic),
	}, &learn.LearningPath{}, &learn.LearningPathList{}, creds, dbname, vnic)
}

func Paths(vnic ifs.IVNic) (ifs.IServiceHandler, bool) {
	return common.ServiceHandler(ServiceName, ServiceArea, vnic)
}

func Path(pathId string, vnic ifs.IVNic) (*learn.LearningPath, error) {
	result, err := common.GetEntity(ServiceName, ServiceArea, &learn.LearningPath{PathId: pathId}, vnic)
	if err != nil || result == nil {
		return nil, err
	}
	return result.(*learn.LearningPath), nil
}
