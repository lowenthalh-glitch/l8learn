/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package ui

import (
	"github.com/saichler/l8types/go/ifs"
)

func RegisterTypes(resources ifs.IResources) {
	registerContentTypes(resources)
	registerStudentTypes(resources)
	registerAdaptiveTypes(resources)
	registerAssessmentTypes(resources)
	registerAnalyticsTypes(resources)
	registerHistoryTypes(resources)
	registerCollabTypes(resources)
}
