/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * Eval processor — two-phase pipeline:
 *   cleanEvalImport()  — runs async after POST, extracts/strips/masks/saves
 *   processAllEvals()  — runs async after PUT/SUBMITTED, batch LLM call
 */
package evalimports

import (
	"encoding/json"
	"fmt"
	"github.com/saichler/l8learn/go/learn/adaptive/engine"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

// cleanEvalImport extracts PDF text, strips headers, masks PII, saves cleaned file.
// No LLM call — guardian must review before processing.
func (h *evalImportHandler) cleanEvalImport(eval *learn.EvalImport, vnic ifs.IVNic) {
	log := func(msg string, args ...interface{}) {
		fmt.Printf("[EvalClean %s] %s\n", eval.ImportId, fmt.Sprintf(msg, args...))
	}

	log("Starting cleaning for student %s", eval.StudentId)
	eval.ProcessingStatus = learn.EvalProcessingStatus_EVAL_PROCESSING_EXTRACTING
	saveEvalImport(eval, vnic)

	// 1. Read uploaded file
	fileData, err := ReadUploadedFile(eval.FilePath, vnic)
	if err != nil {
		h.failEval(eval, "Failed to read file: "+err.Error(), vnic)
		return
	}
	log("File read: %d bytes", len(fileData))

	// 2. Extract text from PDF
	pdfText, err := ExtractTextFromPDF(fileData)
	if err != nil {
		h.failEval(eval, "Failed to extract text: "+err.Error(), vnic)
		return
	}
	if pdfText == "" {
		h.failEval(eval, "PDF has no extractable text layer", vnic)
		return
	}
	log("PDF text extracted: %d chars", len(pdfText))

	// 3. Strip document headers/demographics
	pdfText = StripDocumentHeaders(pdfText)
	log("After header stripping: %d chars", len(pdfText))

	// 4. Build masking context and mask PII
	knownNames := BuildMaskingContext(eval.StudentId, vnic)
	log("Masking context: %d known names: %v", len(knownNames), knownNames)
	tokenMap := engine.NewTokenMap()
	pdfText = h.masker.MaskTextWithNames(pdfText, tokenMap, knownNames)
	log("After PII masking: %d chars", len(pdfText))

	// 5. Save cleaned+masked text and delete original PDF
	cleanedPath, writeErr := WriteCleanedText(eval.FilePath, pdfText, eval.ImportId, vnic)
	if writeErr != nil {
		log("WARNING: Failed to save cleaned text: %s", writeErr.Error())
	} else {
		DeleteOriginalFile(eval.FilePath)
		eval.FilePath = cleanedPath
		log("Saved cleaned text to %s, original PDF deleted", cleanedPath)
	}

	// 6. Set status to CLEANED — ready for guardian review
	eval.ProcessingStatus = learn.EvalProcessingStatus_EVAL_PROCESSING_CLEANED
	eval.ErrorMessage = ""
	saveEvalImport(eval, vnic)
	log("Cleaning COMPLETE — ready for guardian review")
}

// processAllEvals loads all CLEANED evals for a student, sends them to Claude
// in one batch, and stores findings on the trigger eval.
func (h *evalImportHandler) processAllEvals(triggerEval *learn.EvalImport, vnic ifs.IVNic) {
	log := func(msg string, args ...interface{}) {
		fmt.Printf("[EvalBatch %s] %s\n", triggerEval.ImportId, fmt.Sprintf(msg, args...))
	}

	studentId := triggerEval.StudentId
	log("Starting batch processing for student %s", studentId)

	// 1. Load all evals for this student
	allEvals := loadEvalsForStudent(studentId, vnic)
	log("Found %d total evals for student", len(allEvals))

	// 2. Collect CLEANED evals and read their text
	type evalDoc struct {
		eval *learn.EvalImport
		text string
	}
	var docs []evalDoc
	for _, ev := range allEvals {
		if ev.ProcessingStatus == learn.EvalProcessingStatus_EVAL_PROCESSING_CLEANED ||
			ev.ImportId == triggerEval.ImportId {
			text, err := ReadUploadedFile(ev.FilePath, vnic)
			if err != nil {
				log("WARNING: Failed to read cleaned file for %s: %s", ev.ImportId, err.Error())
				continue
			}
			docs = append(docs, evalDoc{eval: ev, text: string(text)})
		}
	}

	if len(docs) == 0 {
		h.failEval(triggerEval, "No cleaned evaluations found to process", vnic)
		return
	}
	log("Processing %d cleaned evaluations", len(docs))

	// 3. Build multi-document prompt
	var documents []engine.EvalDocument
	totalChars := 0
	for i, d := range docs {
		documents = append(documents, engine.EvalDocument{
			Text:         d.text,
			DocumentType: d.eval.DocumentType.String(),
			FileName:     fmt.Sprintf("Document %d", i+1),
		})
		totalChars += len(d.text)
	}

	// 4. Token limit check
	log("Total document text: %d chars (~%d tokens)", totalChars, totalChars/4)
	if totalChars > 400000 {
		h.failEval(triggerEval, fmt.Sprintf("Combined evaluations too large: %d chars (max 400000)", totalChars), vnic)
		return
	}
	if totalChars > 100000 {
		log("WARNING: Large prompt — %d chars may affect quality", totalChars)
	}

	// 5. Load current profile
	currentProfile := loadOrCreateProfile(studentId, vnic)
	profileJSON, _ := json.Marshal(currentProfile)
	log("Current profile: %d bytes", len(profileJSON))

	// 6. Build prompt and call LLM
	systemPrompt, userMessage := engine.BuildBatchEvalPrompt(documents, string(profileJSON))
	log("--- SYSTEM PROMPT (first 2000 chars) ---\n%s\n--- END SYSTEM PROMPT ---", truncate(systemPrompt, 2000))
	log("--- USER MESSAGE (first 3000 chars) ---\n%s\n--- END USER MESSAGE ---", truncate(userMessage, 3000))

	log("Calling LLM (mode=%s)...", h.llmClient.GetMode().String())
	response, err := h.llmClient.Call(
		learn.LLMPromptType_LLM_PROMPT_TYPE_EVAL_IMPORT,
		systemPrompt, userMessage, studentId,
	)
	if err != nil {
		h.failEval(triggerEval, "LLM call failed: "+err.Error(), vnic)
		return
	}
	log("LLM response: %d chars", len(response))
	log("--- LLM RESPONSE ---\n%s\n--- END LLM RESPONSE ---", response)

	// 7. Parse findings
	findings, contradictions := parseEvalResponse(response)
	log("Parsed: %d findings, %d contradictions", len(findings), len(contradictions))
	for i, f := range findings {
		log("  Finding %d: %s.%s = %q (conf=%.2f)", i+1, f.ProfileSection, f.ProfileField, f.NewValue, f.Confidence)
	}

	// 8. Store findings on trigger eval
	triggerEval.Findings = findings
	triggerEval.Contradictions = contradictions
	triggerEval.ProcessingStatus = learn.EvalProcessingStatus_EVAL_PROCESSING_COMPLETE
	triggerEval.ErrorMessage = ""
	saveEvalImport(triggerEval, vnic)

	// 9. Set all other CLEANED evals to COMPLETE (no findings on them)
	for _, d := range docs {
		if d.eval.ImportId != triggerEval.ImportId {
			d.eval.ProcessingStatus = learn.EvalProcessingStatus_EVAL_PROCESSING_COMPLETE
			saveEvalImport(d.eval, vnic)
		}
	}

	log("Batch processing COMPLETE — %d findings on trigger eval", len(findings))
}

func (h *evalImportHandler) failEval(eval *learn.EvalImport, msg string, vnic ifs.IVNic) {
	eval.ProcessingStatus = learn.EvalProcessingStatus_EVAL_PROCESSING_FAILED
	eval.ErrorMessage = msg
	saveEvalImport(eval, vnic)
	fmt.Printf("[EvalProcessor %s] FAILED: %s\n", eval.ImportId, msg)
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "... [TRUNCATED]"
}
