/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package messages

import (
	"fmt"
	"strings"

	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

const maxMessagesPerSession = 20

func newCollabMessageServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.CollabMessage{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.CollabMessage).MessageId }, "MessageId").
		Require(func(v interface{}) string { return v.(*learn.CollabMessage).GroupId }, "GroupId").
		Require(func(v interface{}) string { return v.(*learn.CollabMessage).SenderId }, "SenderId").
		Custom(moderateMessage).
		After(onMessageSent).
		Build()
}

// moderateMessage runs AI moderation BEFORE the message is stored/delivered
func moderateMessage(elem interface{}, vnic ifs.IVNic) error {
	msg := elem.(*learn.CollabMessage)

	// Block personal information
	if containsPersonalInfo(msg.Content) {
		msg.Moderation = learn.ModerationAction_MODERATION_ACTION_BLOCKED
		msg.ModerationReason = "Message contains personal information"
		return fmt.Errorf("message blocked: contains personal information")
	}

	// Block inappropriate language
	if containsInappropriate(msg.Content) {
		msg.Moderation = learn.ModerationAction_MODERATION_ACTION_BLOCKED
		msg.ModerationReason = "Message contains inappropriate language"
		return fmt.Errorf("message blocked: inappropriate content")
	}

	// Coach: if sharing a direct answer, nudge to explain approach instead
	if looksLikeDirectAnswer(msg.Content) {
		msg.Moderation = learn.ModerationAction_MODERATION_ACTION_COACHED
		msg.ModerationReason = "Try explaining HOW you solved it instead of giving the answer"
		// Don't block — deliver with coaching note
	} else {
		msg.Moderation = learn.ModerationAction_MODERATION_ACTION_APPROVED
	}

	return nil
}

// onMessageSent handles post-delivery logic
func onMessageSent(elem interface{}, action ifs.Action, notify bool, vnic ifs.IVNic) (interface{}, bool, error) {
	msg := elem.(*learn.CollabMessage)

	// If message contains an explanation, award helper points
	if msg.ContainsExplanation {
		// TODO: Update sender's GroupMember.HelpsGiven counter
		// TODO: Award "helper" points to sender's EngagementMetric
	}

	// If type is QUESTION, check if AI can provide a hint
	if msg.Type == learn.ChatMessageType_CHAT_MESSAGE_TYPE_QUESTION {
		// TODO: AI generates hint without giving the answer
		// TODO: Post AI_COACH message to group
	}

	// If type is EXPLANATION, verify correctness
	if msg.Type == learn.ChatMessageType_CHAT_MESSAGE_TYPE_EXPLANATION {
		// TODO: AI checks if explanation is mathematically correct
		// If wrong, post gentle AI correction
	}

	_ = msg
	return nil, true, nil
}

func containsPersonalInfo(content string) bool {
	lower := strings.ToLower(content)
	// Basic checks — in production, use AI for nuanced detection
	patterns := []string{"my address is", "my phone", "i live at", "my number is"}
	for _, p := range patterns {
		if strings.Contains(lower, p) {
			return true
		}
	}
	return false
}

func containsInappropriate(content string) bool {
	// Placeholder — in production, use AI content moderation
	// or a profanity filter service
	_ = content
	return false
}

func looksLikeDirectAnswer(content string) bool {
	lower := strings.ToLower(content)
	// Heuristic: very short messages that look like just an answer
	if len(content) < 10 && !strings.Contains(lower, "because") && !strings.Contains(lower, "try") {
		// Could be just "24" or "B" — but could also be "thanks"
		// In production, AI classifies this more accurately
		return false
	}
	// Check for patterns like "the answer is X"
	giveawayPatterns := []string{"the answer is", "it's just", "answer:", "answer ="}
	for _, p := range giveawayPatterns {
		if strings.Contains(lower, p) {
			return true
		}
	}
	return false
}
