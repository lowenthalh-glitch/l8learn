/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * Profile applicator extension — speech, motor, sensory, social, daily living, health.
 */
package evalimports

import (
	"github.com/saichler/l8learn/go/types/learn"
)

func ensureSpeech(p *learn.StudentProfile) *learn.SpeechLanguageProfile {
	if p.Speech == nil {
		p.Speech = &learn.SpeechLanguageProfile{}
	}
	return p.Speech
}

func applySpeech(p *learn.StudentProfile, field, value string) {
	s := ensureSpeech(p)
	switch field {
	case "therapyNeed":
		s.TherapyNeed = value
	case "receivesSpeechTherapy":
		s.ReceivesSpeechTherapy = value == "true"
	case "clarity":
		s.Clarity = value
	case "speechSounds":
		s.SpeechSounds = toStringList(value)
	case "currentGoals":
		s.CurrentGoals = toStringList(value)
	case "helpfulPrompts":
		s.HelpfulPrompts = toStringList(value)
	}
}

func applySpeechArticulation(p *learn.StudentProfile, field, value string) {
	s := ensureSpeech(p)
	if s.Articulation == nil {
		s.Articulation = &learn.ArticulationProfile{}
	}
	switch field {
	case "notedSoundChallenges":
		s.Articulation.NotedSoundChallenges = toStringList(value)
	case "phonologicalProcessesObserved":
		s.Articulation.PhonologicalProcessesObserved = toStringList(value)
	case "functionalImpact":
		s.Articulation.FunctionalImpact = toStringList(value)
	}
}

func applySpeechExpressive(p *learn.StudentProfile, field, value string) {
	s := ensureSpeech(p)
	if s.ExpressiveLanguage == nil {
		s.ExpressiveLanguage = &learn.StrengthsAndNeeds{}
	}
	switch field {
	case "strengths":
		s.ExpressiveLanguage.Strengths = toStringList(value)
	case "needs":
		s.ExpressiveLanguage.Needs = toStringList(value)
	}
}

func applySpeechReceptive(p *learn.StudentProfile, field, value string) {
	s := ensureSpeech(p)
	if s.ReceptiveLanguage == nil {
		s.ReceptiveLanguage = &learn.ReceptiveLanguage{}
	}
	switch field {
	case "needs":
		s.ReceptiveLanguage.Needs = toStringList(value)
	case "bestSupports":
		s.ReceptiveLanguage.BestSupports = toStringList(value)
	}
}

func applySpeechHistory(p *learn.StudentProfile, field, value string) {
	s := ensureSpeech(p)
	if s.LanguageHistory == nil {
		s.LanguageHistory = &learn.LanguageHistory{}
	}
	switch field {
	case "lateLanguageOnset":
		s.LanguageHistory.LateLanguageOnset = value == "true"
	case "singleWordsApproximateAge":
		s.LanguageHistory.SingleWordsApproximateAge = value
	case "sentencesApproximateAge":
		s.LanguageHistory.SentencesApproximateAge = value
	}
}

func applyFineMotor(p *learn.StudentProfile, field, value string) {
	if p.FineMotor == nil {
		p.FineMotor = &learn.FineMotorOTProfile{}
	}
	fm := p.FineMotor
	switch field {
	case "therapyNeed":
		fm.TherapyNeed = value
	case "receivesOccupationalTherapy":
		fm.ReceivesOccupationalTherapy = value == "true"
	case "handDominance":
		fm.HandDominance = value
	case "pencilGrip":
		fm.PencilGrip = value
	case "tracing":
		fm.Tracing = value
	case "cutting":
		fm.Cutting = value
	case "observedNeeds":
		fm.ObservedNeeds = toStringList(value)
	case "helpfulWritingSupports":
		fm.HelpfulWritingSupports = toStringList(value)
	case "helpfulTools":
		fm.HelpfulTools = toStringList(value)
	case "recommendedActivities":
		fm.RecommendedActivities = toStringList(value)
	case "avoid":
		fm.Avoid = toStringList(value)
	}
}

func applyNameWriting(p *learn.StudentProfile, field, value string) {
	if p.FineMotor == nil {
		p.FineMotor = &learn.FineMotorOTProfile{}
	}
	if p.FineMotor.NameWritingStatus == nil {
		p.FineMotor.NameWritingStatus = &learn.NameWritingStatus{}
	}
	nw := p.FineMotor.NameWritingStatus
	switch field {
	case "targetName":
		nw.TargetName = value
	case "currentPattern":
		nw.CurrentPattern = value
	case "lettersNeedingSupport":
		nw.LettersNeedingSupport = toStringList(value)
	case "recommendedPractice":
		nw.RecommendedPractice = toStringList(value)
	}
}

func applyGrossMotor(p *learn.StudentProfile, field, value string) {
	if p.GrossMotor == nil {
		p.GrossMotor = &learn.GrossMotorProfile{}
	}
	gm := p.GrossMotor
	switch field {
	case "overallStatus":
		gm.OverallStatus = value
	case "energyLevel":
		gm.EnergyLevel = value
	case "strengths":
		gm.Strengths = toStringList(value)
	case "needs":
		gm.Needs = toStringList(value)
	case "learningUse":
		gm.LearningUse = value
	case "recommendedMovementBreaks":
		gm.RecommendedMovementBreaks = toStringList(value)
	case "movementBreakFrequencyMinutes":
		gm.MovementBreakFrequencyMinutes = toInt32(value)
	}
}

