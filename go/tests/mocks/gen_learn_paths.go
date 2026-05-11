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

func generatePaths(store *MockDataStore) []*learn.LearningPath {
	var list []*learn.LearningPath
	now := time.Now().Unix()
	for i := 0; i < 1000; i++ {
		list = append(list, &learn.LearningPath{
			PathId:              fmt.Sprintf("PTH-%04d", i+1),
			StudentId:           pickRef(store.StudentIDs, i),
			Subject:            learn.SubjectType(1 + i%2), // Math or Reading
			TargetGrade:        learn.GradeLevel(3 + rand.Intn(4)),
			Status:             learn.PathStatus_PATH_STATUS_ACTIVE,
			CurrentActivityId:  pickRef(store.ActivityIDs, i*3),
			CurrentSkillId:     pickRef(store.SkillIDs, i*2),
			CurrentDifficulty:  learn.DifficultyLevel(2 + rand.Intn(3)),
			ActivitiesCompleted: int32(5 + rand.Intn(50)),
			SkillsMastered:     int32(1 + rand.Intn(10)),
			TotalTimeSeconds:   int32(600 + rand.Intn(7200)),
			StartedDate:        now - int64(rand.Intn(90*24*3600)),
			LastActive:         now - int64(rand.Intn(7*24*3600)),
		})
	}
	return list
}

func generateMastery(store *MockDataStore) []*learn.SkillMastery {
	var list []*learn.SkillMastery
	now := time.Now().Unix()
	levels := []learn.MasteryLevel{
		learn.MasteryLevel_MASTERY_LEVEL_EMERGING,
		learn.MasteryLevel_MASTERY_LEVEL_DEVELOPING,
		learn.MasteryLevel_MASTERY_LEVEL_DEVELOPING,
		learn.MasteryLevel_MASTERY_LEVEL_PROFICIENT,
		learn.MasteryLevel_MASTERY_LEVEL_PROFICIENT,
		learn.MasteryLevel_MASTERY_LEVEL_MASTERED,
	}
	id := 1
	for i := 0; i < 1000; i++ {
		numSkills := 3 + rand.Intn(5)
		for s := 0; s < numSkills; s++ {
			level := levels[rand.Intn(len(levels))]
			accuracy := 0.3 + rand.Float64()*0.6
			list = append(list, &learn.SkillMastery{
				MasteryId:       fmt.Sprintf("MST-%05d", id),
				StudentId:       pickRef(store.StudentIDs, i),
				SkillId:         pickRef(store.SkillIDs, i*5+s),
				Level:           level,
				Confidence:      0.5 + rand.Float64()*0.5,
				AttemptsCount:   int32(5 + rand.Intn(30)),
				CorrectCount:    int32(3 + rand.Intn(20)),
				CurrentAccuracy: accuracy,
				FirstAttempted:  now - int64(rand.Intn(60*24*3600)),
				LastAttempted:   now - int64(rand.Intn(7*24*3600)),
				TotalTimeSeconds: int32(120 + rand.Intn(1800)),
			})
			id++
		}
	}
	return list
}
