/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package engine

import (
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
	"time"
)

// AdaptiveEngine orchestrates the full pipeline:
// interaction scored → mastery updated → rules evaluated → AI fallback → path updated
type AdaptiveEngine struct {
	vnic ifs.IVNic
}

func NewAdaptiveEngine(vnic ifs.IVNic) *AdaptiveEngine {
	return &AdaptiveEngine{vnic: vnic}
}

// OnActivityCompleted is the main entry point, called when a student finishes an activity.
// It runs the full adaptive pipeline and updates the learning path.
func (e *AdaptiveEngine) OnActivityCompleted(
	session *learn.LearningSession,
	path *learn.LearningPath,
	student *learn.Student,
) error {
	// Step 1: Score the latest interactions
	interactions := session.Interactions
	if len(interactions) == 0 {
		return nil
	}

	// Step 2: Update SkillMastery for each skill touched
	skillMasteries, err := e.updateMastery(interactions)
	if err != nil {
		return err
	}

	// Step 3: Load active AdaptationRules
	rules, err := e.loadActiveRules(path.Subject)
	if err != nil {
		return err
	}

	// Step 4: Find the current skill's mastery
	currentMastery := findMastery(path.CurrentSkillId, skillMasteries)

	// Step 5: Evaluate rules (fast, deterministic)
	ruleResult := EvaluateRules(rules, interactions, currentMastery)

	// Step 6: Apply rule result or invoke AI
	if ruleResult != nil && ruleResult.Matched {
		err = e.applyRuleStrategy(path, ruleResult)
	} else {
		err = e.invokeAI(path, student, skillMasteries, interactions, ruleResult)
	}
	if err != nil {
		return err
	}

	// Step 7: Update path metadata
	path.ActivitiesCompleted++
	path.LastActive = time.Now().Unix()
	path.TotalTimeSeconds += totalTimeSpent(interactions)

	return nil
}

// updateMastery recalculates mastery for each skill that appeared in interactions
func (e *AdaptiveEngine) updateMastery(interactions []*learn.Interaction) ([]*learn.SkillMastery, error) {
	// Group interactions by skill
	bySkill := make(map[string][]*learn.Interaction)
	for _, i := range interactions {
		if i.SkillId != "" {
			bySkill[i.SkillId] = append(bySkill[i.SkillId], i)
		}
	}

	var updated []*learn.SkillMastery
	for skillId, skillInteractions := range bySkill {
		mastery, err := e.loadOrCreateMastery(skillId, interactions[0])
		if err != nil {
			return nil, err
		}

		// Update rolling accuracy
		correct := 0
		for _, i := range skillInteractions {
			mastery.AttemptsCount++
			if i.Result == learn.InteractionResult_INTERACTION_RESULT_CORRECT {
				correct++
				mastery.CorrectCount++
			}
		}
		mastery.CurrentAccuracy = float64(correct) / float64(len(skillInteractions))
		mastery.LastAttempted = time.Now().Unix()

		// Update mastery level based on accuracy
		mastery.Level = calculateMasteryLevel(mastery.CurrentAccuracy)
		if mastery.Level >= learn.MasteryLevel_MASTERY_LEVEL_MASTERED && mastery.MasteredDate == 0 {
			mastery.MasteredDate = time.Now().Unix()
		}

		// TODO: PATCH mastery record via service
		updated = append(updated, mastery)
	}
	return updated, nil
}

func calculateMasteryLevel(accuracy float64) learn.MasteryLevel {
	switch {
	case accuracy >= 0.95:
		return learn.MasteryLevel_MASTERY_LEVEL_EXEMPLARY
	case accuracy >= 0.80:
		return learn.MasteryLevel_MASTERY_LEVEL_MASTERED
	case accuracy >= 0.60:
		return learn.MasteryLevel_MASTERY_LEVEL_PROFICIENT
	case accuracy >= 0.40:
		return learn.MasteryLevel_MASTERY_LEVEL_DEVELOPING
	case accuracy > 0:
		return learn.MasteryLevel_MASTERY_LEVEL_EMERGING
	default:
		return learn.MasteryLevel_MASTERY_LEVEL_NOT_STARTED
	}
}

