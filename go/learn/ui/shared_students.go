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

func registerStudentTypes(resources ifs.IResources) {
	
	l8c.RegisterType(resources, &learn.Student{}, &learn.StudentList{}, "StudentId")
	l8c.RegisterType(resources, &learn.Guardian{}, &learn.GuardianList{}, "GuardianId")
	l8c.RegisterType(resources, &learn.Teacher{}, &learn.TeacherList{}, "TeacherId")
	l8c.RegisterType(resources, &learn.Classroom{}, &learn.ClassroomList{}, "ClassroomId")
	l8c.RegisterType(resources, &learn.School{}, &learn.SchoolList{}, "SchoolId")
	l8c.RegisterType(resources, &learn.District{}, &learn.DistrictList{}, "DistrictId")
	l8c.RegisterType(resources, &learn.Enrollment{}, &learn.EnrollmentList{}, "EnrollmentId")
	l8c.RegisterType(resources, &learn.Family{}, &learn.FamilyList{}, "FamilyId")
	l8c.RegisterType(resources, &learn.StateCompliance{}, &learn.StateComplianceList{}, "ComplianceId")
	l8c.RegisterType(resources, &learn.LearningPod{}, &learn.LearningPodList{}, "PodId")
	l8c.RegisterType(resources, &learn.StudentProfile{}, &learn.StudentProfileList{}, "ProfileId")
	l8c.RegisterType(resources, &learn.EvalImport{}, &learn.EvalImportList{}, "ImportId")
}
