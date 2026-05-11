/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package activities

import (
	"fmt"

	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newActivityServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.Activity{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.Activity).ActivityId }, "ActivityId").
		Require(func(v interface{}) string { return v.(*learn.Activity).LessonId }, "LessonId").
		Require(func(v interface{}) string { return v.(*learn.Activity).Name }, "Name").
		Custom(validateQuestions).
		Build()
}

func validateQuestions(elem interface{}, action ifs.Action, vnic ifs.IVNic) error {
	activity := elem.(*learn.Activity)
	if activity.Questions == nil || len(activity.Questions) == 0 {
		return nil
	}

	ids := make(map[string]bool)
	for i, q := range activity.Questions {
		if q.QuestionId == "" {
			return fmt.Errorf("question at index %d has no QuestionId", i)
		}
		if ids[q.QuestionId] {
			return fmt.Errorf("duplicate QuestionId: %s", q.QuestionId)
		}
		ids[q.QuestionId] = true

		if q.QuestionType == learn.QuestionType_QUESTION_TYPE_SINGLE_CHOICE ||
			q.QuestionType == learn.QuestionType_QUESTION_TYPE_MULTI_CHOICE {
			hasCorrect := false
			for _, opt := range q.Options {
				if opt.IsCorrect {
					hasCorrect = true
					break
				}
			}
			if !hasCorrect {
				return fmt.Errorf("question %s has no correct answer option", q.QuestionId)
			}
		}
	}
	return nil
}
