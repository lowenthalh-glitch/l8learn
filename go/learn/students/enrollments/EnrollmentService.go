/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package enrollments

import (
	"github.com/saichler/l8common/go/common"
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "Enroll"
	ServiceArea = byte(20)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	common.ActivateService(common.ServiceConfig{
		ServiceName: ServiceName,
		ServiceArea: ServiceArea,
		PrimaryKey:  "EnrollmentId",
		Callback:    newEnrollmentServiceCallback(vnic),
	}, &learn.Enrollment{}, &learn.EnrollmentList{}, creds, dbname, vnic)
}

func Enrollments(vnic ifs.IVNic) (ifs.IServiceHandler, bool) {
	return common.ServiceHandler(ServiceName, ServiceArea, vnic)
}

func Enrollment(enrollmentId string, vnic ifs.IVNic) (*learn.Enrollment, error) {
	result, err := common.GetEntity(ServiceName, ServiceArea, &learn.Enrollment{EnrollmentId: enrollmentId}, vnic)
	if err != nil || result == nil {
		return nil, err
	}
	return result.(*learn.Enrollment), nil
}
