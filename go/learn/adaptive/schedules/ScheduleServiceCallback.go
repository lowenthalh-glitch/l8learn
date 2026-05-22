/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package schedules

import (
	"fmt"
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
	"time"
)

func newScheduleServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.DailySchedule{}, vnic).
		BeforeAction(func(elem interface{}, action ifs.Action, vnic ifs.IVNic) error {
			if action == ifs.POST {
				sched := elem.(*learn.DailySchedule)
				common.GenerateID(&sched.ScheduleId)
				fmt.Printf("[Schedule] POST for %s, studentId=%s\n",
					sched.ScheduleId, sched.CustomFields["studentId"])
				// Launch async schedule + lesson generation
				if sHandler.llmClient != nil {
					go func() {
						time.Sleep(500 * time.Millisecond)
						sHandler.generateSchedule(sched, vnic)
					}()
				} else {
					fmt.Println("[Schedule] WARNING: llmClient is nil, skipping generation")
				}
			}
			return nil
		}).
		Require(func(v interface{}) string { return v.(*learn.DailySchedule).ScheduleId }, "ScheduleId").
		Require(func(v interface{}) string { return v.(*learn.DailySchedule).FamilyId }, "FamilyId").
		Build()
}
