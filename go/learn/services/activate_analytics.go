/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package services

import (
	"github.com/lowenthalh-glitch/l8learn/go/learn/analytics/engagement"
	"github.com/lowenthalh-glitch/l8learn/go/learn/analytics/progress"
	"github.com/saichler/l8types/go/ifs"
)

func collectAnalyticsActivations(creds, dbname string, nic ifs.IVNic) []func() {
	return []func(){
		func() { progress.Activate(creds, dbname, nic) },
		func() { engagement.Activate(creds, dbname, nic) },
	}
}
