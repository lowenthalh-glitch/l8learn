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
)

func generateCourses() []*learn.Course {
	var list []*learn.Course
	subjects := []learn.SubjectType{
		learn.SubjectType_SUBJECT_TYPE_MATH, learn.SubjectType_SUBJECT_TYPE_READING,
		learn.SubjectType_SUBJECT_TYPE_SCIENCE, learn.SubjectType_SUBJECT_TYPE_WRITING,
	}
	for i := 0; i < 20; i++ {
		list = append(list, &learn.Course{
			CourseId:       fmt.Sprintf("CRS-%03d", i+1),
			Name:          CourseNames[i%len(CourseNames)],
			Subject:       subjects[i%4],
			MinGrade:      learn.GradeLevel(2 + i/5),
			MaxGrade:      learn.GradeLevel(3 + i/5),
			Status:        learn.ContentStatus_CONTENT_STATUS_PUBLISHED,
			EstimatedHours: int32(20 + rand.Intn(40)),
		})
	}
	return list
}

func generateUnits(store *MockDataStore) []*learn.Unit {
	var list []*learn.Unit
	id := 1
	for i := 0; i < len(store.CourseIDs); i++ {
		for u := 0; u < 5; u++ {
			list = append(list, &learn.Unit{
				UnitId:           fmt.Sprintf("UNT-%03d", id),
				CourseId:         store.CourseIDs[i],
				Name:             fmt.Sprintf("Unit %d", u+1),
				SequenceOrder:    int32(u + 1),
				Status:           learn.ContentStatus_CONTENT_STATUS_PUBLISHED,
				EstimatedMinutes: int32(60 + rand.Intn(120)),
			})
			id++
		}
	}
	return list
}

func generateLessons(store *MockDataStore) []*learn.Lesson {
	var list []*learn.Lesson
	id := 1
	difficulties := []learn.DifficultyLevel{
		learn.DifficultyLevel_DIFFICULTY_LEVEL_EASY,
		learn.DifficultyLevel_DIFFICULTY_LEVEL_MEDIUM,
		learn.DifficultyLevel_DIFFICULTY_LEVEL_MEDIUM,
		learn.DifficultyLevel_DIFFICULTY_LEVEL_HARD,
	}
	for i := 0; i < len(store.UnitIDs); i++ {
		for l := 0; l < 4; l++ {
			list = append(list, &learn.Lesson{
				LessonId:         fmt.Sprintf("LSN-%03d", id),
				UnitId:           store.UnitIDs[i],
				Name:             fmt.Sprintf("Lesson %d", l+1),
				SequenceOrder:    int32(l + 1),
				Difficulty:       difficulties[l],
				Status:           learn.ContentStatus_CONTENT_STATUS_PUBLISHED,
				EstimatedMinutes: int32(10 + rand.Intn(20)),
			})
			id++
		}
	}
	return list
}

func generateActivities(store *MockDataStore) []*learn.Activity {
	var list []*learn.Activity
	id := 1
	types := []learn.ActivityType{
		learn.ActivityType_ACTIVITY_TYPE_MULTIPLE_CHOICE,
		learn.ActivityType_ACTIVITY_TYPE_INTERACTIVE,
		learn.ActivityType_ACTIVITY_TYPE_GAME,
	}
	for i := 0; i < len(store.LessonIDs); i++ {
		for a := 0; a < 3; a++ {
			actId := fmt.Sprintf("ACT-%04d", id)
			activity := &learn.Activity{
				ActivityId:       actId,
				LessonId:         store.LessonIDs[i],
				Name:             ActivityNames[(i+a)%len(ActivityNames)],
				ActivityType:     types[a],
				Difficulty:       learn.DifficultyLevel(2 + rand.Intn(3)),
				Status:           learn.ContentStatus_CONTENT_STATUS_PUBLISHED,
				PointsPossible:   int32(10 + rand.Intn(20)),
				EstimatedSeconds: int32(60 + rand.Intn(180)),
				HintsEnabled:     true,
				HintCount:        2,
				Questions:        generateQuestions(actId, 3+rand.Intn(3)),
			}
			list = append(list, activity)
			id++
		}
	}
	return list
}

func generateQuestions(activityId string, count int) []*learn.Question {
	var questions []*learn.Question
	for q := 0; q < count; q++ {
		question := &learn.Question{
			QuestionId:   fmt.Sprintf("%s-Q%d", activityId, q+1),
			SequenceOrder: int32(q + 1),
			QuestionType: learn.QuestionType_QUESTION_TYPE_SINGLE_CHOICE,
			PromptText:   fmt.Sprintf("Question %d for this activity?", q+1),
			Points:       10,
			Options: []*learn.AnswerOption{
				{OptionId: fmt.Sprintf("opt-a-%d", q), Text: "Option A", IsCorrect: q%4 == 0},
				{OptionId: fmt.Sprintf("opt-b-%d", q), Text: "Option B", IsCorrect: q%4 == 1},
				{OptionId: fmt.Sprintf("opt-c-%d", q), Text: "Option C", IsCorrect: q%4 == 2},
				{OptionId: fmt.Sprintf("opt-d-%d", q), Text: "Option D", IsCorrect: q%4 == 3},
			},
			HintTexts: []string{"Think about it carefully.", "Try eliminating wrong answers."},
		}
		questions = append(questions, question)
	}
	return questions
}
