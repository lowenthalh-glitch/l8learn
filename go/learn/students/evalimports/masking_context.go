/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * Builds the PII masking context by loading student-related entities
 * to collect all names that should be masked before sending to the LLM.
 */
package evalimports

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

// BuildMaskingContext loads related entities for a student and returns
// all known names that should be masked in evaluation text.
func BuildMaskingContext(studentId string, vnic ifs.IVNic) []string {
	var names []string

	// Load student
	result, err := common.GetEntity("Student", byte(20), &learn.Student{StudentId: studentId}, vnic)
	if err == nil && result != nil {
		s := result.(*learn.Student)
		addName(&names, s.FirstName)
		addName(&names, s.LastName)
		addName(&names, s.PreferredName)

		// Load school
		if s.SchoolId != "" {
			schoolResult, schoolErr := common.GetEntity("School", byte(20), &learn.School{SchoolId: s.SchoolId}, vnic)
			if schoolErr == nil && schoolResult != nil {
				addName(&names, schoolResult.(*learn.School).Name)
			}
		}
	}

	// Load guardians — use filter by studentId
	guardianResults, gErr := common.GetEntities("Guardian", byte(20), &learn.Guardian{StudentIds: []string{studentId}}, vnic)
	if gErr == nil {
		for _, g := range guardianResults {
			guardian := g.(*learn.Guardian)
			addName(&names, guardian.FirstName)
			addName(&names, guardian.LastName)
		}
	}

	return names
}

func addName(names *[]string, name string) {
	if name != "" && len(name) > 1 {
		*names = append(*names, name)
	}
}
