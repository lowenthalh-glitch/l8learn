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

func generateProfiles(store *MockDataStore) []*learn.StudentProfile {
	var list []*learn.StudentProfile
	now := time.Now().Unix()

	modes := []string{"visual", "auditory", "kinesthetic", "reading"}
	times := []string{"morning", "midday", "afternoon"}
	interests := []string{"dinosaurs", "space", "animals", "sports", "art", "music", "robots", "nature"}
	rewards := []string{"badges", "streaks", "avatar_items", "points"}
	focusSigns := []string{"rushing", "random_answers", "long_pauses", "fidgeting"}
	supports := []string{"timer_visible", "chunked_problems", "movement_breaks", "fidget_tool"}
	readingLevels := []string{"pre-reader", "beginning", "developing", "fluent"}
	mathLevels := []string{"pre-number", "counting", "operations", "fractions"}
	personalities := []string{"encouraging", "patient", "playful", "calm"}

	for i := 0; i < len(store.StudentIDs) && i < 100; i++ {
		profile := &learn.StudentProfile{
			ProfileId:          fmt.Sprintf("PROF-%04d", i+1),
			StudentId:          store.StudentIDs[i],
			CreatedDate:        now - int64(rand.Intn(90*24*3600)),
			LastUpdated:        now - int64(rand.Intn(7*24*3600)),
			ShortSummary: fmt.Sprintf("Student %d shows steady progress with strengths in visual learning.", i+1),
			MainStrengths:      []string{interests[i%len(interests)], "persistence"},
			MainChallenges:     []string{"attention_stamina", "writing"},
			PrimaryGoals:       []string{"improve_reading_fluency", "master_multiplication"},
			Scores: &learn.WorkingScores{
				OverallAcademicReadiness: int32(40 + rand.Intn(60)),
				ReadingReadiness:  int32(30 + rand.Intn(70)),
				MathReadiness:     int32(35 + rand.Intn(65)),
				WritingFineMotor:  int32(25 + rand.Intn(50)),
				SpeechLanguage:    int32(50 + rand.Intn(50)),
				AttentionTaskStamina:  int32(20 + rand.Intn(60)),
				SocialMotivation:   int32(40 + rand.Intn(60)),
				IndependenceDailyLiving:      int32(30 + rand.Intn(70)),
			},
			LearningStyle: &learn.LearningStyle{
				PreferredModes:           []string{modes[i%len(modes)], modes[(i+1)%len(modes)]},
				BestSessionLengthMinutes: int32(10 + rand.Intn(20)),
				BestActivityLengthMinutes: int32(3 + rand.Intn(10)),
				BreakFrequencyMinutes:    int32(8 + rand.Intn(15)),
				BestTimeOfDay:            times[i%len(times)],
			},
			Attention: &learn.AttentionRegulationProfile{
				FocusPreferredActivityMinutes: int32(10 + rand.Intn(20)),
				FocusAcademicTaskMinutes:      int32(5 + rand.Intn(15)),
				LosingFocusSigns:             []string{focusSigns[i%len(focusSigns)]},
				RegulationSupports:              []string{supports[i%len(supports)]},
			},
			Motivation: &learn.MotivationProfile{
				HighInterestActivities:         []string{interests[i%len(interests)], interests[(i+2)%len(interests)]},
				RewardPreferences:  []string{rewards[i%len(rewards)]},
			},
			Literacy: &learn.LiteracyProfile{
				ReadingLevel:      readingLevels[i%len(readingLevels)],
				Comprehension:     "developing",
				ReadingFluencyWpm: int32(30 + rand.Intn(90)),
			},
			Math: &learn.MathProfile{
				Level:          mathLevels[i%len(mathLevels)],
				Addition:       "developing",
				Multiplication: "emerging",
			},
			AdaptiveRules: &learn.AdaptiveLearningRules{
				DefaultSessionLengthMinutes: int32(15),
				MaximumSessionLengthMinutes: int32(30),
				MaximumActivityLengthMinutes: int32(10),
				BreakFrequencyMinutes:       int32(10),
				MaxConsecutiveErrors:        3,
				MaxConsecutiveCorrect:       5,
			},
			AiTutor: &learn.AITutorSettings{
				Personality:    []string{personalities[i%len(personalities)]},
				ShouldDo:       []string{"use_simple_words", "celebrate_effort"},
				ShouldAvoid:    []string{"time_pressure", "negative_feedback"},
			},
		}
		list = append(list, profile)
	}
	return list
}
