/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * Strips identifying headers/footers from evaluation documents.
 * Handles both "Label: Value" (same line) and "Label\nValue" (next line) formats.
 */
package evalimports

import (
	"regexp"
	"strings"
)

// Patterns that identify demographic/administrative lines to strip.
// Same-line patterns: "Label: Value" or "Label Value" on one line.
var sameLinePatterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)^\s*(Patient|Client|Student)\s*(Name|ID)?\s*[:]\s*.+$`),
	regexp.MustCompile(`(?i)^\s*(DOB|Date of Birth|Birth Date)\s*[:]\s*.+$`),
	regexp.MustCompile(`(?i)^\s*(MRN|Medical Record|Record #|Patient ID|Chart #)\s*[:]\s*.+$`),
	regexp.MustCompile(`(?i)^\s*(Date of Evaluation|Date of Report|Report Date)\s*[:]\s*.+$`),
	regexp.MustCompile(`(?i)^\s*(Referred by|Referring Physician)\s*[:]\s*.+$`),
	regexp.MustCompile(`(?i)^\s*(Insurance|Policy #|Member ID|Group #)\s*[:]\s*.+$`),
	regexp.MustCompile(`(?i)^\s*(Fax|Phone|Tel|Office)\s*[:]\s*\(?\d`),
	regexp.MustCompile(`(?i)^\s*Page\s+\d+\s+(of|/)\s+\d+\s*$`),
	regexp.MustCompile(`(?i)^\s*\d{1,5}\s+\w+\s+(Street|St|Avenue|Ave|Boulevard|Blvd|Drive|Dr|Road|Rd|Lane|Ln|Way|Court|Ct|Place|Pl|Circle|Cir|Parkway|Pkwy)\.?\s*,`),
}

// Label-only patterns: these are standalone labels where the value is on the NEXT line.
// When matched, strip both the label line AND the following non-empty line.
var labelOnlyPatterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)^\s*Student\s*Name\s*$`),
	regexp.MustCompile(`(?i)^\s*Patient\s*Name\s*$`),
	regexp.MustCompile(`(?i)^\s*Client\s*Name\s*$`),
	regexp.MustCompile(`(?i)^\s*Full\s*Name\s*$`),
	regexp.MustCompile(`(?i)^\s*Name\s*$`),
	regexp.MustCompile(`(?i)^\s*Age\s*$`),
	regexp.MustCompile(`(?i)^\s*Date\s*of\s*Birth\s*$`),
	regexp.MustCompile(`(?i)^\s*DOB\s*$`),
	regexp.MustCompile(`(?i)^\s*Grade\s*$`),
	regexp.MustCompile(`(?i)^\s*School\s*$`),
	regexp.MustCompile(`(?i)^\s*District\s*$`),
	regexp.MustCompile(`(?i)^\s*Parent(s)?\s*$`),
	regexp.MustCompile(`(?i)^\s*Guardian(s)?\s*$`),
	regexp.MustCompile(`(?i)^\s*Address\s*$`),
	regexp.MustCompile(`(?i)^\s*Phone\s*(Number)?\s*$`),
	regexp.MustCompile(`(?i)^\s*Email\s*$`),
	regexp.MustCompile(`(?i)^\s*Insurance\s*$`),
	regexp.MustCompile(`(?i)^\s*Policy\s*(Number|#)?\s*$`),
	regexp.MustCompile(`(?i)^\s*MRN\s*$`),
	regexp.MustCompile(`(?i)^\s*Medical\s*Record\s*(Number|#)?\s*$`),
	regexp.MustCompile(`(?i)^\s*Date\s*of\s*(Evaluation|Report|Assessment)\s*$`),
	regexp.MustCompile(`(?i)^\s*Evaluator\s*$`),
	regexp.MustCompile(`(?i)^\s*Referred\s*by\s*$`),
	regexp.MustCompile(`(?i)^\s*Referring\s*Physician\s*$`),
}

// StripDocumentHeaders removes identifying lines from evaluation text.
func StripDocumentHeaders(text string) string {
	lines := strings.Split(text, "\n")
	var result []string
	skipNext := false

	for i, line := range lines {
		if skipNext {
			trimmed := strings.TrimSpace(line)
			if trimmed != "" {
				skipNext = false // consumed the value line
				continue
			}
			// empty line — skip it too and keep looking for value
			continue
		}

		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			result = append(result, line)
			continue
		}

		// Check same-line patterns (Label: Value)
		sameLineMatch := false
		for _, pat := range sameLinePatterns {
			if pat.MatchString(trimmed) {
				sameLineMatch = true
				break
			}
		}
		if sameLineMatch {
			continue
		}

		// Check label-only patterns (Label\nValue on next line)
		labelMatch := false
		for _, pat := range labelOnlyPatterns {
			if pat.MatchString(trimmed) {
				labelMatch = true
				break
			}
		}
		if labelMatch {
			// Skip this label line AND the next non-empty line (the value)
			skipNext = true
			_ = i
			continue
		}

		result = append(result, line)
	}

	return strings.TrimSpace(strings.Join(result, "\n"))
}
