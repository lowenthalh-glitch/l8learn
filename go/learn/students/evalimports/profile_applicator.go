/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * Maps accepted EvalImport findings to StudentProfile fields.
 * Extension functions (speech, motor, sensory, etc.) in profile_applicator_ext.go.
 */
package evalimports

import (
	"strconv"
	"strings"

	"github.com/saichler/l8learn/go/types/learn"
)

// ApplyFindingsToProfile applies accepted/edited findings to the StudentProfile.
func ApplyFindingsToProfile(eval *learn.EvalImport, profile *learn.StudentProfile) error {
	for _, finding := range eval.Findings {
		if finding.Status == learn.EvalFindingStatus_EVAL_FINDING_STATUS_ACCEPTED ||
			finding.Status == learn.EvalFindingStatus_EVAL_FINDING_STATUS_EDITED {
			value := finding.NewValue
			if finding.Status == learn.EvalFindingStatus_EVAL_FINDING_STATUS_EDITED && finding.EditedValue != "" {
				value = finding.EditedValue
			}
			applyFinding(profile, finding.ProfileSection, finding.ProfileField, value)
		}
	}
	return nil
}

func applyFinding(profile *learn.StudentProfile, section, field, value string) {
	switch section {
	case "scores":
		applyScores(profile, field, value)
	case "strengths":
		applyStrengths(profile, field, value)
	case "challenges":
		applyChallenges(profile, field, value)
	case "learningStyle":
		applyLearningStyle(profile, field, value)
	case "attention":
		applyAttention(profile, field, value)
	case "motivation":
		applyMotivation(profile, field, value)
	case "literacy":
		applyLiteracy(profile, field, value)
	case "math":
		applyMath(profile, field, value)
	case "speech":
		applySpeech(profile, field, value)
	case "speech.articulation":
		applySpeechArticulation(profile, field, value)
	case "speech.expressiveLanguage":
		applySpeechExpressive(profile, field, value)
	case "speech.receptiveLanguage":
		applySpeechReceptive(profile, field, value)
	case "speech.languageHistory":
		applySpeechHistory(profile, field, value)
	case "fineMotor":
		applyFineMotor(profile, field, value)
	case "fineMotor.nameWritingStatus":
		applyNameWriting(profile, field, value)
	case "grossMotor":
		applyGrossMotor(profile, field, value)
	case "sensory":
		applySensory(profile, field, value)
	case "socialEmotional":
		applySocialEmotional(profile, field, value)
	case "behavior":
		applyBehavior(profile, field, value)
	case "dailyLiving":
		applyDailyLiving(profile, field, value)
	case "dailyLiving.toileting":
		applyToileting(profile, field, value)
	case "dailyLiving.hygiene":
		applyHygiene(profile, field, value)
	case "dailyLiving.dressing":
		applyDressing(profile, field, value)
	case "dailyLiving.feeding":
		applyFeeding(profile, field, value)
	case "health":
		applyHealth(profile, field, value)
	}
}

func toInt32(s string) int32 {
	v, _ := strconv.Atoi(s)
	return int32(v)
}

func toStringList(s string) []string {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	var result []string
	for _, p := range parts {
		t := strings.TrimSpace(p)
		if t != "" {
			result = append(result, t)
		}
	}
	return result
}

func ensureScores(p *learn.StudentProfile) *learn.WorkingScores {
	if p.Scores == nil {
		p.Scores = &learn.WorkingScores{}
	}
	return p.Scores
}

func applyScores(p *learn.StudentProfile, field, value string) {
	s := ensureScores(p)
	switch field {
	case "overallAcademicReadiness":
		s.OverallAcademicReadiness = toInt32(value)
	case "readingReadiness":
		s.ReadingReadiness = toInt32(value)
	case "mathReadiness":
		s.MathReadiness = toInt32(value)
	case "writingFineMotor":
		s.WritingFineMotor = toInt32(value)
	case "speechLanguage":
		s.SpeechLanguage = toInt32(value)
	case "attentionTaskStamina":
		s.AttentionTaskStamina = toInt32(value)
	case "grossMotor":
		s.GrossMotor = toInt32(value)
	case "socialMotivation":
		s.SocialMotivation = toInt32(value)
	case "independenceDailyLiving":
		s.IndependenceDailyLiving = toInt32(value)
	case "confidenceWithLearning":
		s.ConfidenceWithLearning = toInt32(value)
	}
}

func applyStrengths(p *learn.StudentProfile, field, value string) {
	if p.Strengths == nil {
		p.Strengths = &learn.CategorizedStrengths{}
	}
	switch field {
	case "socialEmotional":
		p.Strengths.SocialEmotional = toStringList(value)
	case "playAndMotivation":
		p.Strengths.PlayAndMotivation = toStringList(value)
	case "grossMotor":
		p.Strengths.GrossMotor = toStringList(value)
	case "academic":
		p.Strengths.Academic = toStringList(value)
	case "communication":
		p.Strengths.Communication = toStringList(value)
	}
}

