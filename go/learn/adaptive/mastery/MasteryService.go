/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package mastery

import (
	"github.com/saichler/l8common/go/common"
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "Mastery"
	ServiceArea = byte(30)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	common.ActivateService(common.ServiceConfig{
		ServiceName: ServiceName,
		ServiceArea: ServiceArea,
		PrimaryKey:  "MasteryId",
		Callback:    newMasteryServiceCallback(vnic),
	}, &learn.SkillMastery{}, &learn.SkillMasteryList{}, creds, dbname, vnic)
}

func Masteries(vnic ifs.IVNic) (ifs.IServiceHandler, bool) {
	return common.ServiceHandler(ServiceName, ServiceArea, vnic)
}

func Mastery(masteryId string, vnic ifs.IVNic) (*learn.SkillMastery, error) {
	result, err := common.GetEntity(ServiceName, ServiceArea, &learn.SkillMastery{MasteryId: masteryId}, vnic)
	if err != nil || result == nil {
		return nil, err
	}
	return result.(*learn.SkillMastery), nil
}
