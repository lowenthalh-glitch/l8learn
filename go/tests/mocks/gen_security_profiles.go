/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * Profiles, evaluations, and schedules for the 5 security test students.
 * Each profile is realistic with different challenges and strengths.
 */
package mocks

import (
	"fmt"
	"github.com/saichler/l8learn/go/types/learn"
	"time"
)

func generateSecurityProfiles() []*learn.StudentProfile {
	now := time.Now().Unix()
	return []*learn.StudentProfile{
		diegoProfile(now),
		marcoProfile(now),
		emmaProfile(now),
		zaraProfile(now),
		laylaProfile(now),
	}
}

// Diego Garcia — Grade 3, ADHD, strong math, struggles with reading/writing
func diegoProfile(now int64) *learn.StudentProfile {
	return &learn.StudentProfile{
		ProfileId:         "PROF-SEC-001",
		StudentId:         "STU-SEC-001",
		PrimaryGuardianId: "GRD-SEC-001",
		CreatedDate:       now,
		LastUpdated:       now,
		ShortSummary:      "Diego is an energetic, curious boy with ADHD who excels in math and logical thinking but struggles with sustained reading, written expression, and impulse control. He learns best through hands-on, fast-paced activities with clear structure and immediate feedback.",
		MainStrengths:     []string{"strong logical/math reasoning", "curious and inquisitive", "good spatial awareness", "enjoys building and construction", "responds well to challenge-based learning"},
		MainChallenges:    []string{"ADHD - inattention and impulsivity", "reading comprehension below grade level", "written expression difficulty", "difficulty with multi-step written directions"},
		MainLearningBarriers: []string{"attention wanders during sustained reading tasks", "rushes through written work with errors", "difficulty organizing thoughts for writing", "frustration when tasks feel too slow"},
		Scores: &learn.WorkingScores{
			OverallAcademicReadiness: 6, ReadingReadiness: 4, MathReadiness: 8,
			WritingFineMotor: 4, SpeechLanguage: 7, AttentionTaskStamina: 3,
			GrossMotor: 8, SocialMotivation: 6, IndependenceDailyLiving: 6, ConfidenceWithLearning: 5,
		},
		LearningStyle: &learn.LearningStyle{
			PreferredModes: []string{"kinesthetic", "visual", "hands-on"},
			BestSessionLengthMinutes: 25, BestActivityLengthMinutes: 8,
			MaxSeatedWorkMinutes: 10, BreakFrequencyMinutes: 8, BestTimeOfDay: "morning",
			WorksBestWith:  []string{"timers", "challenges", "building activities", "movement breaks", "competitive games"},
			WorksPoorlyWith: []string{"long reading passages", "extended writing", "waiting", "repetitive drills"},
		},
		Attention: &learn.AttentionRegulationProfile{
			MaxBookSittingTime: "8 to 12 minutes", StructuredTaskStamina: "10 minutes with timer",
			NeedsFrequentBreaks: true, ImpulsivityPresent: true, DistractibilityPresent: true,
			FocusAcademicTaskMinutes: 10, FocusPreferredActivityMinutes: 30,
			LosingFocusSigns: []string{"fidgets", "blurts out", "leaves seat", "starts side conversations"},
			RegulationSupports: []string{"visual timer", "movement breaks", "fidget tool", "clear endpoint", "challenge framing"},
		},
		Literacy: &learn.LiteracyProfile{
			CurrentLevel: "beginning Grade 2 reading level (below Grade 3)", ReadingLevel: "early Grade 2",
			Comprehension: "literal understanding good, inferencing weak",
			PrimaryNeeds: []string{"reading fluency", "inferencing", "multi-paragraph comprehension", "written expression"},
		},
		Math: &learn.MathProfile{
			Level: "above grade level — strong Grade 3/early Grade 4",
			PrioritySkills: []string{"multi-digit multiplication", "fractions introduction", "word problems"},
			RecommendedActivities: []string{"math puzzles", "building challenges with measurement", "competitive math games"},
		},
		Motivation: &learn.MotivationProfile{
			HighInterestActivities: []string{"Legos", "Minecraft", "math challenges", "soccer", "building robots", "science experiments"},
			RewardPreferences: []string{"extra building time", "choice of activity", "challenge badges", "timer to beat"},
		},
	}
}