func applyChallenges(p *learn.StudentProfile, field, value string) {
	if p.Challenges == nil {
		p.Challenges = &learn.CategorizedChallenges{}
	}
	switch field {
	case "speechLanguage":
		p.Challenges.SpeechLanguage = toStringList(value)
	case "attentionExecutiveFunction":
		p.Challenges.AttentionExecutiveFunction = toStringList(value)
	case "fineMotorGraphomotor":
		p.Challenges.FineMotorGraphomotor = toStringList(value)
	case "sensoryMotor":
		p.Challenges.SensoryMotor = toStringList(value)
	case "academicReadiness":
		p.Challenges.AcademicReadiness = toStringList(value)
	}
}

func applyLearningStyle(p *learn.StudentProfile, field, value string) {
	if p.LearningStyle == nil {
		p.LearningStyle = &learn.LearningStyle{}
	}
	ls := p.LearningStyle
	switch field {
	case "preferredModes":
		ls.PreferredModes = toStringList(value)
	case "bestSessionLengthMinutes":
		ls.BestSessionLengthMinutes = toInt32(value)
	case "bestActivityLengthMinutes":
		ls.BestActivityLengthMinutes = toInt32(value)
	case "maxSeatedWorkMinutes":
		ls.MaxSeatedWorkMinutes = toInt32(value)
	case "breakFrequencyMinutes":
		ls.BreakFrequencyMinutes = toInt32(value)
	case "bestTimeOfDay":
		ls.BestTimeOfDay = value
	case "bestLearningFormula":
		ls.BestLearningFormula = toStringList(value)
	case "worksBestWith":
		ls.WorksBestWith = toStringList(value)
	case "worksPoorlyWith":
		ls.WorksPoorlyWith = toStringList(value)
	}
}

func applyAttention(p *learn.StudentProfile, field, value string) {
	if p.Attention == nil {
		p.Attention = &learn.AttentionRegulationProfile{}
	}
	a := p.Attention
	switch field {
	case "maxBookSittingTime":
		a.MaxBookSittingTime = value
	case "structuredTaskStamina":
		a.StructuredTaskStamina = value
	case "needsFrequentBreaks":
		a.NeedsFrequentBreaks = value == "true"
	case "impulsivityPresent":
		a.ImpulsivityPresent = value == "true"
	case "distractibilityPresent":
		a.DistractibilityPresent = value == "true"
	case "focusPreferredActivityMinutes":
		a.FocusPreferredActivityMinutes = toInt32(value)
	case "focusAcademicTaskMinutes":
		a.FocusAcademicTaskMinutes = toInt32(value)
	case "losingFocusSigns":
		a.LosingFocusSigns = toStringList(value)
	case "regulationSupports":
		a.RegulationSupports = toStringList(value)
	}
}

func applyMotivation(p *learn.StudentProfile, field, value string) {
	if p.Motivation == nil {
		p.Motivation = &learn.MotivationProfile{}
	}
	switch field {
	case "highInterestActivities":
		p.Motivation.HighInterestActivities = toStringList(value)
	case "rewardPreferences":
		p.Motivation.RewardPreferences = toStringList(value)
	case "avoidAsReward":
		p.Motivation.AvoidAsReward = toStringList(value)
	case "avoidedActivities":
		p.Motivation.AvoidedActivities = toStringList(value)
	}
}

func applyLiteracy(p *learn.StudentProfile, field, value string) {
	if p.Literacy == nil {
		p.Literacy = &learn.LiteracyProfile{}
	}
	l := p.Literacy
	switch field {
	case "currentLevel":
		l.CurrentLevel = value
	case "readingLevel":
		l.ReadingLevel = value
	case "letterRecognition":
		l.LetterRecognition = value
	case "phonemicAwareness":
		l.PhonemicAwareness = value
	case "sightWords":
		l.SightWords = value
	case "comprehension":
		l.Comprehension = value
	case "bookStaminaMinutes":
		l.BookStaminaMinutes = toInt32(value)
	case "primaryNeeds":
		l.PrimaryNeeds = toStringList(value)
	case "prioritySequence":
		l.PrioritySequence = toStringList(value)
	case "recommendedApproach":
		l.RecommendedApproach = toStringList(value)
	case "avoid":
		l.Avoid = toStringList(value)
	}
}

func applyMath(p *learn.StudentProfile, field, value string) {
	if p.Math == nil {
		p.Math = &learn.MathProfile{}
	}
	m := p.Math
	switch field {
	case "level":
		m.Level = value
	case "recommendedMode":
		m.RecommendedMode = value
	case "counting":
		m.Counting = value
	case "numberRecognition":
		m.NumberRecognition = value
	case "oneToOneCorrespondence":
		m.OneToOneCorrespondence = value
	case "prioritySkills":
		m.PrioritySkills = toStringList(value)
	case "recommendedActivities":
		m.RecommendedActivities = toStringList(value)
	case "errorPatterns":
		m.ErrorPatterns = toStringList(value)
	case "avoid":
		m.Avoid = toStringList(value)
	}
}
