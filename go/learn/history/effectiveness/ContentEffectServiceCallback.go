/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package effectiveness

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newContentEffectServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.ContentEffect{}, vnic).
		BeforeAction(func(elem interface{}, action ifs.Action, vnic ifs.IVNic) error {
			if action == ifs.POST {
				common.GenerateID(&elem.(*learn.ContentEffect).EffectId)
			}
			return nil
		}).
		Require(func(v interface{}) string { return v.(*learn.ContentEffect).EffectId }, "EffectId").
		Require(func(v interface{}) string { return v.(*learn.ContentEffect).ContentId }, "ContentId").
		Require(func(v interface{}) string { return v.(*learn.ContentEffect).ContentType }, "ContentType").
		Require(func(v interface{}) string { return v.(*learn.ContentEffect).AcademicYear }, "AcademicYear").
		Build()
}

// NOTE: Content effectiveness records are COMPUTED by quarterly batch job:
//   1. For each activity/lesson/course:
//   2. Load all LearningSession interactions that include this content
//   3. Compute:
//      - total_attempts, unique_students
//      - mean_completion_rate (completed / started)
//      - mean_time_seconds
//      - mean_score
//      - mean_mastery_gain (average mastery improvement after using this content)
//      - mastery_gain_per_minute (efficiency metric)
//      - students_gained_mastery (moved from <PROFICIENT to >=PROFICIENT)
//   4. Compute effectiveness_percentile vs other content targeting same skills
//   5. Generate AI analysis narrative (why this works or doesn't)
//   6. Store ContentEffect record
