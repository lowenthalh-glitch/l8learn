/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * Security test mock data — 5 extra students (including 2 brothers),
 * 3 guardians, and corresponding user accounts.
 * Used only for testing role-based security rules.
 * Call RunSecurityTestData(client) after RunAllPhases(client).
 */
package mocks

import (
	"fmt"
	"github.com/saichler/l8learn/go/types/learn"
)

// RunSecurityTestData creates test data for security role verification.
// 5 students, 3 guardians (one has 2 kids — brothers), user accounts for all.
func RunSecurityTestData(client Client) {
	fmt.Println("\n=== Security Test Data ===")

	// Guardian 1: Maria Garcia — has 2 sons (brothers)
	g1 := &learn.Guardian{
		GuardianId: "GRD-SEC-001",
		FirstName:  "Maria",
		LastName:   "Garcia",
		Email:      "maria.garcia@family.test",
		Phone:      "5551001001",
		Relation:   1,
		StudentIds: []string{"STU-SEC-001", "STU-SEC-002"},
	}

	// Guardian 2: James Wilson — has 1 daughter
	g2 := &learn.Guardian{
		GuardianId: "GRD-SEC-002",
		FirstName:  "James",
		LastName:   "Wilson",
		Email:      "james.wilson@family.test",
		Phone:      "5551002002",
		Relation:   1,
		StudentIds: []string{"STU-SEC-003"},
	}

	// Guardian 3: Aisha Khan — has 2 daughters
	g3 := &learn.Guardian{
		GuardianId: "GRD-SEC-003",
		FirstName:  "Aisha",
		LastName:   "Khan",
		Email:      "aisha.khan@family.test",
		Phone:      "5551003003",
		Relation:   1,
		StudentIds: []string{"STU-SEC-004", "STU-SEC-005"},
	}

	// Post guardians
	if err := client.Post("/learn/20/Guardian", &learn.GuardianList{List: []*learn.Guardian{g1, g2, g3}}); err != nil {
		fmt.Printf("  ERROR guardians: %v\n", err)
	}
	fmt.Printf("  Guardians: 3 (Maria Garcia, James Wilson, Aisha Khan)\n")

	// Student 1: Diego Garcia (age 8, Maria's older son)
	s1 := &learn.Student{
		StudentId:         "STU-SEC-001",
		FirstName:         "Diego",
		LastName:          "Garcia",
		GradeLevel:        5, // Grade 3
		PrimaryGuardianId: "GRD-SEC-001",
	}

	// Student 2: Marco Garcia (age 6, Maria's younger son — brother of Diego)
	s2 := &learn.Student{
		StudentId:         "STU-SEC-002",
		FirstName:         "Marco",
		LastName:          "Garcia",
		GradeLevel:        3, // Grade 1
		PrimaryGuardianId: "GRD-SEC-001",
	}

	// Student 3: Emma Wilson (age 7, James's daughter)
	s3 := &learn.Student{
		StudentId:         "STU-SEC-003",
		FirstName:         "Emma",
		LastName:          "Wilson",
		GradeLevel:        4, // Grade 2
		PrimaryGuardianId: "GRD-SEC-002",
	}

	// Student 4: Zara Khan (age 9, Aisha's older daughter)
	s4 := &learn.Student{
		StudentId:         "STU-SEC-004",
		FirstName:         "Zara",
		LastName:          "Khan",
		GradeLevel:        6, // Grade 4
		PrimaryGuardianId: "GRD-SEC-003",
	}

	// Student 5: Layla Khan (age 5, Aisha's younger daughter)
	s5 := &learn.Student{
		StudentId:         "STU-SEC-005",
		FirstName:         "Layla",
		LastName:          "Khan",
		GradeLevel:        2, // Kindergarten
		PrimaryGuardianId: "GRD-SEC-003",
	}

	// Post students
	if err := client.Post("/learn/20/Student", &learn.StudentList{List: []*learn.Student{s1, s2, s3, s4, s5}}); err != nil {
		fmt.Printf("  ERROR students: %v\n", err)
	}
	fmt.Printf("  Students: 5 (Diego+Marco Garcia, Emma Wilson, Zara+Layla Khan)\n")

	// Create user accounts
	fmt.Println("  Creating user accounts...")

	// Guardian users (userId = guardianId for deny-scope to work)
	createUser(client, g1.GuardianId, g1.FirstName+" "+g1.LastName, g1.Email, "guardian", "app.html")
	createUser(client, g2.GuardianId, g2.FirstName+" "+g2.LastName, g2.Email, "guardian", "app.html")
	createUser(client, g3.GuardianId, g3.FirstName+" "+g3.LastName, g3.Email, "guardian", "app.html")

	// Student users (userId = studentId for deny-scope to work)
	createUser(client, s1.StudentId, s1.FirstName+" "+s1.LastName, fmt.Sprintf("%s.%s@student.l8learn.local", s1.FirstName, s1.LastName), "student", "app.html")
	createUser(client, s2.StudentId, s2.FirstName+" "+s2.LastName, fmt.Sprintf("%s.%s@student.l8learn.local", s2.FirstName, s2.LastName), "student", "app.html")
	createUser(client, s3.StudentId, s3.FirstName+" "+s3.LastName, fmt.Sprintf("%s.%s@student.l8learn.local", s3.FirstName, s3.LastName), "student", "app.html")
	createUser(client, s4.StudentId, s4.FirstName+" "+s4.LastName, fmt.Sprintf("%s.%s@student.l8learn.local", s4.FirstName, s4.LastName), "student", "app.html")
	createUser(client, s5.StudentId, s5.FirstName+" "+s5.LastName, fmt.Sprintf("%s.%s@student.l8learn.local", s5.FirstName, s5.LastName), "student", "app.html")

	// Upload profiles for all 5 students
	profiles := generateSecurityProfiles()
	if err := client.Post("/learn/20/Profile", &learn.StudentProfileList{List: profiles}); err != nil {
		fmt.Printf("  ERROR profiles: %v\n", err)
	}
	fmt.Printf("  Profiles: %d\n", len(profiles))

	// Upload evaluations
	evals := generateSecurityEvals()
	if err := client.Post("/learn/20/EvalImprt", &learn.EvalImportList{List: evals}); err != nil {
		fmt.Printf("  ERROR evals: %v\n", err)
	}
	fmt.Printf("  Evaluations: %d\n", len(evals))

	// Upload schedules
	schedules := generateSecuritySchedules()
	if err := client.Post("/learn/30/Schedule", &learn.DailyScheduleList{List: schedules}); err != nil {
		fmt.Printf("  ERROR schedules: %v\n", err)
	}
	fmt.Printf("  Schedules: %d\n", len(schedules))

	fmt.Println("\n  === Security Test Accounts ===")
	fmt.Println("  Guardian: maria.garcia@family.test / 12345678 → sees Diego + Marco Garcia only")
	fmt.Println("  Guardian: james.wilson@family.test / 12345678 → sees Emma Wilson only")
	fmt.Println("  Guardian: aisha.khan@family.test / 12345678   → sees Zara + Layla Khan only")
	fmt.Println("  Student:  Diego.Garcia@student.l8learn.local / 12345678 → sees only self")
	fmt.Println("  Student:  Marco.Garcia@student.l8learn.local / 12345678 → sees only self")
	fmt.Println("  Student:  Emma.Wilson@student.l8learn.local / 12345678  → sees only self")
	fmt.Println("  Student:  Zara.Khan@student.l8learn.local / 12345678    → sees only self")
	fmt.Println("  Student:  Layla.Khan@student.l8learn.local / 12345678   → sees only self")
}
