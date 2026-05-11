/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package ui

import (
	l8c "github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func registerCollabTypes(resources ifs.IResources) {
	
	l8c.RegisterType(resources, &learn.CollabGroup{}, &learn.CollabGroupList{}, "GroupId")
	l8c.RegisterType(resources, &learn.CollabMessage{}, &learn.CollabMessageList{}, "MessageId")
	l8c.RegisterType(resources, &learn.TutorMatch{}, &learn.TutorMatchList{}, "MatchId")
	l8c.RegisterType(resources, &learn.Challenge{}, &learn.ChallengeList{}, "ChallengeId")
}
