/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package ui

import (
	l8c "github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func registerHistoryTypes(resources ifs.IResources) {
	
	l8c.RegisterType(resources, &learn.GrowthRecord{}, &learn.GrowthRecordList{}, "GrowthId")
	l8c.RegisterType(resources, &learn.CohortSnapshot{}, &learn.CohortSnapshotList{}, "SnapshotId")
	l8c.RegisterType(resources, &learn.RiskAssessment{}, &learn.RiskAssessmentList{}, "AssessmentId")
	l8c.RegisterType(resources, &learn.StandardMastery{}, &learn.StandardMasteryList{}, "StandardMasteryId")
	l8c.RegisterType(resources, &learn.ContentEffect{}, &learn.ContentEffectList{}, "EffectId")
}
