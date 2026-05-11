/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package main

import (
	"github.com/saichler/l8bus/go/overlay/vnic"
	"github.com/saichler/l8events/go/services"
	"github.com/saichler/l8learn/go/learn/common"
	learnservices "github.com/saichler/l8learn/go/learn/services"
	"github.com/saichler/l8types/go/ifs"
	"os"
)

func main() {
	res := common.CreateResources("L8LearnServices", false)
	ifs.SetNetworkMode(ifs.NETWORK_K8s)
	nic := vnic.NewVirtualNetworkInterface(res, nil)
	nic.Start()
	nic.WaitForConnection()

	if len(os.Args) == 1 {
		common.OpenDBConection("l8learn", nic)
	}

	dbcred := nic.Resources().SysConfig().DataStoreConfig.Type
	dbname := nic.Resources().SysConfig().DataStoreConfig.Name

	// Phase 1: All CRUD services (parallel)
	learnservices.ActivateAllServices(dbcred, dbname, nic)

	// Phase 2: AI chat (needs full introspector populated)
	learnservices.ActivateChatService(dbcred, dbname, nic)

	// Phase 3: Event tracking + notifications
	services.ActivateEvents(dbcred, dbname, nic)

	common.WaitForSignal(res)
}
