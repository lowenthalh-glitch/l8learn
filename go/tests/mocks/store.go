/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package mocks

// MockDataStore holds generated IDs for cross-referencing between phases
type MockDataStore struct {
	// Phase 1: Foundation
	DistrictIDs []string
	SchoolIDs   []string
	SkillIDs    []string
	RuleIDs     []string

	// Phase 2: People
	TeacherIDs   []string
	ClassroomIDs []string
	GuardianIDs  []string
	StudentIDs   []string
	EnrollmentIDs []string
	FamilyIDs    []string

	// Phase 3: Content
	CourseIDs   []string
	UnitIDs     []string
	LessonIDs   []string
	ActivityIDs []string
	BenchmarkIDs []string

	// Phase 4: Learning State
	PathIDs    []string
	MasteryIDs []string

	// Phase 5: Sessions
	SessionIDs []string
	ScoreIDs   []string

	// Phase 6: Worksheets
	WorksheetIDs []string
	ScanIDs      []string

	// Phase 7: Analytics
	ProgressIDs   []string
	EngagementIDs []string

	// Phase 8: Homeschool
	ComplianceIDs     []string
	PodIDs            []string
	FamilyActivityIDs []string
	RealWorldIDs      []string
	ProjectIDs        []string
	ScheduleIDs       []string

	// Phase 9: Collaboration
	GroupIDs     []string
	MessageIDs   []string
	TutorIDs     []string
	ChallengeIDs []string

	// Phase 10: History
	GrowthIDs        []string
	CohortIDs        []string
	RiskIDs          []string
	StandardIDs      []string
	EffectivenessIDs []string
}

func NewMockDataStore() *MockDataStore {
	return &MockDataStore{}
}

// pickRef safely picks from an ID slice using modulo
func pickRef(ids []string, index int) string {
	if len(ids) == 0 {
		return ""
	}
	return ids[index%len(ids)]
}
