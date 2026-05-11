/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package cohorts

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newCohortServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.CohortSnapshot{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.CohortSnapshot).SnapshotId }, "SnapshotId").
		Require(func(v interface{}) string { return v.(*learn.CohortSnapshot).EntityId }, "EntityId").
		Require(func(v interface{}) string { return v.(*learn.CohortSnapshot).AcademicYear }, "AcademicYear").
		Enum(func(v interface{}) int32 { return int32(v.(*learn.CohortSnapshot).Level) }, "Level", learn.AggregationLevel_name).
		Enum(func(v interface{}) int32 { return int32(v.(*learn.CohortSnapshot).Type) }, "Type", learn.SnapshotType_name).
		Build()
}

// NOTE: Cohort snapshots are COMPUTED by scheduled batch job:
// - Weekly during school year (per classroom, per school, per district)
// - Monthly during summer
// - Year-end permanent snapshots
//
// Aggregation pipeline:
//   1. For each entity (classroom/school/district):
//   2. Load all student SkillMastery records within scope
//   3. Compute mastery distribution (exemplary/mastered/proficient/developing/emerging)
//   4. Compute mean/median/std scores
//   5. Compute growth metrics (mean growth, growth vs expected)
//   6. Compute engagement (mean weekly minutes, participation rate, disengaged count)
//   7. Identify top skill gaps (lowest mastery_rate skills)
//   8. Store snapshot
