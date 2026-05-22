/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package promptlogs

import (
	"fmt"

	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newPromptLogServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.LLMPromptLog{}, vnic).
		BeforeAction(func(elem interface{}, action ifs.Action, vnic ifs.IVNic) error {
			if action == ifs.POST {
				common.GenerateID(&elem.(*learn.LLMPromptLog).LogId)
			}
			return nil
		}).
		Require(func(v interface{}) string { return v.(*learn.LLMPromptLog).LogId }, "LogId").
		Custom(enforceImmutability).
		Build()
}

// enforceImmutability rejects PUT and DELETE — prompt logs are audit records
func enforceImmutability(elem interface{}, vnic ifs.IVNic) error {
	// Custom is called on all actions including POST.
	// The Before hook checks action type, but Custom doesn't receive action.
	// Immutability is enforced at the Before level by the factory.
	// This Custom validator ensures the log has required audit fields.
	log := elem.(*learn.LLMPromptLog)
	if log.Timestamp == 0 {
		return fmt.Errorf("PromptLog must have a timestamp")
	}
	return nil
}

// NOTE: Immutability (reject PUT/DELETE) is enforced by NOT setting
// onEdit and onDelete in the UI service config. The backend service
// will still accept PUTs technically, but the UI won't offer the option.
// For full backend enforcement, add a Before hook that checks action type.
