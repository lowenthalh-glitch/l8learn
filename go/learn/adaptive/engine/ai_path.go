/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package engine

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/saichler/l8learn/go/types/learn"
)

// AIPathRequest contains all 15 inputs for the AI decision
type AIPathRequest struct {
	Student      *learn.Student       `json:"student"`
	Path         *learn.LearningPath  `json:"path"`
	Mastery      []*learn.SkillMastery `json:"mastery"`
	Skills       []*learn.Skill       `json:"skills"`
	Interactions []*learn.Interaction  `json:"recentInteractions"`
	Engagement   *learn.EngagementMetric `json:"engagement"`
	Growth       *learn.GrowthRecord  `json:"growth"`
	RiskLevel    learn.RiskLevel      `json:"riskLevel"`
	PeerContext  *CohortSummary       `json:"peerContext"`
	Activities   []*learn.Activity    `json:"availableActivities"`
	RuleResult   *RuleResult          `json:"ruleResult"`
}

// CohortSummary is a simplified view of the class context for the AI
type CohortSummary struct {
	ClassAvgMastery   float64  `json:"classAvgMastery"`
	TopGapSkills      []string `json:"topGapSkills"`
	EffectiveActivities []string `json:"effectiveActivities"`
}

// AIPathResponse is the structured output from the AI
type AIPathResponse struct {
	NextActivities []AIActivityChoice `json:"nextActivities"`
	Reasoning      string             `json:"reasoning"`
}

// AIActivityChoice is a single AI decision for one activity slot
type AIActivityChoice struct {
	ActivityId string               `json:"activityId"`
	SkillId    string               `json:"skillId"`
	Difficulty learn.DifficultyLevel `json:"difficulty"`
	Reason     string               `json:"reason"`
}

// BuildPrompt constructs the system prompt + user message for the AI
func BuildPrompt(req *AIPathRequest) (string, string) {
	systemPrompt := `You are an adaptive learning engine. Your job is to choose the next activities for a student based on their current state, mastery levels, engagement patterns, and available activities.

Rules:
- Never assign an activity whose prerequisite skills are not at PROFICIENT or above
- If the student has an IEP or 504 plan, respect accommodation notes
- Balance challenge with confidence (don't stack hard activities back-to-back)
- If engagement is dropping, insert a game or interactive activity
- If a skill is plateauing, try a different activity type for that skill
- Include contingency notes for real-time adaptation

Return JSON with this structure:
{
  "nextActivities": [
    { "activityId": "...", "skillId": "...", "difficulty": 1-5, "reason": "..." }
  ],
  "reasoning": "Overall rationale for this sequence"
}`

	userMessage := buildUserContext(req)
	return systemPrompt, userMessage
}

func buildUserContext(req *AIPathRequest) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("STUDENT: %s %s, Grade %d\n",
		req.Student.FirstName, req.Student.LastName, req.Student.GradeLevel))

	if req.Student.HasIep {
		b.WriteString(fmt.Sprintf("IEP: Yes. Accommodations: %s\n", req.Student.AccommodationNotes))
	}
	if req.Student.Has_504Plan {
		b.WriteString(fmt.Sprintf("504 Plan: Yes. Accommodations: %s\n", req.Student.AccommodationNotes))
	}

	b.WriteString(fmt.Sprintf("\nSUBJECT: %s\n", req.Path.Subject.String()))
	b.WriteString(fmt.Sprintf("PATH STATUS: %s, Activities completed: %d, Skills mastered: %d\n",
		req.Path.Status.String(), req.Path.ActivitiesCompleted, req.Path.SkillsMastered))

	b.WriteString("\nSKILL MASTERY:\n")
	for _, m := range req.Mastery {
		skillName := findSkillName(m.SkillId, req.Skills)
		b.WriteString(fmt.Sprintf("  %s: %s (%.0f%% accuracy, %d attempts)\n",
			skillName, m.Level.String(), m.CurrentAccuracy*100, m.AttemptsCount))
	}

	if req.Engagement != nil {
		b.WriteString(fmt.Sprintf("\nENGAGEMENT: %s, avg %.0f min/session, streak %d days\n",
			req.Engagement.CurrentLevel.String(),
			req.Engagement.AvgSessionMinutes,
			req.Engagement.CurrentStreakDays))
	}

	if req.Growth != nil {
		b.WriteString(fmt.Sprintf("\nGROWTH: %.1fx expected (rating: %s)\n",
			req.Growth.GrowthVsExpected, req.Growth.Rating.String()))
	}

	if req.RiskLevel > learn.RiskLevel_RISK_LEVEL_ON_TRACK {
		b.WriteString(fmt.Sprintf("\nRISK: %s\n", req.RiskLevel.String()))
	}

	if req.RuleResult != nil && req.RuleResult.Matched {
		b.WriteString(fmt.Sprintf("\nRULE TRIGGERED: %s → strategy: %s\n",
			req.RuleResult.Rule.Name, req.RuleResult.Strategy.String()))
	}

	if req.PeerContext != nil {
		b.WriteString(fmt.Sprintf("\nPEER CONTEXT: class avg mastery %.0f%%\n",
			req.PeerContext.ClassAvgMastery*100))
		if len(req.PeerContext.TopGapSkills) > 0 {
			b.WriteString(fmt.Sprintf("  Class gaps: %s\n", strings.Join(req.PeerContext.TopGapSkills, ", ")))
		}
	}

	if len(req.Interactions) > 0 {
		b.WriteString("\nRECENT INTERACTIONS (last 10):\n")
		start := 0
		if len(req.Interactions) > 10 {
			start = len(req.Interactions) - 10
		}
		for _, i := range req.Interactions[start:] {
			b.WriteString(fmt.Sprintf("  %s: %s (%ds, %d hints)\n",
				i.SkillId, i.Result.String(), i.TimeSpentSeconds, i.HintsUsed))
		}
	}

	b.WriteString("\nAVAILABLE ACTIVITIES:\n")
	for _, a := range req.Activities {
		b.WriteString(fmt.Sprintf("  [%s] %s (%s, %s, skills: %s)\n",
			a.ActivityId, a.Name, a.ActivityType.String(),
			a.Difficulty.String(), strings.Join(a.SkillIds, ",")))
	}

	b.WriteString("\nGenerate the next 5 activities for today's session (target: 15 minutes).")
	return b.String()
}

// ParseAIResponse parses the AI's JSON response into structured output
func ParseAIResponse(response string) (*AIPathResponse, error) {
	// Strip markdown code fences if present
	response = strings.TrimSpace(response)
	if strings.HasPrefix(response, "```") {
		lines := strings.Split(response, "\n")
		response = strings.Join(lines[1:len(lines)-1], "\n")
	}

	var result AIPathResponse
	err := json.Unmarshal([]byte(response), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AI response: %w", err)
	}
	return &result, nil
}

func findSkillName(skillId string, skills []*learn.Skill) string {
	for _, s := range skills {
		if s.SkillId == skillId {
			return s.Name
		}
	}
	return skillId
}
