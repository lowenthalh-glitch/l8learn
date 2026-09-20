/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package genlessons

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "GenLesson"
	ServiceArea = byte(10)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	sla := common.NewOrmSLA(ServiceName, ServiceArea, "GeneratedLessonId", newGenLessonServiceCallback(vnic),
		&learn.GeneratedLesson{}, &learn.GeneratedLessonList{})
	common.ActivateService(sla, creds, dbname, vnic)
}
