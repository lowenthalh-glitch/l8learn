/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package llmconfig

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "LLMConfig"
	ServiceArea = byte(30)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	sla := common.NewOrmSLA(ServiceName, ServiceArea, "ConfigId", newLLMConfigServiceCallback(vnic),
		&learn.LLMConfig{}, &learn.LLMConfigList{})
	common.ActivateService(sla, creds, dbname, vnic)
}
