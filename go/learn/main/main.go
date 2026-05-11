/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/saichler/l8bus/go/overlay/vnic"
	"github.com/saichler/l8learn/go/learn/common"
	learnservices "github.com/saichler/l8learn/go/learn/services"
	"github.com/saichler/l8types/go/ifs"
	"github.com/saichler/l8utils/go/utils/ipsegment"
)

const (
	DB_CREDS = "PostgreSQL"
	DB_NAME  = "l8learn"
)

var dbName = DB_NAME

func main() {
	res := common.CreateResources("L8LearnServices")
	ifs.SetNetworkMode(ifs.NETWORK_K8s)
	nic := vnic.NewVirtualNetworkInterface(res, nil)
	nic.Start()
	nic.WaitForConnection()

	if len(os.Args) == 1 {
		startDb(nic)
	} else {
		ipsegment.MachineIP = "127.0.0.1"
		_, user, pass, _, err := nic.Resources().Security().Credential(DB_CREDS, DB_NAME, nic.Resources())
		if err == nil && user == "admin" && pass == "admin" {
			dbName = "admin"
		}
	}

	// Phase 1: All CRUD services (parallel)
	learnservices.ActivateAllServices(DB_CREDS, dbName, nic)

	// Phase 2: AI chat (needs full introspector populated)
	learnservices.ActivateChatService(DB_CREDS, dbName, nic)

	fmt.Println("[l8learn] All services activated!")
	common.WaitForSignal(res)
}

func startDb(nic ifs.IVNic) {
	_, user, pass, _, err := nic.Resources().Security().Credential(DB_CREDS, DB_NAME, nic.Resources())
	if err != nil {
		panic(DB_CREDS + " " + err.Error())
	}
	if user == "admin" && pass == "admin" {
		dbName = "admin"
	}

	cmd := exec.Command("nohup", "/start-postgres.sh", dbName, user, pass)
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
	time.Sleep(time.Second * 5)
}
