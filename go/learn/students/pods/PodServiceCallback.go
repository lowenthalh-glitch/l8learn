/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package pods

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newPodServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.LearningPod{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.LearningPod).PodId }, "PodId").
		Require(func(v interface{}) string { return v.(*learn.LearningPod).Name }, "Name").
		Require(func(v interface{}) string { return v.(*learn.LearningPod).OrganizerGuardianId }, "OrganizerGuardianId").
		Build()
}
