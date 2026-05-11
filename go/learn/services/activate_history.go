/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package services

import (
	"github.com/saichler/l8learn/go/learn/history/cohorts"
	"github.com/saichler/l8learn/go/learn/history/effectiveness"
	"github.com/saichler/l8learn/go/learn/history/growth"
	"github.com/saichler/l8learn/go/learn/history/risk"
	"github.com/saichler/l8learn/go/learn/history/standards"
	"github.com/saichler/l8types/go/ifs"
)

func collectHistoryActivations(creds, dbname string, nic ifs.IVNic) []func() {
	return []func(){
		func() { growth.Activate(creds, dbname, nic) },
		func() { cohorts.Activate(creds, dbname, nic) },
		func() { risk.Activate(creds, dbname, nic) },
		func() { standards.Activate(creds, dbname, nic) },
		func() { effectiveness.Activate(creds, dbname, nic) },
	}
}
