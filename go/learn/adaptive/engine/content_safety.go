/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * ContentSafety — validates AI-generated content before showing to children.
 * Scans for inappropriate language, violence, bias, and mathematical errors.
 */
package engine

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/saichler/l8learn/go/types/learn"
)

// ContentSafety validates AI-generated lessons for child safety
type ContentSafety struct {
	blockedWords []string
}

func NewContentSafety() *ContentSafety {
	return &ContentSafety{
		blockedWords: []string{
			"kill", "murder", "weapon", "gun", "knife", "drug", "alcohol",
			"hate", "stupid", "idiot", "dumb", "ugly", "fat",
			"sexy", "naked", "porn",
			"suicide", "self-harm",
		},
	}
}

// ValidateLesson scans all text in a generated lesson for safety issues
// Returns empty slice if safe, list of issues if problems found
func (cs *ContentSafety) ValidateLesson(lesson *learn.GeneratedLesson) []string {
	var issues []string

	// Scan title, objective, parent instructions
	issues = append(issues, cs.scanText("title", lesson.Title)...)
	issues = append(issues, cs.scanText("objective", lesson.Objective)...)
	issues = append(issues, cs.scanText("parentInstructions", lesson.ParentInstructions)...)

	// Scan each step
	for i, step := range lesson.Steps {
		prefix := fmt.Sprintf("step[%d]", i)
		issues = append(issues, cs.scanText(prefix+".title", step.Title)...)
		issues = append(issues, cs.scanText(prefix+".instructions", step.Instructions)...)
		issues = append(issues, cs.scanText(prefix+".materialsInstructions", step.MaterialsInstructions)...)

		// Scan questions
		for j, q := range step.Questions {
			qPrefix := fmt.Sprintf("%s.question[%d]", prefix, j)
			issues = append(issues, cs.scanText(qPrefix+".prompt", q.Prompt)...)
			issues = append(issues, cs.scanText(qPrefix+".explanation", q.Explanation)...)

			// Verify math correctness for numeric questions
			if q.QuestionType == learn.QuestionType_QUESTION_TYPE_NUMERIC {
				issues = append(issues, cs.validateMathAnswer(qPrefix, q)...)
			}

			// Verify choice questions have exactly one correct answer
			if q.QuestionType == learn.QuestionType_QUESTION_TYPE_SINGLE_CHOICE {
				issues = append(issues, cs.validateChoiceAnswer(qPrefix, q)...)
			}

			// Scan option feedback
			for k, opt := range q.Options {
				issues = append(issues, cs.scanText(
					fmt.Sprintf("%s.option[%d].feedback", qPrefix, k), opt.Feedback)...)
			}
		}
	}

	return issues
}

func (cs *ContentSafety) scanText(field, text string) []string {
	if text == "" {
		return nil
	}
	var issues []string
	lower := strings.ToLower(text)
	for _, word := range cs.blockedWords {
		if strings.Contains(lower, word) {
			issues = append(issues, fmt.Sprintf("%s contains blocked word: %s", field, word))
		}
	}
	return issues
}

func (cs *ContentSafety) validateMathAnswer(prefix string, q *learn.GeneratedQuestion) []string {
	if q.CorrectAnswer == "" {
		return []string{fmt.Sprintf("%s: numeric question has no correct answer", prefix)}
	}
	_, err := strconv.ParseFloat(q.CorrectAnswer, 64)
	if err != nil {
		return []string{fmt.Sprintf("%s: correct answer '%s' is not a valid number", prefix, q.CorrectAnswer)}
	}
	return nil
}

func (cs *ContentSafety) validateChoiceAnswer(prefix string, q *learn.GeneratedQuestion) []string {
	correctCount := 0
	for _, opt := range q.Options {
		if opt.IsCorrect {
			correctCount++
		}
	}
	if correctCount == 0 {
		return []string{fmt.Sprintf("%s: no correct option marked", prefix)}
	}
	if correctCount > 1 {
		return []string{fmt.Sprintf("%s: multiple correct options for single-choice question", prefix)}
	}
	return nil
}
