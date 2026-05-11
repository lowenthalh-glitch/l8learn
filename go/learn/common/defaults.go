/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package common

import (
	l8c "github.com/saichler/l8common/go/common"
	"github.com/saichler/l8types/go/ifs"
)

const PREFIX = "/learn"

func CreateResources(alias string, logVnet bool) ifs.IResources {
	return l8c.CreateResources(alias, logVnet)
}

var WaitForSignal = l8c.WaitForSignal
var OpenDBConection = l8c.OpenDBConection
