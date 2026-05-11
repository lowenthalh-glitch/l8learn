/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package skills

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "Skill"
	ServiceArea = byte(30)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	common.ActivateService(common.ServiceConfig{
		ServiceName: ServiceName,
		ServiceArea: ServiceArea,
		PrimaryKey:  "SkillId",
		Callback:    newSkillServiceCallback(vnic),
	}, &learn.Skill{}, &learn.SkillList{}, creds, dbname, vnic)
}

func Skills(vnic ifs.IVNic) (ifs.IServiceHandler, bool) {
	return common.ServiceHandler(ServiceName, ServiceArea, vnic)
}

func Skill(skillId string, vnic ifs.IVNic) (*learn.Skill, error) {
	result, err := common.GetEntity(ServiceName, ServiceArea, &learn.Skill{SkillId: skillId}, vnic)
	if err != nil || result == nil {
		return nil, err
	}
	return result.(*learn.Skill), nil
}
