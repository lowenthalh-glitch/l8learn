/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package tutoring

import (
	"github.com/saichler/l8common/go/common"
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newTutorMatchServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.TutorMatch{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.TutorMatch).MatchId }, "MatchId").
		Require(func(v interface{}) string { return v.(*learn.TutorMatch).TutorId }, "TutorId").
		Require(func(v interface{}) string { return v.(*learn.TutorMatch).LearnerId }, "LearnerId").
		Require(func(v interface{}) string { return v.(*learn.TutorMatch).SkillId }, "SkillId").
		After(onTutorMatchChange).
		Build()
}

func onTutorMatchChange(elem interface{}, action ifs.Action, notify bool, vnic ifs.IVNic) (interface{}, bool, error) {
	match := elem.(*learn.TutorMatch)

	if action == ifs.POST {
		// New tutor match created:
		// 1. Create CollabGroup of type TUTORING for the pair
		// 2. Generate tutor guide prompts via AI
		//    ("Here's how to explain equivalent fractions to Jake — try pizza slices")
		// 3. Notify both students via l8notify
		// 4. Schedule check-in after 3 sessions
		_ = match
	}

	if action == ifs.PUT && match.Successful {
		// Tutoring completed successfully (learner reached PROFICIENT):
		// 1. Award badges to both tutor and learner
		// 2. Update both students' EngagementMetric (collaboration scores)
		// 3. Feed tutoring effectiveness into ContentEffect
		//    (peer tutoring is a "content type" with measurable outcomes)
		// 4. Credit tutor with deeper mastery (teaching = highest Bloom's taxonomy)
	}

	return nil, true, nil
}

// NOTE: Weekly AI batch job (outside callback) will:
// 1. For every student at EMERGING or DEVELOPING on a skill:
//    - Find students in same classroom at MASTERED or EXEMPLARY
//    - Filter: tutor has high engagement, not already assigned 2+ pairs
//    - Rank by social compatibility (have they interacted positively?)
//    - Create TutorMatch if good candidate found
