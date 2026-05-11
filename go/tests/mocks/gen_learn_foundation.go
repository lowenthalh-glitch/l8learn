/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package mocks

import (
	"fmt"
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
)

func generateDistricts() []*learn.District {
	var list []*learn.District
	for i := 0; i < 5; i++ {
		list = append(list, &learn.District{
			DistrictId:    fmt.Sprintf("DIST-%03d", i+1),
			Name:          DistrictNames[i],
			StateProvince: StateCodes[i%len(StateCodes)],
			CountryCode:   "US",
			LicenseTier:   "premium",
			MaxStudents:   5000,
		})
	}
	return list
}

func generateSchools(store *MockDataStore) []*learn.School {
	var list []*learn.School
	for i := 0; i < 20; i++ {
		list = append(list, &learn.School{
			SchoolId:      fmt.Sprintf("SCH-%03d", i+1),
			Name:          SchoolNames[i%len(SchoolNames)],
			DistrictId:    pickRef(store.DistrictIDs, i/4),
			StateProvince: StateCodes[i%len(StateCodes)],
			CountryCode:   "US",
			Timezone:      "America/Los_Angeles",
		})
	}
	return list
}

func generateSkills() []*learn.Skill {
	var list []*learn.Skill
	id := 1

	// Math skills across grades
	for g := 2; g <= 8; g++ {
		for _, name := range MathSkillNames[:12] {
			list = append(list, &learn.Skill{
				SkillId:              fmt.Sprintf("SKL-%03d", id),
				Name:                 name,
				Subject:             learn.SubjectType_SUBJECT_TYPE_MATH,
				GradeLevel:          learn.GradeLevel(g),
				Domain:              SkillDomains[id%8],
				TypicalMasteryMinutes: int32(30 + (id%5)*15),
			})
			id++
		}
	}

	// Reading skills across grades
	for g := 2; g <= 7; g++ {
		for _, name := range ReadingSkillNames[:10] {
			list = append(list, &learn.Skill{
				SkillId:              fmt.Sprintf("SKL-%03d", id),
				Name:                 name,
				Subject:             learn.SubjectType_SUBJECT_TYPE_READING,
				GradeLevel:          learn.GradeLevel(g),
				Domain:              SkillDomains[8+id%6],
				TypicalMasteryMinutes: int32(25 + (id%4)*10),
			})
			id++
		}
	}

	return list
}

func generateAdaptRules() []*learn.AdaptationRule {
	rules := []*learn.AdaptationRule{
		{RuleId: "RULE-001", Name: "Struggle: lower difficulty", Status: learn.AdaptRuleStatus_ADAPT_RULE_STATUS_ACTIVE,
			Priority: 1, Trigger: learn.AdaptTrigger_ADAPT_TRIGGER_STREAK_INCORRECT, TriggerThreshold: 3, TriggerWindow: 5,
			Strategy: learn.AdaptStrategy_ADAPT_STRATEGY_REPEAT},
		{RuleId: "RULE-002", Name: "Mastery: advance", Status: learn.AdaptRuleStatus_ADAPT_RULE_STATUS_ACTIVE,
			Priority: 2, Trigger: learn.AdaptTrigger_ADAPT_TRIGGER_SCORE_ABOVE, TriggerThreshold: 90, TriggerWindow: 10,
			Strategy: learn.AdaptStrategy_ADAPT_STRATEGY_ADVANCE},
		{RuleId: "RULE-003", Name: "Streak: enrich", Status: learn.AdaptRuleStatus_ADAPT_RULE_STATUS_ACTIVE,
			Priority: 3, Trigger: learn.AdaptTrigger_ADAPT_TRIGGER_STREAK_CORRECT, TriggerThreshold: 5, TriggerWindow: 5,
			Strategy: learn.AdaptStrategy_ADAPT_STRATEGY_ENRICH},
		{RuleId: "RULE-004", Name: "Bored: try different type", Status: learn.AdaptRuleStatus_ADAPT_RULE_STATUS_ACTIVE,
			Priority: 4, Trigger: learn.AdaptTrigger_ADAPT_TRIGGER_TIME_TOO_FAST, TriggerThreshold: 3, TriggerWindow: 5,
			Strategy: learn.AdaptStrategy_ADAPT_STRATEGY_ALTERNATE},
		{RuleId: "RULE-005", Name: "Stuck: scaffold", Status: learn.AdaptRuleStatus_ADAPT_RULE_STATUS_ACTIVE,
			Priority: 5, Trigger: learn.AdaptTrigger_ADAPT_TRIGGER_SCORE_BELOW, TriggerThreshold: 30, TriggerWindow: 10,
			Strategy: learn.AdaptStrategy_ADAPT_STRATEGY_SCAFFOLD},
		{RuleId: "RULE-006", Name: "Engagement drop: game break", Status: learn.AdaptRuleStatus_ADAPT_RULE_STATUS_ACTIVE,
			Priority: 6, Trigger: learn.AdaptTrigger_ADAPT_TRIGGER_ENGAGEMENT_DROP, TriggerThreshold: 50, TriggerWindow: 8,
			Strategy: learn.AdaptStrategy_ADAPT_STRATEGY_BREAK},
		{RuleId: "RULE-007", Name: "Hints exhausted: review prereqs", Status: learn.AdaptRuleStatus_ADAPT_RULE_STATUS_ACTIVE,
			Priority: 7, Trigger: learn.AdaptTrigger_ADAPT_TRIGGER_HINTS_EXHAUSTED, TriggerThreshold: 0, TriggerWindow: 3,
			Strategy: learn.AdaptStrategy_ADAPT_STRATEGY_REVIEW},
	}
	return rules
}
