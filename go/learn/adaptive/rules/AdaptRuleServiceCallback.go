/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package rules

import (
	"fmt"

	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newAdaptRuleServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.AdaptationRule{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.AdaptationRule).RuleId }, "RuleId").
		Require(func(v interface{}) string { return v.(*learn.AdaptationRule).Name }, "Name").
		Custom(validateTriggerStrategy).
		Build()
}

func validateTriggerStrategy(elem interface{}, vnic ifs.IVNic) error {
	rule := elem.(*learn.AdaptationRule)

	// SCORE_ABOVE should not trigger REVIEW (reviewing easier material when scoring high makes no sense)
	if rule.Trigger == learn.AdaptTrigger_ADAPT_TRIGGER_SCORE_ABOVE &&
		rule.Strategy == learn.AdaptStrategy_ADAPT_STRATEGY_REVIEW {
		return fmt.Errorf("invalid combination: SCORE_ABOVE trigger cannot use REVIEW strategy")
	}

	// STREAK_CORRECT should not trigger REPEAT (repeating when getting everything right is wasteful)
	if rule.Trigger == learn.AdaptTrigger_ADAPT_TRIGGER_STREAK_CORRECT &&
		rule.Strategy == learn.AdaptStrategy_ADAPT_STRATEGY_REPEAT {
		return fmt.Errorf("invalid combination: STREAK_CORRECT trigger cannot use REPEAT strategy")
	}

	return nil
}
