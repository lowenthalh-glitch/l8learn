/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package challenges

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newChallengeServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.Challenge{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.Challenge).ChallengeId }, "ChallengeId").
		Require(func(v interface{}) string { return v.(*learn.Challenge).Name }, "Name").
		Require(func(v interface{}) string { return v.(*learn.Challenge).ClassroomId }, "ClassroomId").
		After(onChallengeCreate).
		Build()
}

func onChallengeCreate(elem interface{}, action ifs.Action, notify bool, vnic ifs.IVNic) (interface{}, bool, error) {
	challenge := elem.(*learn.Challenge)
	if action == ifs.POST {
		// Challenge created by teacher:
		// 1. Auto-assign students to balanced teams (mix of mastery levels)
		// 2. Create team CollabGroups (one per team)
		// 3. Notify all participating students
		// 4. Generate daily leaderboard update schedule
		_ = challenge
	}
	return nil, true, nil
}
