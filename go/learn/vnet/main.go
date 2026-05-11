/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package main

import (
	"github.com/saichler/l8bus/go/overlay/vnet"
	"github.com/lowenthalh-glitch/l8learn/go/learn/common"
	"github.com/saichler/l8types/go/ifs"
)

func main() {
	res := common.CreateResources("L8LearnVNet", true)
	ifs.SetNetworkMode(ifs.NETWORK_K8s)
	v := vnet.NewVNet(res)
	v.Start()
	common.WaitForSignal(res)
}
