/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package ui

import (
	l8c "github.com/saichler/l8common/go/common"
	"github.com/saichler/l8types/go/ifs"
	l8events "github.com/saichler/l8types/go/types/l8events"
)

func RegisterTypes(resources ifs.IResources) {
	registerContentTypes(resources)
	registerStudentTypes(resources)
	registerAdaptiveTypes(resources)
	registerAssessmentTypes(resources)
	registerAnalyticsTypes(resources)
	registerHistoryTypes(resources)
	registerCollabTypes(resources)
	l8c.RegisterType(resources, &l8events.EventRecord{}, &l8events.EventRecordList{}, "EventId")
}
