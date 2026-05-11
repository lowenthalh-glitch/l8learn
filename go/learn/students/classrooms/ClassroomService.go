/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package classrooms

import (
	"github.com/saichler/l8common/go/common"
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "Classroom"
	ServiceArea = byte(20)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	common.ActivateService(common.ServiceConfig{
		ServiceName: ServiceName,
		ServiceArea: ServiceArea,
		PrimaryKey:  "ClassroomId",
		Callback:    newClassroomServiceCallback(vnic),
	}, &learn.Classroom{}, &learn.ClassroomList{}, creds, dbname, vnic)
}
