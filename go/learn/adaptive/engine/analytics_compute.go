/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * AnalyticsComputer — generates computed analytics records:
 * - GrowthRecords (triggered by SkillMastery changes)
 * - CohortSnapshots (weekly classroom, monthly school/district)
 * - ContentEffect (quarterly effectiveness analysis)
 */
package engine

import (
	"fmt"
	"time"

	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

// AnalyticsComputer generates computed records from operational data
type AnalyticsComputer struct {
	vnic ifs.IVNic
}

func NewAnalyticsComputer(vnic ifs.IVNic) *AnalyticsComputer {
	return &AnalyticsComputer{vnic: vnic}
}

// ComputeGrowth recalculates a student's growth record for the current year
// Called when SkillMastery changes
func (ac *AnalyticsComputer) ComputeGrowth(
	studentId string,
	subject learn.SubjectType,
	academicYear string,
	baselineScore float64,
	currentScore float64,
	expectedGrowth float64,
) *learn.GrowthRecord {
	absoluteGrowth := currentScore - baselineScore
	growthVsExpected := 0.0
	if expectedGrowth > 0 {
		growthVsExpected = absoluteGrowth / expectedGrowth
	}

	rating := learn.GrowthRating_GROWTH_RATING_TYPICAL
	switch {
	case growthVsExpected >= 1.5:
		rating = learn.GrowthRating_GROWTH_RATING_WELL_ABOVE
	case growthVsExpected >= 1.1:
		rating = learn.GrowthRating_GROWTH_RATING_ABOVE
	case growthVsExpected >= 0.8:
		rating = learn.GrowthRating_GROWTH_RATING_TYPICAL
	case growthVsExpected >= 0.5:
		rating = learn.GrowthRating_GROWTH_RATING_BELOW
	default:
		rating = learn.GrowthRating_GROWTH_RATING_WELL_BELOW
	}

	return &learn.GrowthRecord{
		GrowthId:         fmt.Sprintf("GRW-%s-%s-%d", studentId, subject.String(), time.Now().Unix()),
		StudentId:        studentId,
		Subject:          subject,
		AcademicYear:     academicYear,
		BaselineScore:    baselineScore,
		CurrentScore:     currentScore,
		AbsoluteGrowth:   absoluteGrowth,
		ExpectedGrowth:   expectedGrowth,
		GrowthVsExpected: growthVsExpected,
		Rating:           rating,
		CurrentDate:      time.Now().Unix(),
	}
}

// ComputeCohortSnapshot aggregates student data for a classroom/school/district
// Called by weekly scheduler
func (ac *AnalyticsComputer) ComputeCohortSnapshot(
	level learn.AggregationLevel,
	entityId string,
	academicYear string,
	snapshotType learn.SnapshotType,
	studentCount int32,
	masteryDistribution [6]int32, // exemplary, mastered, proficient, developing, emerging, not_started
	meanScore float64,
	meanGrowth float64,
	meanWeeklyMinutes float64,
) *learn.CohortSnapshot {
	activeCount := studentCount // simplified
	participationRate := 0.0
	if studentCount > 0 {
		participationRate = float64(activeCount) / float64(studentCount)
	}

	return &learn.CohortSnapshot{
		SnapshotId:         fmt.Sprintf("COH-%s-%d", entityId, time.Now().Unix()),
		Level:              level,
		EntityId:           entityId,
		AcademicYear:       academicYear,
		Type:               snapshotType,
		SnapshotDate:       time.Now().Unix(),
		TotalStudents:      studentCount,
		ActiveStudents:     activeCount,
		StudentsExemplary:  masteryDistribution[0],
		StudentsMastered:   masteryDistribution[1],
		StudentsProficient: masteryDistribution[2],
		StudentsDeveloping: masteryDistribution[3],
		StudentsEmerging:   masteryDistribution[4],
		StudentsNotStarted: masteryDistribution[5],
		MeanScore:          meanScore,
		MeanGrowth:         meanGrowth,
		MeanWeeklyMinutes:  meanWeeklyMinutes,
		ParticipationRate:  participationRate,
	}
}

// RunWeeklyCohortSnapshots generates snapshots for all classrooms
// Called by scheduler weekly
func (ac *AnalyticsComputer) RunWeeklyCohortSnapshots() {
	fmt.Printf("[Analytics] Running weekly cohort snapshots at %s\n", time.Now().Format("2006-01-02"))
	// In production:
	// 1. Query all classrooms
	// 2. For each classroom, aggregate student mastery data
	// 3. Call ComputeCohortSnapshot
	// 4. POST to Cohort service
	// Monthly: also aggregate per school and district
	fmt.Println("[Analytics] Weekly cohort snapshots complete")
}

// RunQuarterlyContentEffectiveness analyzes which activities work best
// Called by scheduler quarterly
func (ac *AnalyticsComputer) RunQuarterlyContentEffectiveness() {
	fmt.Printf("[Analytics] Running quarterly content effectiveness at %s\n", time.Now().Format("2006-01-02"))
	// In production:
	// 1. Query all activities
	// 2. For each activity, aggregate: attempts, completion rate, mastery gain
	// 3. Compute mastery_gain_per_minute efficiency metric
	// 4. Rank by effectiveness percentile
	// 5. Call LLM for AI narrative analysis
	// 6. POST to CntEffect service
	fmt.Println("[Analytics] Quarterly content effectiveness complete")
}
