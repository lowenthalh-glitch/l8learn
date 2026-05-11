/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package main

import (
	"github.com/saichler/l8common/go/common"
	"github.com/lowenthalh-glitch/l8learn/go/learn/ui"
)

func main() {
	svr := common.CreateWebServer("web", ui.RegisterTypes)
	svr.Start()
}
