/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package lessons

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "Lesson"
	ServiceArea = byte(10)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	common.ActivateService(common.ServiceConfig{
		ServiceName: ServiceName,
		ServiceArea: ServiceArea,
		PrimaryKey:  "LessonId",
		Callback:    newLessonServiceCallback(vnic),
	}, &learn.Lesson{}, &learn.LessonList{}, creds, dbname, vnic)
}

func Lessons(vnic ifs.IVNic) (ifs.IServiceHandler, bool) {
	return common.ServiceHandler(ServiceName, ServiceArea, vnic)
}

func Lesson(lessonId string, vnic ifs.IVNic) (*learn.Lesson, error) {
	result, err := common.GetEntity(ServiceName, ServiceArea, &learn.Lesson{LessonId: lessonId}, vnic)
	if err != nil || result == nil {
		return nil, err
	}
	return result.(*learn.Lesson), nil
}
