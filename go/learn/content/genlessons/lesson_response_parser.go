/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package genlessons

import (
	"fmt"
	"github.com/saichler/l8learn/go/learn/adaptive/engine"
	"github.com/saichler/l8learn/go/types/learn"
)

type llmLessonResponse struct {
	Title              string          `json:"title"`
	Objective          string          `json:"objective"`
	Theme              string          `json:"theme"`
	EstimatedMinutes   int32           `json:"estimatedMinutes"`
	MaterialsNeeded    []string        `json:"materialsNeeded"`
	ParentInstructions string          `json:"parentInstructions"`
	Steps              []llmStep       `json:"steps"`
	MinCorrectToAdvance int32          `json:"minCorrectToAdvance"`
	MinCorrectToPass   int32           `json:"minCorrectToPass"`
	OnStruggleStrategy string          `json:"onStruggleStrategy"`
}

type llmStep struct {
	StepNumber            int32          `json:"stepNumber"`
	StepType              string         `json:"stepType"`
	Title                 string         `json:"title"`
	Instructions          string         `json:"instructions"`
	DurationMinutes       int32          `json:"durationMinutes"`
	ParentRole            string         `json:"parentRole"`
	MaterialsInstructions string         `json:"materialsInstructions"`
	Questions             []llmQuestion  `json:"questions"`
}

type llmQuestion struct {
	Prompt        string      `json:"prompt"`
	QuestionType  int32       `json:"questionType"`
	CorrectAnswer string      `json:"correctAnswer"`
	Explanation   string      `json:"explanation"`
	Hints         []string    `json:"hints"`
	Difficulty    int32       `json:"difficulty"`
	Options       []llmOption `json:"options"`
}

type llmOption struct {
	Text      string `json:"text"`
	IsCorrect bool   `json:"isCorrect"`
	Feedback  string `json:"feedback"`
}

// ParseLessonResponse parses Claude's JSON into a GeneratedLesson proto.
func ParseLessonResponse(jsonResponse string) (*learn.GeneratedLesson, error) {
	var resp llmLessonResponse
	if err := engine.ParseLLMResponse(jsonResponse, &resp); err != nil {
		return nil, fmt.Errorf("lesson parse: %w", err)
	}

	lesson := &learn.GeneratedLesson{
		Title:               resp.Title,
		Objective:           resp.Objective,
		Theme:               resp.Theme,
		EstimatedMinutes:    resp.EstimatedMinutes,
		MaterialsNeeded:     resp.MaterialsNeeded,
		ParentInstructions:  resp.ParentInstructions,
		MinCorrectToAdvance: resp.MinCorrectToAdvance,
		MinCorrectToPass:    resp.MinCorrectToPass,
		OnStruggleStrategy:  resp.OnStruggleStrategy,
	}

	for _, s := range resp.Steps {
		step := &learn.LessonStep{
			StepNumber:            s.StepNumber,
			StepType:              s.StepType,
			Title:                 s.Title,
			Instructions:          s.Instructions,
			DurationMinutes:       s.DurationMinutes,
			ParentRole:            s.ParentRole,
			MaterialsInstructions: s.MaterialsInstructions,
		}
		for qi, q := range s.Questions {
			question := &learn.GeneratedQuestion{
				QuestionId:    fmt.Sprintf("Q%d-%d", s.StepNumber, qi+1),
				Prompt:        q.Prompt,
				QuestionType:  learn.QuestionType(q.QuestionType),
				CorrectAnswer: q.CorrectAnswer,
				Explanation:   q.Explanation,
				Hints:         q.Hints,
				Difficulty:    learn.DifficultyLevel(q.Difficulty),
			}
			for oi, o := range q.Options {
				question.Options = append(question.Options, &learn.GeneratedOption{
					OptionId:  fmt.Sprintf("O%d-%d-%d", s.StepNumber, qi+1, oi+1),
					Text:      o.Text,
					IsCorrect: o.IsCorrect,
					Feedback:  o.Feedback,
				})
			}
			step.Questions = append(step.Questions, question)
		}
		lesson.Steps = append(lesson.Steps, step)
	}

	return lesson, nil
}
