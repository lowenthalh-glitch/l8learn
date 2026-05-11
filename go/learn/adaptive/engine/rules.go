/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package engine

import (
	"github.com/lowenthalh-glitch/l8learn/go/types/learn"
)

// RuleResult holds the outcome of evaluating a single rule
type RuleResult struct {
	Rule     *learn.AdaptationRule
	Matched  bool
	Strategy learn.AdaptStrategy
}

// EvaluateRules checks all active rules against the student's recent interaction data.
// Returns the highest-priority matching rule's strategy, or nil if no rule matches.
func EvaluateRules(
	rules []*learn.AdaptationRule,
	interactions []*learn.Interaction,
	mastery *learn.SkillMastery,
) *RuleResult {
	for _, rule := range rules {
		if rule.Status != learn.AdaptRuleStatus_ADAPT_RULE_STATUS_ACTIVE {
			continue
		}
		if matchesRule(rule, interactions, mastery) {
			return &RuleResult{
				Rule:     rule,
				Matched:  true,
				Strategy: rule.Strategy,
			}
		}
	}
	return nil
}

func matchesRule(
	rule *learn.AdaptationRule,
	interactions []*learn.Interaction,
	mastery *learn.SkillMastery,
) bool {
	if len(interactions) == 0 {
		return false
	}

	window := int(rule.TriggerWindow)
	if window <= 0 || window > len(interactions) {
		window = len(interactions)
	}
	recent := interactions[len(interactions)-window:]

	switch rule.Trigger {
	case learn.AdaptTrigger_ADAPT_TRIGGER_SCORE_BELOW:
		return checkScoreBelow(recent, rule.TriggerThreshold)
	case learn.AdaptTrigger_ADAPT_TRIGGER_SCORE_ABOVE:
		return checkScoreAbove(recent, rule.TriggerThreshold)
	case learn.AdaptTrigger_ADAPT_TRIGGER_STREAK_CORRECT:
		return checkStreakCorrect(recent, rule.TriggerThreshold)
	case learn.AdaptTrigger_ADAPT_TRIGGER_STREAK_INCORRECT:
		return checkStreakIncorrect(recent, rule.TriggerThreshold)
	case learn.AdaptTrigger_ADAPT_TRIGGER_TIME_EXCEEDED:
		return checkTimeExceeded(recent, rule.TriggerThreshold)
	case learn.AdaptTrigger_ADAPT_TRIGGER_TIME_TOO_FAST:
		return checkTimeTooFast(recent, rule.TriggerThreshold)
	case learn.AdaptTrigger_ADAPT_TRIGGER_HINTS_EXHAUSTED:
		return checkHintsExhausted(recent)
	case learn.AdaptTrigger_ADAPT_TRIGGER_ENGAGEMENT_DROP:
		return checkEngagementDrop(recent, rule.TriggerThreshold)
	case learn.AdaptTrigger_ADAPT_TRIGGER_MASTERY_ACHIEVED:
		return checkMasteryAchieved(mastery, rule.TriggerThreshold)
	}
	return false
}

func checkScoreBelow(interactions []*learn.Interaction, threshold int32) bool {
	correct := 0
	for _, i := range interactions {
		if i.Result == learn.InteractionResult_INTERACTION_RESULT_CORRECT {
			correct++
		}
	}
	pct := int32(correct * 100 / len(interactions))
	return pct < threshold
}

func checkScoreAbove(interactions []*learn.Interaction, threshold int32) bool {
	correct := 0
	for _, i := range interactions {
		if i.Result == learn.InteractionResult_INTERACTION_RESULT_CORRECT {
			correct++
		}
	}
	pct := int32(correct * 100 / len(interactions))
	return pct > threshold
}

func checkStreakCorrect(interactions []*learn.Interaction, threshold int32) bool {
	streak := int32(0)
	for j := len(interactions) - 1; j >= 0; j-- {
		if interactions[j].Result == learn.InteractionResult_INTERACTION_RESULT_CORRECT {
			streak++
		} else {
			break
		}
	}
	return streak >= threshold
}

func checkStreakIncorrect(interactions []*learn.Interaction, threshold int32) bool {
	streak := int32(0)
	for j := len(interactions) - 1; j >= 0; j-- {
		if interactions[j].Result == learn.InteractionResult_INTERACTION_RESULT_INCORRECT {
			streak++
		} else {
			break
		}
	}
	return streak >= threshold
}

func checkTimeExceeded(interactions []*learn.Interaction, thresholdSecs int32) bool {
	if len(interactions) == 0 {
		return false
	}
	last := interactions[len(interactions)-1]
	return last.TimeSpentSeconds > thresholdSecs
}

func checkTimeTooFast(interactions []*learn.Interaction, thresholdSecs int32) bool {
	if len(interactions) == 0 {
		return false
	}
	last := interactions[len(interactions)-1]
	return last.TimeSpentSeconds > 0 && last.TimeSpentSeconds < thresholdSecs
}

func checkHintsExhausted(interactions []*learn.Interaction) bool {
	if len(interactions) == 0 {
		return false
	}
	last := interactions[len(interactions)-1]
	return last.HintsUsed > 0 && last.Result != learn.InteractionResult_INTERACTION_RESULT_CORRECT
}

func checkEngagementDrop(interactions []*learn.Interaction, threshold int32) bool {
	if len(interactions) < 4 {
		return false
	}
	// Compare average time-per-question in first half vs second half
	mid := len(interactions) / 2
	firstHalf := avgTimeSpent(interactions[:mid])
	secondHalf := avgTimeSpent(interactions[mid:])
	if firstHalf == 0 {
		return false
	}
	// If second half takes significantly longer, engagement may be dropping
	ratio := int32(secondHalf * 100 / firstHalf)
	return ratio > (100 + threshold)
}

func checkMasteryAchieved(mastery *learn.SkillMastery, threshold int32) bool {
	if mastery == nil {
		return false
	}
	return int32(mastery.CurrentAccuracy*100) >= threshold
}

func avgTimeSpent(interactions []*learn.Interaction) int32 {
	if len(interactions) == 0 {
		return 0
	}
	total := int32(0)
	for _, i := range interactions {
		total += i.TimeSpentSeconds
	}
	return total / int32(len(interactions))
}
