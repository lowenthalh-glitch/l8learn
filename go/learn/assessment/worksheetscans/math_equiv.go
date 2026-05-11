/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package worksheetscans

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

const epsilon = 0.0001

// AreEquivalent checks if two mathematical expressions represent the same value.
// Handles: integers, decimals, fractions, mixed numbers, percentages.
// Examples: "5/4" == "1 1/4" == "1.25" == "125%"
func AreEquivalent(studentAnswer, correctAnswer string) bool {
	studentVal, okS := parseToFloat(normalize(studentAnswer))
	correctVal, okC := parseToFloat(normalize(correctAnswer))
	if okS && okC {
		return math.Abs(studentVal-correctVal) < epsilon
	}
	// If numeric parsing fails, fall back to string comparison
	return strings.EqualFold(
		strings.TrimSpace(studentAnswer),
		strings.TrimSpace(correctAnswer),
	)
}

func normalize(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, ",", "")   // remove thousands separators
	s = strings.ReplaceAll(s, "$", "")   // remove currency symbols
	s = strings.ReplaceAll(s, " ", " ")  // normalize whitespace (non-breaking)
	return s
}

// parseToFloat attempts to interpret a string as a numeric value
func parseToFloat(s string) (float64, bool) {
	// Try direct float
	if v, err := strconv.ParseFloat(s, 64); err == nil {
		return v, true
	}

	// Try percentage: "75%" → 0.75
	if strings.HasSuffix(s, "%") {
		pct := strings.TrimSuffix(s, "%")
		if v, err := strconv.ParseFloat(pct, 64); err == nil {
			return v / 100, true
		}
	}

	// Try mixed number: "1 1/4" or "2 3/8"
	if v, ok := parseMixedNumber(s); ok {
		return v, true
	}

	// Try fraction: "5/4", "3/8"
	if v, ok := parseFraction(s); ok {
		return v, true
	}

	return 0, false
}

var fractionRegex = regexp.MustCompile(`^(-?\d+)\s*/\s*(\d+)$`)
var mixedRegex = regexp.MustCompile(`^(-?\d+)\s+(\d+)\s*/\s*(\d+)$`)

func parseFraction(s string) (float64, bool) {
	matches := fractionRegex.FindStringSubmatch(s)
	if matches == nil {
		return 0, false
	}
	num, _ := strconv.ParseFloat(matches[1], 64)
	den, _ := strconv.ParseFloat(matches[2], 64)
	if den == 0 {
		return 0, false
	}
	return num / den, true
}

func parseMixedNumber(s string) (float64, bool) {
	matches := mixedRegex.FindStringSubmatch(s)
	if matches == nil {
		return 0, false
	}
	whole, _ := strconv.ParseFloat(matches[1], 64)
	num, _ := strconv.ParseFloat(matches[2], 64)
	den, _ := strconv.ParseFloat(matches[3], 64)
	if den == 0 {
		return 0, false
	}
	if whole < 0 {
		return whole - (num / den), true
	}
	return whole + (num/den), true
}

// GradeAnswer compares a student's extracted answer against the correct answer.
// Returns: correct (bool), confidence adjustment (float64), reasoning (string)
func GradeAnswer(extracted, correct string) (bool, string) {
	if extracted == "" {
		return false, "no answer extracted"
	}

	// Direct string match (case-insensitive)
	if strings.EqualFold(strings.TrimSpace(extracted), strings.TrimSpace(correct)) {
		return true, "exact match"
	}

	// Math equivalence check
	if AreEquivalent(extracted, correct) {
		return true, "mathematically equivalent"
	}

	return false, "incorrect"
}
