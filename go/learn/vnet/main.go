/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package main

import (
	"os"

	"github.com/saichler/l8bus/go/overlay/vnet"
	"github.com/saichler/l8learn/go/learn/common"
)

func main() {
	resources := common.CreateResources("vnet-" + os.Getenv("HOSTNAME"))
	net := vnet.NewVNet(resources)
	net.Start()
	resources.Logger().Info("vnet started!")
	common.WaitForSignal(resources)
}
