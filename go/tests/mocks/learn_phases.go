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

	// Create one student and one full-depth profile
	fmt.Println("=== Demo Student + Profile ===")

	student := &learn.Student{
		StudentId:          "STU-0001",
		FirstName:          "Luca",
		LastName:           "Lowenthal",
		PreferredName:      "Luca",
		GradeLevel:         1,
		SchoolId:           "SCH-0001",
		DistrictId:         "DIST-0001",
		LanguagePreference: "en",
		HasIep:             true,
		AccommodationNotes: "Extended time on assessments, visual aids preferred",
	}
	if err := client.Post("/learn/20/Student", &learn.StudentList{List: []*learn.Student{student}}); err != nil {
		fmt.Printf("  ERROR Student: %v\n", err)
	}
	fmt.Println("  Student: Jake Martinez (STU-0001)")

	store.StudentIDs = append(store.StudentIDs, student.StudentId)

	profiles := generateProfiles(store)
	if err := client.Post("/learn/20/Profile", &learn.StudentProfileList{List: profiles}); err != nil {
		fmt.Printf("  ERROR Profile: %v\n", err)
	}
	fmt.Printf("  Profiles: %d\n", len(profiles))

	fmt.Printf("\n=== Summary ===\n")
	fmt.Printf("  Student: 1\n")
	fmt.Printf("  Profiles: %d\n", len(profiles))
}

/*
func RunAllPhasesFullData(client Client) {
	store := NewMockDataStore()

	fmt.Println("=== Phase 1: Foundation ===")
	runPhase1(client, store)

	fmt.Println("=== Phase 2: People ===")
	runPhase2(client, store)

	fmt.Println("=== Phase 3: Content ===")
	runPhase3(client, store)

	fmt.Println("=== Phase 4: Learning State ===")
	runPhase4(client, store)

	fmt.Println("=== Phase 5: Adaptive Intelligence ===")
	runPhase5(client, store)

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
	fmt.Printf("  Profiles:    %d\n", len(store.ProfileIDs))
	fmt.Printf("  PromptLogs:  %d\n", len(store.PromptLogIDs))
	fmt.Printf("  EvalImports: %d\n", len(store.EvalImportIDs))
	fmt.Printf("  GenLessons:  %d\n", len(store.GeneratedLessonIDs))
}
*/

func runPhase1(client Client, store *MockDataStore) {
	districts := generateDistricts()
	if err := client.Post("/learn/20/District", &learn.DistrictList{List: districts}); err != nil { fmt.Printf("  ERROR: %v\n", err) }
	for _, d := range districts {
		store.DistrictIDs = append(store.DistrictIDs, d.DistrictId)
	}
	fmt.Printf("  Districts: %d\n", len(districts))

	schools := generateSchools(store)
	if err := client.Post("/learn/20/School", &learn.SchoolList{List: schools}); err != nil { fmt.Printf("  ERROR: %v\n", err) }
	for _, s := range schools {
		store.SchoolIDs = append(store.SchoolIDs, s.SchoolId)
	}
	fmt.Printf("  Schools: %d\n", len(schools))

	skills := generateSkills()
	if err := client.Post("/learn/30/Skill", &learn.SkillList{List: skills}); err != nil { fmt.Printf("  ERROR: %v\n", err) }
	for _, s := range skills {
		store.SkillIDs = append(store.SkillIDs, s.SkillId)
	}
	fmt.Printf("  Skills: %d\n", len(skills))

	rules := generateAdaptRules()
	if err := client.Post("/learn/30/AdaptRule", &learn.AdaptationRuleList{List: rules}); err != nil { fmt.Printf("  ERROR: %v\n", err) }
	for _, r := range rules {
		store.RuleIDs = append(store.RuleIDs, r.RuleId)
	}
	fmt.Printf("  Rules: %d\n", len(rules))
}

func runPhase2(client Client, store *MockDataStore) {
	teachers := generateTeachers(store)
	if err := client.Post("/learn/20/Teacher", &learn.TeacherList{List: teachers}); err != nil { fmt.Printf("  ERROR: %v\n", err) }
	for _, t := range teachers {
		store.TeacherIDs = append(store.TeacherIDs, t.TeacherId)
	}
	fmt.Printf("  Teachers: %d\n", len(teachers))

	classrooms := generateClassrooms(store)
	if err := client.Post("/learn/20/Classroom", &learn.ClassroomList{List: classrooms}); err != nil { fmt.Printf("  ERROR: %v\n", err) }
	for _, c := range classrooms {
		store.ClassroomIDs = append(store.ClassroomIDs, c.ClassroomId)
	}
	fmt.Printf("  Classrooms: %d\n", len(classrooms))

	guardians := generateGuardians(store)
	if err := client.Post("/learn/20/Guardian", &learn.GuardianList{List: guardians}); err != nil { fmt.Printf("  ERROR: %v\n", err) }
	for _, g := range guardians {
		store.GuardianIDs = append(store.GuardianIDs, g.GuardianId)
	}
	fmt.Printf("  Guardians: %d\n", len(guardians))

	students := generateStudents(store)
	if err := client.Post("/learn/20/Student", &learn.StudentList{List: students}); err != nil { fmt.Printf("  ERROR: %v\n", err) }
	for _, s := range students {
		store.StudentIDs = append(store.StudentIDs, s.StudentId)
	}
	fmt.Printf("  Students: %d\n", len(students))
}

