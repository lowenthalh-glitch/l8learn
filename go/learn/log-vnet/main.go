/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package main

import (
	"github.com/saichler/l8bus/go/overlay/vnet"
	"github.com/saichler/l8learn/go/learn/common"
	"github.com/saichler/l8logfusion/go/agent/logserver"
)

func main() {
	logsDbDirectory := "/data/logsdb/l8learn"
	resources := common.CreateResources("log-vnet")
	resources.SysConfig().VnetPort = resources.SysConfig().LogConfig.VnetPort
	net := vnet.NewVNet(resources, true)
	net.Start()
	logserver.ActivateLogService(logsDbDirectory, net.VnetVnic())
	resources.Logger().Info("logs vnet started!")
	common.WaitForSignal(resources)
}
