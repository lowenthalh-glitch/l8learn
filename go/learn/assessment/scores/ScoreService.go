/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package scores

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "Score"
	ServiceArea = byte(40)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	sla := common.NewOrmSLA(ServiceName, ServiceArea, "ScoreId", newScoreServiceCallback(vnic),
		&learn.Score{}, &learn.ScoreList{})
	common.ActivateService(sla, creds, dbname, vnic)
}
