/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package services

import (
	"github.com/lowenthalh-glitch/l8learn/go/learn/assessment/benchmarks"
	"github.com/lowenthalh-glitch/l8learn/go/learn/assessment/scores"
	"github.com/lowenthalh-glitch/l8learn/go/learn/assessment/sessions"
	"github.com/lowenthalh-glitch/l8learn/go/learn/assessment/worksheetscans"
	"github.com/saichler/l8types/go/ifs"
)

func collectAssessmentActivations(creds, dbname string, nic ifs.IVNic) []func() {
	return []func(){
		func() { sessions.Activate(creds, dbname, nic) },
		func() { scores.Activate(creds, dbname, nic) },
		func() { benchmarks.Activate(creds, dbname, nic) },
		func() { worksheetscans.Activate(creds, dbname, nic) },
	}
}
