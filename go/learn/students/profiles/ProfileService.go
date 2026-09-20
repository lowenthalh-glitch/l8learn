/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package profiles

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "Profile"
	ServiceArea = byte(20)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	sla := common.NewOrmSLA(ServiceName, ServiceArea, "ProfileId", newProfileServiceCallback(vnic),
		&learn.StudentProfile{}, &learn.StudentProfileList{})
	common.ActivateService(sla, creds, dbname, vnic)
}
