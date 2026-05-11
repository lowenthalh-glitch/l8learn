/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package services

import (
	"github.com/saichler/l8learn/go/learn/students/classrooms"
	"github.com/saichler/l8learn/go/learn/students/compliance"
	"github.com/saichler/l8learn/go/learn/students/districts"
	"github.com/saichler/l8learn/go/learn/students/enrollments"
	"github.com/saichler/l8learn/go/learn/students/evalimports"
	"github.com/saichler/l8learn/go/learn/students/families"
	"github.com/saichler/l8learn/go/learn/students/guardians"
	"github.com/saichler/l8learn/go/learn/students/pods"
	"github.com/saichler/l8learn/go/learn/students/profiles"
	"github.com/saichler/l8learn/go/learn/students/schools"
	"github.com/saichler/l8learn/go/learn/students/students"
	"github.com/saichler/l8learn/go/learn/students/teachers"
	"github.com/saichler/l8types/go/ifs"
)

func collectStudentActivations(creds, dbname string, nic ifs.IVNic) []func() {
	return []func(){
		func() { students.Activate(creds, dbname, nic) },
		func() { guardians.Activate(creds, dbname, nic) },
		func() { teachers.Activate(creds, dbname, nic) },
		func() { classrooms.Activate(creds, dbname, nic) },
		func() { schools.Activate(creds, dbname, nic) },
		func() { districts.Activate(creds, dbname, nic) },
		func() { enrollments.Activate(creds, dbname, nic) },
		func() { families.Activate(creds, dbname, nic) },
		func() { compliance.Activate(creds, dbname, nic) },
		func() { pods.Activate(creds, dbname, nic) },
		func() { profiles.Activate(creds, dbname, nic) },
		func() { evalimports.Activate(creds, dbname, nic) },
	}
}
