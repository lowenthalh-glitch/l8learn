/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package services

import (
	"github.com/lowenthalh-glitch/l8learn/go/learn/collab/challenges"
	"github.com/lowenthalh-glitch/l8learn/go/learn/collab/groups"
	"github.com/lowenthalh-glitch/l8learn/go/learn/collab/messages"
	"github.com/lowenthalh-glitch/l8learn/go/learn/collab/tutoring"
	"github.com/saichler/l8types/go/ifs"
)

func collectCollabActivations(creds, dbname string, nic ifs.IVNic) []func() {
	return []func(){
		func() { groups.Activate(creds, dbname, nic) },
		func() { messages.Activate(creds, dbname, nic) },
		func() { tutoring.Activate(creds, dbname, nic) },
		func() { challenges.Activate(creds, dbname, nic) },
	}
}
