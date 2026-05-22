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
	"github.com/saichler/l8learn/go/learn/adaptive/engine"
	"github.com/saichler/l8learn/go/learn/common"
	learnservices "github.com/saichler/l8learn/go/learn/services"
	"github.com/saichler/l8learn/go/learn/adaptive/schedules"
	"github.com/saichler/l8learn/go/learn/students/evalimports"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
	"github.com/saichler/l8utils/go/utils/ipsegment"
)

const (
	DB_CREDS = "postgres"
	DB_NAME  = "admin"
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

	// Phase 3: Initialize LLM client for eval import pipeline
	initLLMClient(nic)

	fmt.Println("[l8learn] All services activated!")
	common.WaitForSignal(res)
}

func initLLMClient(nic ifs.IVNic) {
	_, _, apiKey, _, err := nic.Resources().Security().Credential("Anthropic", "API_KEY", nic.Resources())
	masker := engine.NewPIIMasker()
	logger := engine.NewPromptLogger(nic)
	var llmClient engine.LLMClient
	if err == nil && apiKey != "" {
		keyPreview := apiKey
		if len(keyPreview) > 12 {
			keyPreview = apiKey[:8] + "..." + apiKey[len(apiKey)-4:]
		}
		fmt.Printf("[l8learn] API key loaded: %s (len=%d)\n", keyPreview, len(apiKey))
		llmClient = engine.NewLLMClient(learn.LLMMode_LLM_MODE_LIVE, apiKey, masker, logger)
		fmt.Println("[l8learn] LLM client initialized in LIVE mode")
	} else {
		llmClient = engine.NewLLMClient(learn.LLMMode_LLM_MODE_SIMULATE, "", masker, logger)
		fmt.Println("[l8learn] No Anthropic API key found, using LLM simulator")
	}
	evalimports.SetLLMClient(llmClient, masker)
	schedules.SetLLMClient(llmClient, masker)
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
