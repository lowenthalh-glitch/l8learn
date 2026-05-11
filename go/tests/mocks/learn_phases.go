/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package mocks

import (
	"fmt"
	"github.com/saichler/l8learn/go/types/learn"
)

// Client is the HTTP client interface for posting mock data
type Client interface {
	Post(endpoint string, data interface{}) error
}

// RunAllPhases executes all mock data generation in dependency order
func RunAllPhases(client Client) {
	store := NewMockDataStore()

	fmt.Println("=== Phase 1: Foundation ===")
	runPhase1(client, store)

	fmt.Println("=== Phase 2: People ===")
	runPhase2(client, store)

	fmt.Println("=== Phase 3: Content ===")
	runPhase3(client, store)

	fmt.Println("=== Phase 4: Learning State ===")
	runPhase4(client, store)

	fmt.Printf("\n=== Summary ===\n")
	fmt.Printf("  Districts:   %d\n", len(store.DistrictIDs))
	fmt.Printf("  Schools:     %d\n", len(store.SchoolIDs))
	fmt.Printf("  Skills:      %d\n", len(store.SkillIDs))
	fmt.Printf("  Teachers:    %d\n", len(store.TeacherIDs))
	fmt.Printf("  Classrooms:  %d\n", len(store.ClassroomIDs))
	fmt.Printf("  Guardians:   %d\n", len(store.GuardianIDs))
	fmt.Printf("  Students:    %d\n", len(store.StudentIDs))
	fmt.Printf("  Courses:     %d\n", len(store.CourseIDs))
	fmt.Printf("  Units:       %d\n", len(store.UnitIDs))
	fmt.Printf("  Lessons:     %d\n", len(store.LessonIDs))
	fmt.Printf("  Activities:  %d\n", len(store.ActivityIDs))
	fmt.Printf("  Paths:       %d\n", len(store.PathIDs))
	fmt.Printf("  Mastery:     %d\n", len(store.MasteryIDs))
}

func runPhase1(client Client, store *MockDataStore) {
	districts := generateDistricts()
	client.Post("/learn/20/District", &learn.DistrictList{List: districts})
	for _, d := range districts {
		store.DistrictIDs = append(store.DistrictIDs, d.DistrictId)
	}
	fmt.Printf("  Districts: %d\n", len(districts))

	schools := generateSchools(store)
	client.Post("/learn/20/School", &learn.SchoolList{List: schools})
	for _, s := range schools {
		store.SchoolIDs = append(store.SchoolIDs, s.SchoolId)
	}
	fmt.Printf("  Schools: %d\n", len(schools))

	skills := generateSkills()
	client.Post("/learn/30/Skill", &learn.SkillList{List: skills})
	for _, s := range skills {
		store.SkillIDs = append(store.SkillIDs, s.SkillId)
	}
	fmt.Printf("  Skills: %d\n", len(skills))

	rules := generateAdaptRules()
	client.Post("/learn/30/AdaptRule", &learn.AdaptationRuleList{List: rules})
	for _, r := range rules {
		store.RuleIDs = append(store.RuleIDs, r.RuleId)
	}
	fmt.Printf("  Rules: %d\n", len(rules))
}

func runPhase2(client Client, store *MockDataStore) {
	teachers := generateTeachers(store)
	client.Post("/learn/20/Teacher", &learn.TeacherList{List: teachers})
	for _, t := range teachers {
		store.TeacherIDs = append(store.TeacherIDs, t.TeacherId)
	}
	fmt.Printf("  Teachers: %d\n", len(teachers))

	classrooms := generateClassrooms(store)
	client.Post("/learn/20/Classroom", &learn.ClassroomList{List: classrooms})
	for _, c := range classrooms {
		store.ClassroomIDs = append(store.ClassroomIDs, c.ClassroomId)
	}
	fmt.Printf("  Classrooms: %d\n", len(classrooms))

	guardians := generateGuardians(store)
	client.Post("/learn/20/Guardian", &learn.GuardianList{List: guardians})
	for _, g := range guardians {
		store.GuardianIDs = append(store.GuardianIDs, g.GuardianId)
	}
	fmt.Printf("  Guardians: %d\n", len(guardians))

	students := generateStudents(store)
	client.Post("/learn/20/Student", &learn.StudentList{List: students})
	for _, s := range students {
		store.StudentIDs = append(store.StudentIDs, s.StudentId)
	}
	fmt.Printf("  Students: %d\n", len(students))
}

func runPhase3(client Client, store *MockDataStore) {
	courses := generateCourses()
	client.Post("/learn/10/Course", &learn.CourseList{List: courses})
	for _, c := range courses {
		store.CourseIDs = append(store.CourseIDs, c.CourseId)
	}
	fmt.Printf("  Courses: %d\n", len(courses))

	units := generateUnits(store)
	client.Post("/learn/10/Unit", &learn.UnitList{List: units})
	for _, u := range units {
		store.UnitIDs = append(store.UnitIDs, u.UnitId)
	}
	fmt.Printf("  Units: %d\n", len(units))

	lessons := generateLessons(store)
	client.Post("/learn/10/Lesson", &learn.LessonList{List: lessons})
	for _, l := range lessons {
		store.LessonIDs = append(store.LessonIDs, l.LessonId)
	}
	fmt.Printf("  Lessons: %d\n", len(lessons))

	activities := generateActivities(store)
	client.Post("/learn/10/Activity", &learn.ActivityList{List: activities})
	for _, a := range activities {
		store.ActivityIDs = append(store.ActivityIDs, a.ActivityId)
	}
	fmt.Printf("  Activities: %d\n", len(activities))
}

func runPhase4(client Client, store *MockDataStore) {
	paths := generatePaths(store)
	client.Post("/learn/30/LearnPath", &learn.LearningPathList{List: paths})
	for _, p := range paths {
		store.PathIDs = append(store.PathIDs, p.PathId)
	}
	fmt.Printf("  Paths: %d\n", len(paths))

	mastery := generateMastery(store)
	client.Post("/learn/30/Mastery", &learn.SkillMasteryList{List: mastery})
	for _, m := range mastery {
		store.MasteryIDs = append(store.MasteryIDs, m.MasteryId)
	}
	fmt.Printf("  Mastery: %d\n", len(mastery))
}
