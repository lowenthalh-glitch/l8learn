/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package sessions

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newSessionServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.LearningSession{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.LearningSession).SessionId }, "SessionId").
		Require(func(v interface{}) string { return v.(*learn.LearningSession).StudentId }, "StudentId").
		After(onSessionChange).
		Build()
}

func onSessionChange(elem interface{}, action ifs.Action, notify bool, vnic ifs.IVNic) (interface{}, bool, error) {
	session := elem.(*learn.LearningSession)

	if (action == ifs.PUT || action == ifs.PATCH) &&
		session.Status == learn.SessionStatus_SESSION_STATUS_COMPLETED {
		// Session completed — trigger adaptive engine pipeline:
		// 1. Update SkillMastery records for all skills in this session's interactions
		// 2. Trigger adaptive engine to evaluate rules and refresh LearningPath
		// 3. Update EngagementMetric for this student
		//
		// engine := engine.NewAdaptiveEngine(vnic)
		// path := loadPath(session.PathId, vnic)
		// student := loadStudent(session.StudentId, vnic)
		// engine.OnActivityCompleted(session, path, student)
		_ = session
	}

	return nil, true, nil
}