func applySensory(p *learn.StudentProfile, field, value string) {
	if p.Sensory == nil {
		p.Sensory = &learn.SensoryProfile{}
	}
	switch field {
	case "sensoryPattern":
		p.Sensory.SensoryPattern = value
	case "sensitivities":
		p.Sensory.Sensitivities = toStringList(value)
	case "registrationNeeds":
		p.Sensory.RegistrationNeeds = toStringList(value)
	case "functionalImpact":
		p.Sensory.FunctionalImpact = toStringList(value)
	case "helpfulSupports":
		p.Sensory.HelpfulSupports = toStringList(value)
	case "avoid":
		p.Sensory.Avoid = toStringList(value)
	}
}

func applySocialEmotional(p *learn.StudentProfile, field, value string) {
	if p.SocialEmotional == nil {
		p.SocialEmotional = &learn.SocialEmotionalProfile{}
	}
	se := p.SocialEmotional
	switch field {
	case "confidence":
		se.Confidence = value
	case "peerInteraction":
		se.PeerInteraction = value
	case "turnTaking":
		se.TurnTaking = value
	case "emotionNaming":
		se.EmotionNaming = value
	case "strengths":
		se.Strengths = toStringList(value)
	case "needs":
		se.Needs = toStringList(value)
	case "frustrationTriggers":
		se.FrustrationTriggers = toStringList(value)
	case "calmingStrategies":
		se.CalmingStrategies = toStringList(value)
	case "recommendedEmotionalSupports":
		se.RecommendedEmotionalSupports = toStringList(value)
	}
}

func applyBehavior(p *learn.StudentProfile, field, value string) {
	if p.Behavior == nil {
		p.Behavior = &learn.BehaviorProfile{}
	}
	switch field {
	case "avoidanceBehaviors":
		p.Behavior.AvoidanceBehaviors = toStringList(value)
	case "redirectStrategies":
		p.Behavior.RedirectStrategies = toStringList(value)
	case "successfulSupports":
		p.Behavior.SuccessfulSupports = toStringList(value)
	}
}

func ensureDailyLiving(p *learn.StudentProfile) *learn.DailyLivingProfile {
	if p.DailyLiving == nil {
		p.DailyLiving = &learn.DailyLivingProfile{}
	}
	return p.DailyLiving
}

func applyDailyLiving(p *learn.StudentProfile, field, value string) {
	dl := ensureDailyLiving(p)
	switch field {
	case "homePracticeIdeas":
		dl.HomePracticeIdeas = toStringList(value)
	}
}

func applyToileting(p *learn.StudentProfile, field, value string) {
	dl := ensureDailyLiving(p)
	if dl.Toileting == nil {
		dl.Toileting = &learn.ToiletingProfile{}
	}
	switch field {
	case "urination":
		dl.Toileting.Urination = value
	case "bowelHygiene":
		dl.Toileting.BowelHygiene = value
	}
}

func applyHygiene(p *learn.StudentProfile, field, value string) {
	dl := ensureDailyLiving(p)
	if dl.Hygiene == nil {
		dl.Hygiene = &learn.HygieneProfile{}
	}
	switch field {
	case "handWashing":
		dl.Hygiene.HandWashing = value
	case "toothBrushing":
		dl.Hygiene.ToothBrushing = value
	case "bathing":
		dl.Hygiene.Bathing = value
	}
}

func applyDressing(p *learn.StudentProfile, field, value string) {
	dl := ensureDailyLiving(p)
	if dl.Dressing == nil {
		dl.Dressing = &learn.DressingProfile{}
	}
	d := dl.Dressing
	switch field {
	case "lowerBody":
		d.LowerBody = value
	case "upperBody":
		d.UpperBody = value
	case "socks":
		d.Socks = value
	case "shoes":
		d.Shoes = value
	case "zippers":
		d.Zippers = value
	case "buttons":
		d.Buttons = value
	case "beltBuckle":
		d.BeltBuckle = value
	}
}

func applyFeeding(p *learn.StudentProfile, field, value string) {
	dl := ensureDailyLiving(p)
	if dl.Feeding == nil {
		dl.Feeding = &learn.FeedingProfile{}
	}
	f := dl.Feeding
	switch field {
	case "spoon":
		f.Spoon = value
	case "fork":
		f.Fork = value
	case "knife":
		f.Knife = value
	case "preference":
		f.Preference = value
	case "foodSelectivity":
		f.FoodSelectivity = value
	}
}

func applyHealth(p *learn.StudentProfile, field, value string) {
	if p.Health == nil {
		p.Health = &learn.HealthSafety{}
	}
	switch field {
	case "medicalConditions":
		p.Health.MedicalConditions = toStringList(value)
	case "visionConcerns":
		p.Health.VisionConcerns = value
	case "hearingConcerns":
		p.Health.HearingConcerns = value
	case "safetyConcerns":
		p.Health.SafetyConcerns = toStringList(value)
	}
}
