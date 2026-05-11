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

func registerAssessmentTypes(resources ifs.IResources) {
	
	l8c.RegisterType(resources, &learn.LearningSession{}, &learn.LearningSessionList{}, "SessionId")
	l8c.RegisterType(resources, &learn.Score{}, &learn.ScoreList{}, "ScoreId")
	l8c.RegisterType(resources, &learn.Benchmark{}, &learn.BenchmarkList{}, "BenchmarkId")
	l8c.RegisterType(resources, &learn.WorksheetScan{}, &learn.WorksheetScanList{}, "ScanId")
}