// Marco Garcia — Grade 1, speech/language delay (Diego's younger brother)
func marcoProfile(now int64) *learn.StudentProfile {
	return &learn.StudentProfile{
		ProfileId:         "PROF-SEC-002",
		StudentId:         "STU-SEC-002",
		PrimaryGuardianId: "GRD-SEC-001",
		CreatedDate:       now,
		LastUpdated:       now,
		ShortSummary:      "Marco is a shy, sweet boy with a moderate speech-language delay who is making progress with therapy. He has strong visual-spatial skills and enjoys puzzles and drawing. He needs support with articulation, sentence structure, following directions, and early literacy.",
		MainStrengths:     []string{"strong visual-spatial skills", "enjoys puzzles and drawing", "gentle and cooperative", "good fine motor for age", "responds to encouragement"},
		MainChallenges:    []string{"moderate speech-language delay", "articulation difficulty", "shy in group settings", "slow processing speed"},
		MainLearningBarriers: []string{"difficulty being understood by unfamiliar listeners", "avoids speaking in groups", "needs extra processing time", "frustration when communication fails"},
		Scores: &learn.WorkingScores{
			OverallAcademicReadiness: 4, ReadingReadiness: 3, MathReadiness: 5,
			WritingFineMotor: 6, SpeechLanguage: 3, AttentionTaskStamina: 5,
			GrossMotor: 5, SocialMotivation: 4, IndependenceDailyLiving: 5, ConfidenceWithLearning: 4,
		},
		LearningStyle: &learn.LearningStyle{
			PreferredModes: []string{"visual", "hands-on"},
			BestSessionLengthMinutes: 25, BestActivityLengthMinutes: 7,
			MaxSeatedWorkMinutes: 8, BreakFrequencyMinutes: 7, BestTimeOfDay: "morning",
			WorksBestWith:  []string{"visual models", "drawing", "puzzles", "quiet environment", "one-on-one support"},
			WorksPoorlyWith: []string{"group discussions", "oral presentations", "noisy environments", "time pressure"},
		},
		Attention: &learn.AttentionRegulationProfile{
			MaxBookSittingTime: "10 to 15 minutes", StructuredTaskStamina: "8 minutes",
			FocusAcademicTaskMinutes: 8, FocusPreferredActivityMinutes: 20,
		},
		Speech: &learn.SpeechLanguageProfile{
			TherapyNeed: "speech-language therapy ongoing", ReceivesSpeechTherapy: true,
			Clarity: "intelligible to familiar listeners, reduced with unfamiliar",
			Articulation: &learn.ArticulationProfile{
				NotedSoundChallenges: []string{"R sound", "L blends", "TH sound"},
				FunctionalImpact:     []string{"reduced intelligibility with strangers", "avoids speaking in class"},
			},
			ExpressiveLanguage: &learn.StrengthsAndNeeds{
				Strengths: []string{"uses gestures effectively", "can describe pictures with support"},
				Needs:     []string{"sentence structure", "vocabulary expansion", "narrative skills"},
			},
		},
		Motivation: &learn.MotivationProfile{
			HighInterestActivities: []string{"puzzles", "drawing", "coloring", "Legos", "animals", "dinosaurs"},
			RewardPreferences: []string{"stickers", "drawing time", "quiet praise", "showing work to parent"},
		},
	}
}

