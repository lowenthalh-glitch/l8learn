/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package services

import (
	"github.com/saichler/l8types/go/ifs"
	"sync"
)

const parallelWorkers = 20

func ActivateAllServices(creds, dbname string, nic ifs.IVNic) {
	var all []func()
	all = append(all, collectContentActivations(creds, dbname, nic)...)
	all = append(all, collectStudentActivations(creds, dbname, nic)...)
	all = append(all, collectAdaptiveActivations(creds, dbname, nic)...)
	all = append(all, collectAssessmentActivations(creds, dbname, nic)...)
	all = append(all, collectAnalyticsActivations(creds, dbname, nic)...)
	all = append(all, collectHistoryActivations(creds, dbname, nic)...)
	all = append(all, collectCollabActivations(creds, dbname, nic)...)

	sem := make(chan struct{}, parallelWorkers)
	var wg sync.WaitGroup
	for _, fn := range all {
		wg.Add(1)
		sem <- struct{}{}
		go func(f func()) {
			defer wg.Done()
			defer func() { <-sem }()
			f()
		}(fn)
	}
	wg.Wait()
}

func ActivateChatService(creds, dbname string, nic ifs.IVNic) {
	// AI agent chat service — activated after all CRUD services
	// so the introspector has all types registered
	// TODO: agent.InitializeChat(creds, dbname, nic)
}
