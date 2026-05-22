/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package enrollments

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newEnrollmentServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.Enrollment{}, vnic).
		BeforeAction(func(elem interface{}, action ifs.Action, vnic ifs.IVNic) error {
			if action == ifs.POST {
				common.GenerateID(&elem.(*learn.Enrollment).EnrollmentId)
			}
			return nil
		}).
		Require(func(v interface{}) string { return v.(*learn.Enrollment).EnrollmentId }, "EnrollmentId").
		Require(func(v interface{}) string { return v.(*learn.Enrollment).StudentId }, "StudentId").
		Require(func(v interface{}) string { return v.(*learn.Enrollment).SchoolId }, "SchoolId").
		Require(func(v interface{}) string { return v.(*learn.Enrollment).DistrictId }, "DistrictId").
		After(onEnrollmentChange).
		Build()
}

func onEnrollmentChange(elem interface{}, action ifs.Action, vnic ifs.IVNic) error {
	enroll := elem.(*learn.Enrollment)

	if action == ifs.POST {
		// 1. Create Student record if not existing
		// 2. Create Guardian record if not existing
		// 3. Link guardian to student
		// 4. Generate unique invite token
		// 5. Send guardian invite email via l8notify
		// 6. Update status: DRAFT → INVITED
		_ = enroll
	}

	if action == ifs.PUT {
		switch enroll.Status {
		case learn.EnrollmentStatus_ENROLLMENT_STATUS_CONSENTED:
			// Validate all required ConsentRecords are signed
			if !validateConsent(enroll) {
				// Return error — cannot transition to CONSENTED without all required consents
			}

		case learn.EnrollmentStatus_ENROLLMENT_STATUS_ACTIVE:
			// Student activated — trigger onboarding:
			// 1. Create student login via ISecurityProvider
			// 2. Assign "student" role
			// 3. Create empty StudentProfile (to be populated by diagnostic)
			// 4. Mark enrollment.DiagnosticComplete = false
			// 5. On first student login, student player checks DiagnosticComplete
			//    and triggers the diagnostic flow if false
			// 6. After diagnostic completes:
			//    a. SkillMastery records created from diagnostic results
			//    b. StudentProfile.readiness populated
			//    c. LearningPaths created (one per subject)
			//    d. enrollment.DiagnosticComplete = true
			//    e. PATH_DECISION prompt logged to PromptLog
			// 5. Notify teacher: "Student has been activated"
			// 6. Notify guardian: "Account is ready"
		}
	}

	return nil
}

func validateConsent(enroll *learn.Enrollment) bool {
	if enroll.Consents == nil || len(enroll.Consents) == 0 {
		return false
	}

	requiredTypes := []learn.ConsentType{
		learn.ConsentType_CONSENT_TYPE_COPPA,
		learn.ConsentType_CONSENT_TYPE_DATA_COLLECTION,
		learn.ConsentType_CONSENT_TYPE_AI_PERSONALIZATION,
		learn.ConsentType_CONSENT_TYPE_PROGRESS_SHARING,
	}

	signed := make(map[learn.ConsentType]bool)
	for _, c := range enroll.Consents {
		if c.Granted {
			signed[c.Type] = true
		}
	}

	for _, required := range requiredTypes {
		if !signed[required] {
			return false
		}
	}
	return true
}
