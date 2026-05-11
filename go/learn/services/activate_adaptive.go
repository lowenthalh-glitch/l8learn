/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package services

import (
	"github.com/saichler/l8learn/go/learn/adaptive/llmconfig"
	"github.com/saichler/l8learn/go/learn/adaptive/mastery"
	"github.com/saichler/l8learn/go/learn/adaptive/paths"
	"github.com/saichler/l8learn/go/learn/adaptive/promptlogs"
	"github.com/saichler/l8learn/go/learn/adaptive/rules"
	"github.com/saichler/l8learn/go/learn/adaptive/schedules"
	"github.com/saichler/l8learn/go/learn/adaptive/skills"
	"github.com/saichler/l8types/go/ifs"
)

func collectAdaptiveActivations(creds, dbname string, nic ifs.IVNic) []func() {
	return []func(){
		func() { skills.Activate(creds, dbname, nic) },
		func() { mastery.Activate(creds, dbname, nic) },
		func() { paths.Activate(creds, dbname, nic) },
		func() { rules.Activate(creds, dbname, nic) },
		func() { schedules.Activate(creds, dbname, nic) },
		func() { promptlogs.Activate(creds, dbname, nic) },
		func() { llmconfig.Activate(creds, dbname, nic) },
	}
}
