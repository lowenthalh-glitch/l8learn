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

func registerAdaptiveTypes(resources ifs.IResources) {
	
	l8c.RegisterType(resources, &learn.LearningPath{}, &learn.LearningPathList{}, "PathId")
	l8c.RegisterType(resources, &learn.SkillMastery{}, &learn.SkillMasteryList{}, "MasteryId")
	l8c.RegisterType(resources, &learn.Skill{}, &learn.SkillList{}, "SkillId")
	l8c.RegisterType(resources, &learn.AdaptationRule{}, &learn.AdaptationRuleList{}, "RuleId")
	l8c.RegisterType(resources, &learn.DailySchedule{}, &learn.DailyScheduleList{}, "ScheduleId")
}
