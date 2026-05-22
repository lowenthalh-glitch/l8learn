/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package growth

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newGrowthServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.GrowthRecord{}, vnic).
		BeforeAction(func(elem interface{}, action ifs.Action, vnic ifs.IVNic) error {
			if action == ifs.POST {
				common.GenerateID(&elem.(*learn.GrowthRecord).GrowthId)
			}
			return nil
		}).
		Require(func(v interface{}) string { return v.(*learn.GrowthRecord).GrowthId }, "GrowthId").
		Require(func(v interface{}) string { return v.(*learn.GrowthRecord).StudentId }, "StudentId").
		Require(func(v interface{}) string { return v.(*learn.GrowthRecord).AcademicYear }, "AcademicYear").
		Build()
}

// NOTE: Growth records are COMPUTED, not manually created.
// Triggered by SkillMastery callback (After PATCH):
//   1. Recalculate current_score from all SkillMastery records for this student+subject
//   2. Compute absolute_growth = current_score - baseline_score
//   3. Compare to expected_growth for this starting level (from historical peer data)
//   4. Set growth_vs_expected = absolute_growth / expected_growth
//   5. Assign rating based on growth_percentile
//
// Year-end scheduled job freezes the record as a permanent snapshot.
