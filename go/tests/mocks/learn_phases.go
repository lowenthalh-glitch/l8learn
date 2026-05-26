/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package mocks

import (
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/saichler/l8learn/go/types/learn"
)

//go:embed fixtures/l8learn_schedules.json
var schedulesFixture []byte

//go:embed fixtures/l8learn_lessons.json
var lessonsFixture []byte

// Client is the HTTP client interface for posting mock data
type Client interface {
	Post(endpoint string, data interface{}) error
	PostRawJSON(endpoint string, jsonData []byte) error
}

// RunWithMode executes mock data based on the selected mode.
func RunWithMode(client Client, mode string) {
	switch mode {
	case "security":
		runDemoData(client)
		RunSecurityTestData(client)
	case "full":
		runDemoData(client)
		RunSecurityTestData(client)
		// TODO: RunFullTestData(client) — profiles + evals + schedules for all students
	default: // "demo"
		runDemoData(client)
	}
}

// RunAllPhases is the legacy entry point — runs demo + security.
func RunAllPhases(client Client) {
	RunWithMode(client, "security")
}

func runDemoData(client Client) {
	store := NewMockDataStore()

	fmt.Println("=== Demo Data: Student + Guardian + Teacher + Profile ===")

	// Guardian (Luca's parent)
	guardian := &learn.Guardian{
		GuardianId: "GRD-0001",
		FirstName:  "Erica",
		LastName:   "Lowenthal",
		Email:      "erica@lowenthal.family",
		Phone:      "5551234567",
		Relation:   1, // Parent
		StudentIds: []string{"STU-0001"},
	}
	if err := client.Post("/learn/20/Guardian", &learn.GuardianList{List: []*learn.Guardian{guardian}}); err != nil {
		fmt.Printf("  ERROR Guardian: %v\n", err)
	}
	store.GuardianIDs = append(store.GuardianIDs, guardian.GuardianId)
	fmt.Printf("  Guardian: %s %s\n", guardian.FirstName, guardian.LastName)

	// Teacher
	teacher := &learn.Teacher{
		TeacherId: "TCH-0001",
		FirstName: "Sarah",
		LastName:  "Miller",
		Email:     "sarah.miller@school.edu",
		Role:      1, // Primary
	}
	if err := client.Post("/learn/20/Teacher", &learn.TeacherList{List: []*learn.Teacher{teacher}}); err != nil {
		fmt.Printf("  ERROR Teacher: %v\n", err)
	}
	store.TeacherIDs = append(store.TeacherIDs, teacher.TeacherId)
	fmt.Printf("  Teacher: %s %s\n", teacher.FirstName, teacher.LastName)

	// Student
	student := &learn.Student{
		StudentId:          "STU-0001",
		FirstName:          "Luca",
		LastName:           "Lowenthal",
		PreferredName:      "Luca",
		GradeLevel:         1,
		PrimaryGuardianId:  "GRD-0001",
		LanguagePreference: "en",
		HasIep:             true,
		AccommodationNotes: "Extended time on assessments, visual aids preferred",
	}
	if err := client.Post("/learn/20/Student", &learn.StudentList{List: []*learn.Student{student}}); err != nil {
		fmt.Printf("  ERROR Student: %v\n", err)
	}
	store.StudentIDs = append(store.StudentIDs, student.StudentId)
	fmt.Printf("  Student: %s %s\n", student.FirstName, student.LastName)

	// Profile
	profiles := generateProfiles(store)
	if err := client.Post("/learn/20/Profile", &learn.StudentProfileList{List: profiles}); err != nil {
		fmt.Printf("  ERROR Profile: %v\n", err)
	}
	fmt.Printf("  Profiles: %d\n", len(profiles))

	// User accounts
	fmt.Println("\n=== User Accounts ===")
	createUser(client, guardian.GuardianId, guardian.FirstName+" "+guardian.LastName, guardian.Email, "guardian", "guardian.html", guardian.StudentIds)
	createUser(client, teacher.TeacherId, teacher.FirstName+" "+teacher.LastName, teacher.Email, "teacher", "teacher.html", nil)
	studentEmail := fmt.Sprintf("%s.%s@student.l8learn.local", student.FirstName, student.LastName)
	createUser(client, student.StudentId, student.FirstName+" "+student.LastName, studentEmail, "student", "student.html", nil)

	// Upload LLM-generated fixtures (schedule + lessons) — embedded, no LLM call needed
	fmt.Println("\n=== LLM Fixture Data (pre-generated) ===")
	uploadFixture(client, "/learn/30/Schedule", schedulesFixture, "Schedules")
	uploadFixture(client, "/learn/10/GenLesson", lessonsFixture, "Lessons")

	fmt.Printf("\n=== Demo Summary ===\n")
	fmt.Printf("  Guardian:  1 (Erica)\n")
	fmt.Printf("  Teacher:   1 (Sarah)\n")
	fmt.Printf("  Student:   1 (Luca)\n")
	fmt.Printf("  Profiles:  %d\n", len(profiles))
	fmt.Printf("  Users:     3\n")
}

func createUser(client Client, userId, fullName, email, role, portal string, associateIds []string) {
	userData := map[string]interface{}{
		"userId":        userId,
		"fullName":      fullName,
		"email":         email,
		"portal":        portal,
		"password":      map[string]string{"hash": "12345678"},
		"accountStatus": "ACCOUNT_STATUS_ACTIVE",
		"roles":         map[string]bool{role: true},
	}
	if len(associateIds) > 0 {
		userData["associateIds"] = associateIds
	}
	if err := client.Post("/learn/73/users", userData); err != nil {
		fmt.Printf("  FAIL %s user (%s): %v\n", role, email, err)
	} else {
		fmt.Printf("  Created %s user: %s (%s)\n", role, fullName, email)
	}
}

func uploadFixture(client Client, endpoint string, data []byte, label string) {
	if len(data) == 0 {
		fmt.Printf("  %s: empty fixture data\n", label)
		return
	}
	if err := client.PostRawJSON(endpoint, data); err != nil {
		fmt.Printf("  %s: upload failed: %v\n", label, err)
	} else {
		var wrapper map[string]json.RawMessage
		json.Unmarshal(data, &wrapper)
		var items []json.RawMessage
		json.Unmarshal(wrapper["list"], &items)
		fmt.Printf("  %s: %d records uploaded\n", label, len(items))
	}
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