// applyRuleStrategy maps a rule's strategy to a concrete path change
func (e *AdaptiveEngine) applyRuleStrategy(path *learn.LearningPath, result *RuleResult) error {
	logEntry := &learn.AdaptationLog{
		Timestamp:      time.Now().Unix(),
		Trigger:        result.Rule.Trigger,
		Strategy:       result.Strategy,
		FromActivityId: path.CurrentActivityId,
		Reasoning:      "Rule: " + result.Rule.Name,
	}

	switch result.Strategy {
	case learn.AdaptStrategy_ADAPT_STRATEGY_REPEAT:
		// Stay on same skill, lower difficulty
		if path.CurrentDifficulty > learn.DifficultyLevel_DIFFICULTY_LEVEL_INTRO {
			path.CurrentDifficulty--
		}
	case learn.AdaptStrategy_ADAPT_STRATEGY_SCAFFOLD:
		// Break into sub-skills — find prerequisite skills
		// TODO: Look up skill.PrerequisiteSkillIds and target those
	case learn.AdaptStrategy_ADAPT_STRATEGY_ALTERNATE:
		// Same skill, different activity type
		// TODO: Find activity with same skill but different ActivityType
	case learn.AdaptStrategy_ADAPT_STRATEGY_REVIEW:
		// Go back to prerequisite skill
		// TODO: Look up skill prerequisites, pick weakest
	case learn.AdaptStrategy_ADAPT_STRATEGY_ADVANCE:
		// Move to next skill in the graph
		// TODO: Find next skill where prerequisites are met
	case learn.AdaptStrategy_ADAPT_STRATEGY_ENRICH:
		// Harder problems on same skill
		if path.CurrentDifficulty < learn.DifficultyLevel_DIFFICULTY_LEVEL_CHALLENGE {
			path.CurrentDifficulty++
		}
	case learn.AdaptStrategy_ADAPT_STRATEGY_BREAK:
		// Switch to engagement activity (game)
		// TODO: Find a GAME type activity at easy difficulty
	}

	path.AdaptationLog = append(path.AdaptationLog, logEntry)
	return nil
}

// invokeAI calls the LLM when rules don't produce a clear action
func (e *AdaptiveEngine) invokeAI(
	path *learn.LearningPath,
	student *learn.Student,
	masteries []*learn.SkillMastery,
	interactions []*learn.Interaction,
	ruleResult *RuleResult,
) error {
	// Build the full AI context
	req := &AIPathRequest{
		Student:      student,
		Path:         path,
		Mastery:      masteries,
		Interactions: interactions,
		RuleResult:   ruleResult,
	}

	// Load additional context
	req.Skills = e.loadSkills(path.Subject)
	req.Activities = e.loadActivities(path.Subject, path.TargetGrade)
	req.Engagement = e.loadEngagement(student.StudentId)
	req.Growth = e.loadGrowth(student.StudentId, path.Subject)
	req.PeerContext = e.loadPeerContext(student.ClassroomId)

	// Build prompt
	systemPrompt, userMessage := BuildPrompt(req)

	// Call LLM via l8agent pattern
	// TODO: Use l8agent orchestration (same pattern as l8agent/go/services/chat/orchestrate.go)
	_ = systemPrompt
	_ = userMessage

	// For now, placeholder: parse response and update path
	// response, err := llmClient.Call(systemPrompt, userMessage)
	// aiResult, err := ParseAIResponse(response)
	// Convert aiResult.NextActivities into path.UpcomingQueue

	return nil
}

// Helper loaders — these will call services via VNic
func (e *AdaptiveEngine) loadActiveRules(subject learn.SubjectType) ([]*learn.AdaptationRule, error) {
	// TODO: Query AdaptRule service for active rules matching this subject
	return nil, nil
}

func (e *AdaptiveEngine) loadOrCreateMastery(skillId string, sample *learn.Interaction) (*learn.SkillMastery, error) {
	// TODO: Query Mastery service by studentId + skillId, create if not found
	return &learn.SkillMastery{SkillId: skillId}, nil
}

func (e *AdaptiveEngine) loadSkills(subject learn.SubjectType) []*learn.Skill {
	// TODO: Query Skill service by subject
	return nil
}

func (e *AdaptiveEngine) loadActivities(subject learn.SubjectType, grade learn.GradeLevel) []*learn.Activity {
	// TODO: Query Activity service by subject + grade
	return nil
}

func (e *AdaptiveEngine) loadEngagement(studentId string) *learn.EngagementMetric {
	// TODO: Query Engage service by studentId
	return nil
}

func (e *AdaptiveEngine) loadGrowth(studentId string, subject learn.SubjectType) *learn.GrowthRecord {
	// TODO: Query Growth service by studentId + subject
	return nil
}

func (e *AdaptiveEngine) loadPeerContext(classroomId string) *CohortSummary {
	// TODO: Query Cohort service for latest classroom snapshot
	return nil
}

func findMastery(skillId string, masteries []*learn.SkillMastery) *learn.SkillMastery {
	for _, m := range masteries {
		if m.SkillId == skillId {
			return m
		}
	}
	return nil
}

func totalTimeSpent(interactions []*learn.Interaction) int32 {
	total := int32(0)
	for _, i := range interactions {
		total += i.TimeSpentSeconds
	}
	return total
}