// Emma Wilson — Grade 2, fine motor difficulty, strong verbal skills
func emmaProfile(now int64) *learn.StudentProfile {
	return &learn.StudentProfile{
		ProfileId:         "PROF-SEC-003",
		StudentId:         "STU-SEC-003",
		PrimaryGuardianId: "GRD-SEC-002",
		CreatedDate:       now,
		LastUpdated:       now,
		ShortSummary:      "Emma is a verbally advanced, creative girl with fine motor and graphomotor delays. She has strong oral language and reading comprehension but struggles with handwriting, cutting, and written output. She benefits from assistive technology and alternative output methods.",
		MainStrengths:     []string{"advanced verbal skills", "strong reading comprehension", "creative storytelling", "good social skills", "loves music and drama"},
		MainChallenges:    []string{"fine motor and graphomotor delays", "handwriting fatigue", "slow written output", "avoids writing tasks"},
		Scores: &learn.WorkingScores{
			OverallAcademicReadiness: 7, ReadingReadiness: 8, MathReadiness: 6,
			WritingFineMotor: 3, SpeechLanguage: 9, AttentionTaskStamina: 6,
			GrossMotor: 5, SocialMotivation: 8, IndependenceDailyLiving: 6, ConfidenceWithLearning: 6,
		},
		LearningStyle: &learn.LearningStyle{
			PreferredModes: []string{"auditory", "verbal", "visual"},
			BestSessionLengthMinutes: 30, BestActivityLengthMinutes: 10,
			MaxSeatedWorkMinutes: 12, BreakFrequencyMinutes: 10, BestTimeOfDay: "morning",
			WorksBestWith:  []string{"oral discussions", "storytelling", "drama", "music", "dictation instead of writing"},
			WorksPoorlyWith: []string{"long handwriting tasks", "timed writing", "small print worksheets"},
		},
		FineMotor: &learn.FineMotorOTProfile{
			TherapyNeed: "occupational therapy recommended", ReceivesOccupationalTherapy: true,
			HandDominance: "right", PencilGrip: "modified tripod, fatigues quickly",
			ObservedNeeds: []string{"handwriting endurance", "letter sizing", "spacing", "cutting curves"},
			HelpfulWritingSupports: []string{"pencil grip", "raised-line paper", "short writing tasks", "oral alternatives"},
		},
		Motivation: &learn.MotivationProfile{
			HighInterestActivities: []string{"reading", "storytelling", "drama", "singing", "art", "playing teacher"},
			RewardPreferences: []string{"read-aloud time", "drama activity", "verbal praise", "choice of story"},
		},
	}
}

// Zara Khan — Grade 4, gifted with anxiety and social challenges
func zaraProfile(now int64) *learn.StudentProfile {
	return &learn.StudentProfile{
		ProfileId:         "PROF-SEC-004",
		StudentId:         "STU-SEC-004",
		PrimaryGuardianId: "GRD-SEC-003",
		CreatedDate:       now,
		LastUpdated:       now,
		ShortSummary:      "Zara is academically gifted (reading at Grade 6 level, math at Grade 5) but experiences anxiety around performance, perfectionism, and social interactions. She benefits from emotional support, flexible pacing, and social skills coaching.",
		MainStrengths:     []string{"advanced academic skills", "strong reading and writing", "independent learner", "deep focus on interests", "creative problem solver"},
		MainChallenges:    []string{"performance anxiety", "perfectionism", "difficulty with peer interaction", "avoids new social situations", "meltdowns when frustrated"},
		Scores: &learn.WorkingScores{
			OverallAcademicReadiness: 9, ReadingReadiness: 9, MathReadiness: 8,
			WritingFineMotor: 7, SpeechLanguage: 8, AttentionTaskStamina: 7,
			GrossMotor: 5, SocialMotivation: 3, IndependenceDailyLiving: 7, ConfidenceWithLearning: 4,
		},
		LearningStyle: &learn.LearningStyle{
			PreferredModes: []string{"reading", "visual", "independent"},
			BestSessionLengthMinutes: 40, BestActivityLengthMinutes: 15,
			MaxSeatedWorkMinutes: 20, BreakFrequencyMinutes: 15, BestTimeOfDay: "morning",
			WorksBestWith:  []string{"independent work", "deep reading", "creative projects", "predictable routine", "quiet environment"},
			WorksPoorlyWith: []string{"group work", "oral presentations", "competitive activities", "unexpected changes", "being wrong publicly"},
		},
		SocialEmotional: &learn.SocialEmotionalProfile{
			Confidence: "high academically, low socially",
			Strengths:  []string{"empathetic", "kind to younger children", "strong sense of justice"},
			Needs:      []string{"anxiety management", "social skills practice", "flexibility with mistakes", "peer interaction coaching"},
			FrustrationTriggers: []string{"making mistakes", "being corrected publicly", "group projects", "time pressure"},
			CalmingStrategies:   []string{"deep breathing", "reading alone", "drawing", "talking to parent"},
		},
		Motivation: &learn.MotivationProfile{
			HighInterestActivities: []string{"reading chapter books", "writing stories", "science", "coding", "chess", "astronomy"},
			RewardPreferences: []string{"new book", "extra reading time", "science experiment", "quiet recognition"},
		},
	}
}

