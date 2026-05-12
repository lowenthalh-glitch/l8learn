/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * Diagnostic Benchmark Engine — adaptive placement algorithm.
 * Determines a new student's starting level across skills.
 */
package engine

import (
	"fmt"
	"time"

	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const (
	questionsPerSkill  = 4
	advanceThreshold   = 0.80 // >80% correct → try harder skill
	dropThreshold      = 0.40 // <40% correct → stop going harder
	maxSkillsPerSubject = 10  // max skills to test per subject
)

// DiagnosticEngine runs the adaptive placement assessment
type DiagnosticEngine struct {
	vnic      ifs.IVNic
	llmClient LLMClient
}

func NewDiagnosticEngine(vnic ifs.IVNic, llmClient LLMClient) *DiagnosticEngine {
	return &DiagnosticEngine{vnic: vnic, llmClient: llmClient}
}

// DiagnosticResult holds the outcome of a diagnostic session
type DiagnosticResult struct {
	SkillResults   []*SkillDiagResult
	ReadinessScores *learn.WorkingScores
	MasteryRecords []*learn.SkillMastery
	ProfileUpdates *learn.StudentProfile
}

// SkillDiagResult holds results for one skill in the diagnostic
type SkillDiagResult struct {
	SkillId     string
	SkillName   string
	Subject     learn.SubjectType
	GradeLevel  learn.GradeLevel
	Correct     int
	Total       int
	Accuracy    float64
	Level       learn.MasteryLevel
	IsCeiling   bool // Student couldn't go higher
	IsFloor     bool // Student couldn't go lower (already at easiest)
}

// RunDiagnostic executes the adaptive placement for a student.
// Returns results that should be used to create SkillMastery records
// and populate the StudentProfile.
func (d *DiagnosticEngine) RunDiagnostic(
	student *learn.Student,
	skills []*learn.Skill,
	activities []*learn.Activity,
) *DiagnosticResult {
	result := &DiagnosticResult{}

	// Group skills by subject, ordered by grade level
	mathSkills := filterSkillsBySubject(skills, learn.SubjectType_SUBJECT_TYPE_MATH)
	readingSkills := filterSkillsBySubject(skills, learn.SubjectType_SUBJECT_TYPE_READING)

	// Start at the student's enrolled grade level
	startGrade := student.GradeLevel

	// Run diagnostic per subject
	mathResults := d.diagnoseSubject(mathSkills, startGrade, activities)
	readingResults := d.diagnoseSubject(readingSkills, startGrade, activities)

	result.SkillResults = append(result.SkillResults, mathResults...)
	result.SkillResults = append(result.SkillResults, readingResults...)

	// Compute readiness scores from results
	result.ReadinessScores = computeReadiness(mathResults, readingResults)

	// Create SkillMastery records from results
	now := time.Now().Unix()
	for _, sr := range result.SkillResults {
		mastery := &learn.SkillMastery{
			MasteryId:       fmt.Sprintf("MST-DIAG-%s", sr.SkillId),
			StudentId:       student.StudentId,
			SkillId:         sr.SkillId,
			Level:           sr.Level,
			Confidence:      0.7, // Diagnostic confidence (limited questions)
			AttemptsCount:   int32(sr.Total),
			CorrectCount:    int32(sr.Correct),
			CurrentAccuracy: sr.Accuracy,
			FirstAttempted:  now,
			LastAttempted:   now,
		}
		if sr.Level >= learn.MasteryLevel_MASTERY_LEVEL_MASTERED {
			mastery.MasteredDate = now
		}
		result.MasteryRecords = append(result.MasteryRecords, mastery)
	}

	return result
}

// diagnoseSubject runs the adaptive algorithm for one subject
func (d *DiagnosticEngine) diagnoseSubject(
	skills []*learn.Skill,
	startGrade learn.GradeLevel,
	activities []*learn.Activity,
) []*SkillDiagResult {
	var results []*SkillDiagResult

	// Find skills at the starting grade level
	gradeSkills := filterSkillsByGrade(skills, startGrade)
	if len(gradeSkills) == 0 {
		return results
	}

	// Test skills at this grade
	for _, skill := range gradeSkills {
		if len(results) >= maxSkillsPerSubject {
			break
		}
		sr := d.testSkill(skill, activities)
		results = append(results, sr)
	}

	// If doing well (>advanceThreshold avg), try next grade up
	avgAccuracy := averageAccuracy(results)
	if avgAccuracy > advanceThreshold {
		nextGrade := startGrade + 1
		upperSkills := filterSkillsByGrade(skills, nextGrade)
		for _, skill := range upperSkills {
			if len(results) >= maxSkillsPerSubject {
				break
			}
			sr := d.testSkill(skill, activities)
			sr.IsCeiling = sr.Accuracy < dropThreshold
			results = append(results, sr)
			if sr.IsCeiling {
				break // Found ceiling
			}
		}
	}

	// If doing poorly (<dropThreshold avg), try grade below
	if avgAccuracy < dropThreshold && startGrade > learn.GradeLevel_GRADE_LEVEL_K {
		prevGrade := startGrade - 1
		lowerSkills := filterSkillsByGrade(skills, prevGrade)
		for _, skill := range lowerSkills {
			if len(results) >= maxSkillsPerSubject {
				break
			}
			sr := d.testSkill(skill, activities)
			sr.IsFloor = sr.Accuracy > advanceThreshold
			results = append(results, sr)
			if sr.IsFloor {
				break // Found floor
			}
		}
	}

	return results
}

// testSkill simulates testing a student on one skill
// In production, this would present questions via the student player
// and collect real answers. For now, it generates a simulated result.
func (d *DiagnosticEngine) testSkill(skill *learn.Skill, activities []*learn.Activity) *SkillDiagResult {
	// Find activities for this skill
	// In a real implementation, this would:
	// 1. Select questionsPerSkill questions from matching activities
	// 2. Present them to the student via the student player
	// 3. Collect answers and score them
	// For now, return a simulated result

	sr := &SkillDiagResult{
		SkillId:    skill.SkillId,
		SkillName:  skill.Name,
		Subject:    skill.Subject,
		GradeLevel: skill.GradeLevel,
		Total:      questionsPerSkill,
	}

	// Simulated scoring (in production, this comes from actual student answers)
	// The simulation uses the skill's grade level relative to the student
	sr.Correct = 2 // Default to 50% for simulation
	sr.Accuracy = float64(sr.Correct) / float64(sr.Total)
	sr.Level = calculateMasteryLevel(sr.Accuracy)

	return sr
}

func computeReadiness(mathResults, readingResults []*SkillDiagResult) *learn.WorkingScores {
	return &learn.WorkingScores{
		OverallAcademicReadiness: int32(averageAccuracy(append(mathResults, readingResults...)) * 100),
		ReadingReadiness:  int32(averageAccuracy(readingResults) * 100),
		MathReadiness:     int32(averageAccuracy(mathResults) * 100),
	}
}

func filterSkillsBySubject(skills []*learn.Skill, subject learn.SubjectType) []*learn.Skill {
	var filtered []*learn.Skill
	for _, s := range skills {
		if s.Subject == subject {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

func filterSkillsByGrade(skills []*learn.Skill, grade learn.GradeLevel) []*learn.Skill {
	var filtered []*learn.Skill
	for _, s := range skills {
		if s.GradeLevel == grade {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

func averageAccuracy(results []*SkillDiagResult) float64 {
	if len(results) == 0 {
		return 0
	}
	total := 0.0
	for _, r := range results {
		total += r.Accuracy
	}
	return total / float64(len(results))
}
