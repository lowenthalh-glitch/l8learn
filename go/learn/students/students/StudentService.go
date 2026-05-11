/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package students

import (
	"github.com/saichler/l8common/go/common"
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "Student"
	ServiceArea = byte(20)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	common.ActivateService(common.ServiceConfig{
		ServiceName: ServiceName,
		ServiceArea: ServiceArea,
		PrimaryKey:  "StudentId",
		Callback:    newStudentServiceCallback(vnic),
	}, &learn.Student{}, &learn.StudentList{}, creds, dbname, vnic)
}

func Students(vnic ifs.IVNic) (ifs.IServiceHandler, bool) {
	return common.ServiceHandler(ServiceName, ServiceArea, vnic)
}

func Student(studentId string, vnic ifs.IVNic) (*learn.Student, error) {
	result, err := common.GetEntity(ServiceName, ServiceArea, &learn.Student{StudentId: studentId}, vnic)
	if err != nil || result == nil {
		return nil, err
	}
	return result.(*learn.Student), nil
}
