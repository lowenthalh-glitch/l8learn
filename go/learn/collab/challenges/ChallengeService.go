/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package challenges

import (
	"github.com/saichler/l8common/go/common"
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "Challenge"
	ServiceArea = byte(70)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	common.ActivateService(common.ServiceConfig{
		ServiceName: ServiceName,
		ServiceArea: ServiceArea,
		PrimaryKey:  "ChallengeId",
		Callback:    newChallengeServiceCallback(vnic),
	}, &learn.Challenge{}, &learn.ChallengeList{}, creds, dbname, vnic)
}
