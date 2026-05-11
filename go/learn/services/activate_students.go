/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package services

import (
	"github.com/lowenthalh-glitch/l8learn/go/learn/students/classrooms"
	"github.com/lowenthalh-glitch/l8learn/go/learn/students/compliance"
	"github.com/lowenthalh-glitch/l8learn/go/learn/students/districts"
	"github.com/lowenthalh-glitch/l8learn/go/learn/students/enrollments"
	"github.com/lowenthalh-glitch/l8learn/go/learn/students/families"
	"github.com/lowenthalh-glitch/l8learn/go/learn/students/guardians"
	"github.com/lowenthalh-glitch/l8learn/go/learn/students/pods"
	"github.com/lowenthalh-glitch/l8learn/go/learn/students/schools"
	"github.com/lowenthalh-glitch/l8learn/go/learn/students/students"
	"github.com/lowenthalh-glitch/l8learn/go/learn/students/teachers"
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
	}
}