// Layla Khan — Kindergarten, developing normally, benefits from structured play
func laylaProfile(now int64) *learn.StudentProfile {
	return &learn.StudentProfile{
		ProfileId:         "PROF-SEC-005",
		StudentId:         "STU-SEC-005",
		PrimaryGuardianId: "GRD-SEC-003",
		CreatedDate:       now,
		LastUpdated:       now,
		ShortSummary:      "Layla is a cheerful, active 5-year-old developing on track for Kindergarten. She enjoys social play, songs, and movement activities. She is beginning to recognize letters and numbers and benefits from structured, playful learning with lots of encouragement.",
		MainStrengths:     []string{"cheerful and social", "loves songs and dance", "good gross motor skills", "eager to learn", "follows simple routines well"},
		MainChallenges:    []string{"still developing letter recognition", "short attention span (age-appropriate)", "needs support with pencil grip", "emerging number sense"},
		Scores: &learn.WorkingScores{
			OverallAcademicReadiness: 5, ReadingReadiness: 4, MathReadiness: 4,
			WritingFineMotor: 4, SpeechLanguage: 6, AttentionTaskStamina: 4,
			GrossMotor: 7, SocialMotivation: 8, IndependenceDailyLiving: 5, ConfidenceWithLearning: 7,
		},
		LearningStyle: &learn.LearningStyle{
			PreferredModes: []string{"kinesthetic", "musical", "social"},
			BestSessionLengthMinutes: 20, BestActivityLengthMinutes: 5,
			MaxSeatedWorkMinutes: 5, BreakFrequencyMinutes: 5, BestTimeOfDay: "morning",
			WorksBestWith:  []string{"songs", "movement games", "social play", "colorful materials", "short tasks"},
			WorksPoorlyWith: []string{"long seated work", "worksheets", "working alone", "abstract concepts"},
		},
		Attention: &learn.AttentionRegulationProfile{
			MaxBookSittingTime: "5 to 8 minutes", StructuredTaskStamina: "5 minutes",
			FocusAcademicTaskMinutes: 5, FocusPreferredActivityMinutes: 15,
		},
		Motivation: &learn.MotivationProfile{
			HighInterestActivities: []string{"singing", "dancing", "playing with dolls", "playing house", "drawing", "outdoor play", "bubbles"},
			RewardPreferences: []string{"stickers", "songs", "dance break", "playing with sister"},
		},
	}
}

