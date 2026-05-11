/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package ui

import (
	l8c "github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func registerContentTypes(resources ifs.IResources) {
	
	l8c.RegisterType(resources, &learn.Course{}, &learn.CourseList{}, "CourseId")
	l8c.RegisterType(resources, &learn.Unit{}, &learn.UnitList{}, "UnitId")
	l8c.RegisterType(resources, &learn.Lesson{}, &learn.LessonList{}, "LessonId")
	l8c.RegisterType(resources, &learn.Activity{}, &learn.ActivityList{}, "ActivityId")
	l8c.RegisterType(resources, &learn.Worksheet{}, &learn.WorksheetList{}, "WorksheetId")
	l8c.RegisterType(resources, &learn.FamilyActivity{}, &learn.FamilyActivityList{}, "FamilyActivityId")
	l8c.RegisterType(resources, &learn.RealWorldLesson{}, &learn.RealWorldLessonList{}, "LessonId")
	l8c.RegisterType(resources, &learn.Project{}, &learn.ProjectList{}, "ProjectId")
	l8c.RegisterType(resources, &learn.GeneratedLesson{}, &learn.GeneratedLessonList{}, "GeneratedLessonId")
}
