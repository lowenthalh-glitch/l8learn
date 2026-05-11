/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package services

import (
	"github.com/saichler/l8learn/go/learn/content/activities"
	"github.com/saichler/l8learn/go/learn/content/courses"
	"github.com/saichler/l8learn/go/learn/content/familyactivities"
	"github.com/saichler/l8learn/go/learn/content/lessons"
	"github.com/saichler/l8learn/go/learn/content/projects"
	"github.com/saichler/l8learn/go/learn/content/realworld"
	"github.com/saichler/l8learn/go/learn/content/units"
	"github.com/saichler/l8learn/go/learn/content/worksheets"
	"github.com/saichler/l8types/go/ifs"
)

func collectContentActivations(creds, dbname string, nic ifs.IVNic) []func() {
	return []func(){
		func() { courses.Activate(creds, dbname, nic) },
		func() { units.Activate(creds, dbname, nic) },
		func() { lessons.Activate(creds, dbname, nic) },
		func() { activities.Activate(creds, dbname, nic) },
		func() { worksheets.Activate(creds, dbname, nic) },
		func() { familyactivities.Activate(creds, dbname, nic) },
		func() { realworld.Activate(creds, dbname, nic) },
		func() { projects.Activate(creds, dbname, nic) },
	}
}