func generateSecurityEvals() []*learn.EvalImport {
	now := time.Now().Unix()
	return []*learn.EvalImport{
		{ImportId: "EVAL-SEC-001", StudentId: "STU-SEC-001", PrimaryGuardianId: "GRD-SEC-001",
			DocumentType: learn.EvalDocumentType_EVAL_DOCUMENT_TYPE_PSYCHOLOGICAL, ProfessionalName: "Dr. Robert Chen, PhD",
			EvaluationDate: now - 90*24*3600, ProcessingStatus: learn.EvalProcessingStatus_EVAL_PROCESSING_COMPLETE},
		{ImportId: "EVAL-SEC-002", StudentId: "STU-SEC-001", PrimaryGuardianId: "GRD-SEC-001",
			DocumentType: learn.EvalDocumentType_EVAL_DOCUMENT_TYPE_READING_SPECIALIST, ProfessionalName: "Ms. Lisa Park, Reading Specialist",
			EvaluationDate: now - 60*24*3600, ProcessingStatus: learn.EvalProcessingStatus_EVAL_PROCESSING_COMPLETE},
		{ImportId: "EVAL-SEC-003", StudentId: "STU-SEC-002", PrimaryGuardianId: "GRD-SEC-001",
			DocumentType: learn.EvalDocumentType_EVAL_DOCUMENT_TYPE_SPEECH, ProfessionalName: "Dr. Anna Reyes, SLP",
			EvaluationDate: now - 45*24*3600, ProcessingStatus: learn.EvalProcessingStatus_EVAL_PROCESSING_COMPLETE},
		{ImportId: "EVAL-SEC-004", StudentId: "STU-SEC-003", PrimaryGuardianId: "GRD-SEC-002",
			DocumentType: learn.EvalDocumentType_EVAL_DOCUMENT_TYPE_OCCUPATIONAL_THERAPY, ProfessionalName: "Dr. Mike Torres, OTR/L",
			EvaluationDate: now - 30*24*3600, ProcessingStatus: learn.EvalProcessingStatus_EVAL_PROCESSING_COMPLETE},
		{ImportId: "EVAL-SEC-005", StudentId: "STU-SEC-004", PrimaryGuardianId: "GRD-SEC-003",
			DocumentType: learn.EvalDocumentType_EVAL_DOCUMENT_TYPE_PSYCHOLOGICAL, ProfessionalName: "Dr. Priya Sharma, PsyD",
			EvaluationDate: now - 120*24*3600, ProcessingStatus: learn.EvalProcessingStatus_EVAL_PROCESSING_COMPLETE},
		{ImportId: "EVAL-SEC-006", StudentId: "STU-SEC-005", PrimaryGuardianId: "GRD-SEC-003",
			DocumentType: learn.EvalDocumentType_EVAL_DOCUMENT_TYPE_DEVELOPMENTAL, ProfessionalName: "Dr. Noor Hassan, Developmental Pediatrician",
			EvaluationDate: now - 15*24*3600, ProcessingStatus: learn.EvalProcessingStatus_EVAL_PROCESSING_COMPLETE},
	}
}

func generateSecuritySchedules() []*learn.DailySchedule {
	students := []struct{ id, guardianId, name string }{
		{"STU-SEC-001", "GRD-SEC-001", "Diego"},
		{"STU-SEC-002", "GRD-SEC-001", "Marco"},
		{"STU-SEC-003", "GRD-SEC-002", "Emma"},
		{"STU-SEC-004", "GRD-SEC-003", "Zara"},
		{"STU-SEC-005", "GRD-SEC-003", "Layla"},
	}

	var schedules []*learn.DailySchedule
	for _, s := range students {
		sched := &learn.DailySchedule{
			ScheduleId:       fmt.Sprintf("SCHED-%s", s.id),
			FamilyId:         fmt.Sprintf("FAM-%s", s.guardianId),
			ScheduleDate:     time.Now().Unix(),
			AvailableHours:   4,
			ParentEnergy:     "medium",
			Weather:          "sunny",
			LessonsTotal:     4,
			LessonsGenerated: 4,
			CustomFields:     map[string]string{"studentId": s.id},
			Blocks: []*learn.ScheduleBlock{
				{BlockId: s.id + "-mon-1", StartMinute: 540, DurationMinutes: 5, ActivityType: "movement_warmup", Description: "Morning movement warm-up for " + s.name},
				{BlockId: s.id + "-mon-2", StartMinute: 545, DurationMinutes: 7, ActivityType: "academic", Description: "Literacy activity for " + s.name},
				{BlockId: s.id + "-mon-3", StartMinute: 552, DurationMinutes: 3, ActivityType: "break", Description: "Movement break"},
				{BlockId: s.id + "-mon-4", StartMinute: 555, DurationMinutes: 7, ActivityType: "therapy", Description: "Therapy-aligned practice for " + s.name},
				{BlockId: s.id + "-mon-5", StartMinute: 562, DurationMinutes: 5, ActivityType: "creative", Description: "Creative activity"},
				{BlockId: s.id + "-mon-6", StartMinute: 567, DurationMinutes: 3, ActivityType: "cleanup", Description: "Cleanup and celebration"},
			},
		}
		schedules = append(schedules, sched)
	}
	return schedules
}
