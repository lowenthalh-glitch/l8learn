/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * Shared LLM response parsing utilities.
 * Used by eval, schedule, and lesson response parsers.
 */
package engine

import (
	"encoding/json"
	"fmt"
	"strings"
)

// StripCodeFences removes markdown code fences and any text before/after the JSON.
// Handles: ```json ... ```, text before fences, and text after fences.
func StripCodeFences(s string) string {
	s = strings.TrimSpace(s)

	// If response contains code fences, extract content between them
	fenceStart := strings.Index(s, "```")
	if fenceStart >= 0 {
		// Skip the opening fence line
		afterFence := s[fenceStart+3:]
		newline := strings.Index(afterFence, "\n")
		if newline >= 0 {
			afterFence = afterFence[newline+1:]
		}
		// Find closing fence
		fenceEnd := strings.LastIndex(afterFence, "```")
		if fenceEnd >= 0 {
			s = afterFence[:fenceEnd]
		} else {
			s = afterFence
		}
		return strings.TrimSpace(s)
	}

	// No fences — try to find JSON object/array directly
	jsonStart := strings.IndexAny(s, "{[")
	if jsonStart > 0 {
		s = s[jsonStart:]
	}
	return strings.TrimSpace(s)
}

// ParseLLMResponse strips code fences and unmarshals JSON into the target struct.
func ParseLLMResponse(jsonResponse string, target interface{}) error {
	cleaned := StripCodeFences(jsonResponse)
	if err := json.Unmarshal([]byte(cleaned), target); err != nil {
		preview := cleaned
		if len(preview) > 200 {
			preview = preview[:200]
		}
		return fmt.Errorf("JSON parse error: %w (first 200 chars: %s)", err, preview)
	}
	return nil
}
