/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package courses

import (
	"github.com/saichler/l8common/go/common"
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "Course"
	ServiceArea = byte(10)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	common.ActivateService(common.ServiceConfig{
		ServiceName: ServiceName,
		ServiceArea: ServiceArea,
		PrimaryKey:  "CourseId",
		Callback:    newCourseServiceCallback(vnic),
	}, &learn.Course{}, &learn.CourseList{}, creds, dbname, vnic)
}

func Courses(vnic ifs.IVNic) (ifs.IServiceHandler, bool) {
	return common.ServiceHandler(ServiceName, ServiceArea, vnic)
}

func Course(courseId string, vnic ifs.IVNic) (*learn.Course, error) {
	result, err := common.GetEntity(ServiceName, ServiceArea, &learn.Course{CourseId: courseId}, vnic)
	if err != nil || result == nil {
		return nil, err
	}
	return result.(*learn.Course), nil
}
