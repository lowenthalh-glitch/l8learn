/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * Verifies student profile has minimum required data for schedule generation.
 */
package schedules

import (
	"github.com/saichler/l8learn/go/types/learn"
)

// CheckProfileReady verifies the profile has minimum data for schedule generation.
func CheckProfileReady(profile *learn.StudentProfile) (bool, string) {
	if profile == nil {
		return false, "No student profile found. Upload and process evaluations first."
	}
	if profile.Scores == nil || profile.Scores.OverallAcademicReadiness == 0 {
		return false, "Profile has no readiness scores. Upload and process evaluations first."
	}
	if profile.LearningStyle == nil || profile.LearningStyle.BestSessionLengthMinutes == 0 {
		return false, "Profile has no learning style data. Upload and process evaluations first."
	}
	if profile.Attention == nil || profile.Attention.FocusAcademicTaskMinutes == 0 {
		return false, "Profile has no attention data. Upload and process evaluations first."
	}
	return true, ""
}
