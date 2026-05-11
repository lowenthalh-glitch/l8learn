/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package enrollments

import (
	"github.com/saichler/l8common/go/common"
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newEnrollmentServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.Enrollment{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.Enrollment).EnrollmentId }, "EnrollmentId").
		Require(func(v interface{}) string { return v.(*learn.Enrollment).StudentId }, "StudentId").
		Require(func(v interface{}) string { return v.(*learn.Enrollment).SchoolId }, "SchoolId").
		Require(func(v interface{}) string { return v.(*learn.Enrollment).DistrictId }, "DistrictId").
		Enum(func(v interface{}) int32 { return int32(v.(*learn.Enrollment).Status) }, "Status", learn.EnrollmentStatus_name).
		After(onEnrollmentChange).
		Build()
}

func onEnrollmentChange(elem interface{}, action ifs.Action, notify bool, vnic ifs.IVNic) (interface{}, bool, error) {
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
			// 1. Create student login via ISecurityProvider
			// 2. Assign "student" role
			// 3. Create initial LearningPaths (one per configured subject)
			// 4. Schedule diagnostic benchmark
			// 5. Notify teacher: "Student has been activated"
			// 6. Notify guardian: "Account is ready"
		}
	}

	return nil, true, nil
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
