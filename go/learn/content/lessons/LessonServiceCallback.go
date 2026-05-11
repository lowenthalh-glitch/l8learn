/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package lessons

import (
	"github.com/saichler/l8common/go/common"
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newLessonServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.Lesson{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.Lesson).LessonId }, "LessonId").
		Require(func(v interface{}) string { return v.(*learn.Lesson).UnitId }, "UnitId").
		Require(func(v interface{}) string { return v.(*learn.Lesson).Name }, "Name").
		Enum(func(v interface{}) int32 { return int32(v.(*learn.Lesson).Difficulty) }, "Difficulty", learn.DifficultyLevel_name).
		Enum(func(v interface{}) int32 { return int32(v.(*learn.Lesson).Status) }, "Status", learn.ContentStatus_name).
		Build()
}
