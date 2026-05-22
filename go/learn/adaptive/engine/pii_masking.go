/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 * Pattern adapted from l8agent/go/masking/proxy.go
 */
package engine

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
)

// PIIReport contains the results of scanning text for PII
type PIIReport struct {
	HasNames bool
	HasPII   bool
	Fields   []string
}

// TokenMap maintains mask/unmask mappings for a single request
type TokenMap struct {
	mu       sync.Mutex
	tokens   map[string]string // "[NAME_1]" → "Jake Martinez"
	counters map[string]int    // "NAME" → next index
}

func NewTokenMap() *TokenMap {
	return &TokenMap{
		tokens:   make(map[string]string),
		counters: make(map[string]int),
	}
}

func (m *TokenMap) Mask(prefix string, value string) string {
	m.mu.Lock()
	defer m.mu.Unlock()
	idx := m.counters[prefix]
	m.counters[prefix] = idx + 1
	token := fmt.Sprintf("[%s_%d]", prefix, idx+1)
	m.tokens[token] = value
	return token
}

func (m *TokenMap) Unmask(text string) string {
	m.mu.Lock()
	defer m.mu.Unlock()
	for token, value := range m.tokens {
		text = strings.ReplaceAll(text, token, value)
	}
	return text
}

// Compiled regexes for PII detection (from l8agent/masking/proxy.go)
var (
	ssnRegex       = regexp.MustCompile(`\b\d{3}-\d{2}-\d{4}\b`)
	emailRegex     = regexp.MustCompile(`\b[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}\b`)
	phoneRegex     = regexp.MustCompile(`\(?\d{3}\)?[\s\-]?\d{3}[\s\-]?\d{4}`)
	dobRegex       = regexp.MustCompile(`\b\d{1,2}/\d{1,2}/\d{4}\b`)
	addressRegex   = regexp.MustCompile(`\b\d{1,5}\s+\w+\s+(St|Ave|Blvd|Dr|Rd|Ln|Way|Ct|Pl|Cir|Pkwy)\.?\b`)
	mrnRegex       = regexp.MustCompile(`(?i)\b(?:MRN|Medical Record|Record #|Patient ID)[:\s#]*[\w\-]+\b`)
	insuranceRegex = regexp.MustCompile(`(?i)\b(?:Policy|Member ID|Group #|Insurance ID)[:\s#]*[\w\-]+\b`)
	dateFullRegex  = regexp.MustCompile(`\b(?:January|February|March|April|May|June|July|August|September|October|November|December)\s+\d{1,2},?\s+\d{4}\b`)
)

// PIIMasker scans and masks PII in text
type PIIMasker struct {
	knownNames []string // Student names to detect
}

func NewPIIMasker() *PIIMasker {
	return &PIIMasker{}
}

func (p *PIIMasker) SetKnownNames(names []string) {
	p.knownNames = names
}

// Scan checks text for PII without masking it
func (p *PIIMasker) Scan(text string) *PIIReport {
	report := &PIIReport{}

	for _, name := range p.knownNames {
		if name != "" && strings.Contains(text, name) {
			report.HasNames = true
			report.Fields = append(report.Fields, "student_name: "+name)
		}
	}

	if ssnRegex.MatchString(text) {
		report.HasPII = true
		report.Fields = append(report.Fields, "ssn_pattern")
	}
	if emailRegex.MatchString(text) {
		report.Fields = append(report.Fields, "email_pattern")
	}
	if phoneRegex.MatchString(text) {
		report.Fields = append(report.Fields, "phone_pattern")
	}
	if dobRegex.MatchString(text) {
		report.HasPII = true
		report.Fields = append(report.Fields, "dob_pattern")
	}

	report.HasPII = report.HasPII || report.HasNames
	return report
}

// MaskText replaces PII patterns with tokens using the instance's knownNames.
func (p *PIIMasker) MaskText(text string, tokenMap *TokenMap) string {
	return maskTextInternal(text, tokenMap, p.knownNames)
}

// MaskTextWithNames replaces PII patterns using per-request names (thread-safe).
// Use this instead of SetKnownNames + MaskText when processing concurrently.
func (p *PIIMasker) MaskTextWithNames(text string, tokenMap *TokenMap, knownNames []string) string {
	return maskTextInternal(text, tokenMap, knownNames)
}

func maskTextInternal(text string, tokenMap *TokenMap, knownNames []string) string {
	// Mask known names (longest first to avoid partial matches)
	for _, name := range knownNames {
		if name != "" && len(name) > 1 && strings.Contains(text, name) {
			token := tokenMap.Mask("PERSON", name)
			text = strings.ReplaceAll(text, name, token)
		}
	}

	// Mask SSN patterns
	text = ssnRegex.ReplaceAllStringFunc(text, func(match string) string {
		return "[MASKED_SSN]"
	})

	// Mask email patterns
	text = emailRegex.ReplaceAllStringFunc(text, func(match string) string {
		return tokenMap.Mask("EMAIL", match)
	})

	// Mask phone patterns
	text = phoneRegex.ReplaceAllStringFunc(text, func(match string) string {
		return tokenMap.Mask("PHONE", match)
	})

	// Mask DOB patterns (MM/DD/YYYY)
	text = dobRegex.ReplaceAllStringFunc(text, func(match string) string {
		return tokenMap.Mask("DOB", match)
	})

	// Mask full date strings (January 15, 2020)
	text = dateFullRegex.ReplaceAllStringFunc(text, func(match string) string {
		return tokenMap.Mask("DATE", match)
	})

	// Mask address patterns
	text = addressRegex.ReplaceAllStringFunc(text, func(match string) string {
		return "[MASKED_ADDRESS]"
	})

	// Mask medical record numbers
	text = mrnRegex.ReplaceAllStringFunc(text, func(match string) string {
		return "[MASKED_MRN]"
	})

	// Mask insurance identifiers
	text = insuranceRegex.ReplaceAllStringFunc(text, func(match string) string {
		return "[MASKED_INSURANCE]"
	})

	return text
}