func runPhase3(client Client, store *MockDataStore) {
	courses := generateCourses()
	if err := client.Post("/learn/10/Course", &learn.CourseList{List: courses}); err != nil { fmt.Printf("  ERROR: %v\n", err) }
	for _, c := range courses {
		store.CourseIDs = append(store.CourseIDs, c.CourseId)
	}
	fmt.Printf("  Courses: %d\n", len(courses))

	units := generateUnits(store)
	if err := client.Post("/learn/10/Unit", &learn.UnitList{List: units}); err != nil { fmt.Printf("  ERROR: %v\n", err) }
	for _, u := range units {
		store.UnitIDs = append(store.UnitIDs, u.UnitId)
	}
	fmt.Printf("  Units: %d\n", len(units))

	lessons := generateLessons(store)
	if err := client.Post("/learn/10/Lesson", &learn.LessonList{List: lessons}); err != nil { fmt.Printf("  ERROR: %v\n", err) }
	for _, l := range lessons {
		store.LessonIDs = append(store.LessonIDs, l.LessonId)
	}
	fmt.Printf("  Lessons: %d\n", len(lessons))

	activities := generateActivities(store)
	if err := client.Post("/learn/10/Activity", &learn.ActivityList{List: activities}); err != nil { fmt.Printf("  ERROR: %v\n", err) }
	for _, a := range activities {
		store.ActivityIDs = append(store.ActivityIDs, a.ActivityId)
	}
	fmt.Printf("  Activities: %d\n", len(activities))
}

func runPhase4(client Client, store *MockDataStore) {
	paths := generatePaths(store)
	if err := client.Post("/learn/30/LearnPath", &learn.LearningPathList{List: paths}); err != nil { fmt.Printf("  ERROR: %v\n", err) }
	for _, p := range paths {
		store.PathIDs = append(store.PathIDs, p.PathId)
	}
	fmt.Printf("  Paths: %d\n", len(paths))

	mastery := generateMastery(store)
	if err := client.Post("/learn/30/Mastery", &learn.SkillMasteryList{List: mastery}); err != nil { fmt.Printf("  ERROR: %v\n", err) }
	for _, m := range mastery {
		store.MasteryIDs = append(store.MasteryIDs, m.MasteryId)
	}
	fmt.Printf("  Mastery: %d\n", len(mastery))
}

func runPhase5(client Client, store *MockDataStore) {
	// Student Profiles (depends on StudentIDs)
	profiles := generateProfiles(store)
	if err := client.Post("/learn/20/Profile", &learn.StudentProfileList{List: profiles}); err != nil { fmt.Printf("  ERROR: %v\n", err) }
	for _, p := range profiles {
		store.ProfileIDs = append(store.ProfileIDs, p.ProfileId)
	}
	fmt.Printf("  Profiles: %d\n", len(profiles))

	// LLM Config (singleton)
	config := generateLLMConfig()
	if err := client.Post("/learn/30/LLMConfig", &learn.LLMConfigList{List: []*learn.LLMConfig{config}}); err != nil { fmt.Printf("  ERROR: %v\n", err) }
	fmt.Printf("  LLMConfig: 1\n")

	// Prompt Logs (depends on StudentIDs)
	logs := generatePromptLogs(store)
	if err := client.Post("/learn/30/PromptLog", &learn.LLMPromptLogList{List: logs}); err != nil { fmt.Printf("  ERROR: %v\n", err) }
	for _, l := range logs {
		store.PromptLogIDs = append(store.PromptLogIDs, l.LogId)
	}
	fmt.Printf("  PromptLogs: %d\n", len(logs))

	// Eval Imports (depends on StudentIDs, GuardianIDs)
	evals := generateEvalImports(store)
	if err := client.Post("/learn/20/EvalImprt", &learn.EvalImportList{List: evals}); err != nil { fmt.Printf("  ERROR: %v\n", err) }
	for _, e := range evals {
		store.EvalImportIDs = append(store.EvalImportIDs, e.ImportId)
	}
	fmt.Printf("  EvalImports: %d\n", len(evals))

	// Generated Lessons (depends on StudentIDs, PathIDs, SkillIDs)
	genLessons := generateGeneratedLessons(store)
	if err := client.Post("/learn/10/GenLesson", &learn.GeneratedLessonList{List: genLessons}); err != nil { fmt.Printf("  ERROR: %v\n", err) }
	for _, gl := range genLessons {
		store.GeneratedLessonIDs = append(store.GeneratedLessonIDs, gl.GeneratedLessonId)
	}
	fmt.Printf("  GenLessons: %d\n", len(genLessons))
}
