/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package standards

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newStandardMasteryServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.StandardMastery{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.StandardMastery).StandardMasteryId }, "StandardMasteryId").
		Require(func(v interface{}) string { return v.(*learn.StandardMastery).StudentId }, "StudentId").
		Require(func(v interface{}) string { return v.(*learn.StandardMastery).StandardId }, "StandardId").
		Require(func(v interface{}) string { return v.(*learn.StandardMastery).AcademicYear }, "AcademicYear").
		Enum(func(v interface{}) int32 { return int32(v.(*learn.StandardMastery).Subject) }, "Subject", learn.SubjectType_name).
		Enum(func(v interface{}) int32 { return int32(v.(*learn.StandardMastery).Level) }, "Level", learn.MasteryLevel_name).
		Build()
}

// NOTE: Standard mastery records are COMPUTED, triggered by SkillMastery callback:
//   1. When a skill's mastery changes, look up Skill.StandardIds
//   2. For each standard that includes this skill:
//      - Load all SkillMastery records for skills mapped to this standard
//      - Compute skills_mastered / skills_in_standard
//      - Set level based on ratio (same thresholds as skill mastery)
//      - Update score (0-100)
//   3. PATCH the StandardMastery record (or POST if first time)
