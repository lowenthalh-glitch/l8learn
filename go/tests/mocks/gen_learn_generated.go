/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package mocks

import (
	"fmt"
	"github.com/saichler/l8learn/go/types/learn"
	"math/rand"
	"time"
)

func generateGeneratedLessons(store *MockDataStore) []*learn.GeneratedLesson {
	var list []*learn.GeneratedLesson
	now := time.Now().Unix()

	themes := []string{"Dinosaurs", "Space", "Animals", "Sports", "Robots", "Ocean", "Music", "Nature"}
	topics := []string{"Multiplication by 7s", "Fraction Comparison", "Reading Comprehension", "Addition Facts", "Subtraction Regrouping"}
	strategies := []string{"scaffold", "alternate", "break", "review"}

	statuses := []learn.GeneratedLessonStatus{
		learn.GeneratedLessonStatus_GENERATED_LESSON_STATUS_READY,
		learn.GeneratedLessonStatus_GENERATED_LESSON_STATUS_COMPLETED,
		learn.GeneratedLessonStatus_GENERATED_LESSON_STATUS_COMPLETED,
		learn.GeneratedLessonStatus_GENERATED_LESSON_STATUS_IN_PROGRESS,
	}

	for i := 0; i < 20; i++ {
		theme := themes[i%len(themes)]
		topic := topics[i%len(topics)]
		status := statuses[i%len(statuses)]
		correct := int32(0)
		total := int32(4)
		if status == learn.GeneratedLessonStatus_GENERATED_LESSON_STATUS_COMPLETED {
			correct = int32(2 + rand.Intn(3))
		}

		lesson := &learn.GeneratedLesson{
			GeneratedLessonId: fmt.Sprintf("GL-%04d", i+1),
			StudentId:         pickRef(store.StudentIDs, i),
			PathId:            pickRef(store.PathIDs, i),
			Status:            status,
			SkillIds:          []string{pickRef(store.SkillIDs, i*2)},
			Subject:           learn.SubjectType(1 + i%2),
			Difficulty:        learn.DifficultyLevel(2 + rand.Intn(3)),
			Topic:             topic,
			Theme:             theme,
			Title:             fmt.Sprintf("%s %s Adventure", theme, topic),
			Objective:         fmt.Sprintf("Student will practice %s using %s theme", topic, theme),
			EstimatedMinutes:  int32(8 + rand.Intn(7)),
			MaterialsNeeded:   []string{"pencil", "paper"},
			ParentInstructions: "Observe and encourage. Help with physical materials if needed.",
			MinCorrectToAdvance: 3,
			MinCorrectToPass:   2,
			OnStruggleStrategy: strategies[i%len(strategies)],
			QuestionsCorrect:   correct,
			QuestionsTotal:     total,
			GeneratedAt:        now - int64(rand.Intn(7*24*3600)),
			Steps: []*learn.LessonStep{
				{
					StepNumber:   1,
					StepType:     "screen",
					Title:        "Practice Problems",
					Instructions: fmt.Sprintf("Answer these %s questions", topic),
					DurationMinutes: int32(8 + rand.Intn(5)),
					ParentRole:   "none",
					Questions: []*learn.GeneratedQuestion{
						{
							QuestionId:    fmt.Sprintf("GQ-%04d-1", i+1),
							Prompt:        "What is 7 × 8?",
							QuestionType:  learn.QuestionType_QUESTION_TYPE_NUMERIC,
							CorrectAnswer: "56",
							Explanation:   "7 groups of 8 = 56",
							Hints:         []string{"Try counting by 7s", "7×8 is close to 7×7+7"},
							Difficulty:    learn.DifficultyLevel_DIFFICULTY_LEVEL_MEDIUM,
						},
						{
							QuestionId:    fmt.Sprintf("GQ-%04d-2", i+1),
							Prompt:        "What is 6 × 9?",
							QuestionType:  learn.QuestionType_QUESTION_TYPE_NUMERIC,
							CorrectAnswer: "54",
							Explanation:   "6 groups of 9 = 54. Try 6×10 minus 6.",
							Hints:         []string{"Try 6 × 10 first, then subtract 6"},
							Difficulty:    learn.DifficultyLevel_DIFFICULTY_LEVEL_MEDIUM,
						},
					},
				},
			},
		}
		list = append(list, lesson)
	}
	return list
}
