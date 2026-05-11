/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package mocks

import (
	"fmt"
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
	"math/rand"
	"time"
)

func generateTeachers(store *MockDataStore) []*learn.Teacher {
	var list []*learn.Teacher
	for i := 0; i < 80; i++ {
		first := TeacherFirstNames[i%len(TeacherFirstNames)]
		last := LastNames[i%len(LastNames)]
		list = append(list, &learn.Teacher{
			TeacherId:  fmt.Sprintf("TCH-%03d", i+1),
			FirstName:  first,
			LastName:   last,
			Email:      fmt.Sprintf("%s.%s@school.edu", first, last),
			Role:       learn.TeacherRole_TEACHER_ROLE_PRIMARY,
			SchoolId:   pickRef(store.SchoolIDs, i/4),
			DistrictId: pickRef(store.DistrictIDs, i/16),
		})
	}
	return list
}

func generateClassrooms(store *MockDataStore) []*learn.Classroom {
	var list []*learn.Classroom
	grades := []learn.GradeLevel{
		learn.GradeLevel_GRADE_LEVEL_K, learn.GradeLevel_GRADE_LEVEL_1,
		learn.GradeLevel_GRADE_LEVEL_2, learn.GradeLevel_GRADE_LEVEL_3,
		learn.GradeLevel_GRADE_LEVEL_4,
	}
	for i := 0; i < 100; i++ {
		list = append(list, &learn.Classroom{
			ClassroomId:      fmt.Sprintf("CLS-%03d", i+1),
			Name:             fmt.Sprintf("Room %d", 100+i),
			GradeLevel:       grades[i%len(grades)],
			PrimaryTeacherId: pickRef(store.TeacherIDs, i),
			SchoolId:         pickRef(store.SchoolIDs, i/5),
			AcademicYear:     "2026-2027",
			StudentCount:     10,
		})
	}
	return list
}

func generateGuardians(store *MockDataStore) []*learn.Guardian {
	var list []*learn.Guardian
	for i := 0; i < 500; i++ {
		first := FirstNames[(i+10)%len(FirstNames)]
		last := LastNames[i%len(LastNames)]
		list = append(list, &learn.Guardian{
			GuardianId:       fmt.Sprintf("GRD-%03d", i+1),
			FirstName:        first,
			LastName:         last,
			Email:            fmt.Sprintf("%s.%s.parent@email.com", first, last),
			Phone:            fmt.Sprintf("555-%03d-%04d", i/100, i%10000),
			Relation:         learn.GuardianRelation_GUARDIAN_RELATION_PARENT,
			ReceivesReports:  true,
			ReportFrequency:  "weekly",
		})
	}
	return list
}

func generateStudents(store *MockDataStore) []*learn.Student {
	var list []*learn.Student
	grades := []learn.GradeLevel{
		learn.GradeLevel_GRADE_LEVEL_K, learn.GradeLevel_GRADE_LEVEL_1,
		learn.GradeLevel_GRADE_LEVEL_2, learn.GradeLevel_GRADE_LEVEL_3,
		learn.GradeLevel_GRADE_LEVEL_4, learn.GradeLevel_GRADE_LEVEL_5,
	}
	now := time.Now().Unix()

	for i := 0; i < 1000; i++ {
		first := FirstNames[i%len(FirstNames)]
		last := LastNames[(i/25)%len(LastNames)]
		grade := grades[i%len(grades)]
		list = append(list, &learn.Student{
			StudentId:          fmt.Sprintf("STU-%04d", i+1),
			FirstName:          first,
			LastName:           last,
			GradeLevel:         grade,
			Status:             learn.StudentStatus_STUDENT_STATUS_ACTIVE,
			ClassroomId:        pickRef(store.ClassroomIDs, i/10),
			SchoolId:           pickRef(store.SchoolIDs, i/50),
			DistrictId:         pickRef(store.DistrictIDs, i/200),
			PrimaryGuardianId:  pickRef(store.GuardianIDs, i/2),
			LanguagePreference: "en",
			WeeklyGoalMinutes:  int32(60 + rand.Intn(60)),
			DailyGoalMinutes:   int32(15 + rand.Intn(15)),
			EnrollmentDate:     now - int64(rand.Intn(365*24*3600)),
		})
	}
	return list
}
