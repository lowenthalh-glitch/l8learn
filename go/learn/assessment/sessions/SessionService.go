/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package sessions

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "LearnSess"
	ServiceArea = byte(40)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	common.ActivateService(common.ServiceConfig{
		ServiceName: ServiceName,
		ServiceArea: ServiceArea,
		PrimaryKey:  "SessionId",
		Callback:    newSessionServiceCallback(vnic),
	}, &learn.LearningSession{}, &learn.LearningSessionList{}, creds, dbname, vnic)
}

func Sessions(vnic ifs.IVNic) (ifs.IServiceHandler, bool) {
	return common.ServiceHandler(ServiceName, ServiceArea, vnic)
}

func Session(sessionId string, vnic ifs.IVNic) (*learn.LearningSession, error) {
	result, err := common.GetEntity(ServiceName, ServiceArea, &learn.LearningSession{SessionId: sessionId}, vnic)
	if err != nil || result == nil {
		return nil, err
	}
	return result.(*learn.LearningSession), nil
}
