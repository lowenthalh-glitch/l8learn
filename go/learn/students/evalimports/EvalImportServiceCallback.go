/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package evalimports

import (
	"fmt"
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/learn/adaptive/engine"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
	"time"
)

// evalImportHandler holds LLM dependencies injected after service activation.
type evalImportHandler struct {
	llmClient engine.LLMClient
	masker    *engine.PIIMasker
}

// Global handler instance — set during Activate, LLM deps injected later via SetLLMClient.
var handler = &evalImportHandler{}

// SetLLMClient injects the LLM client and masker after initialization.
func SetLLMClient(client engine.LLMClient, masker *engine.PIIMasker) {
	handler.llmClient = client
	handler.masker = masker
	fmt.Printf("[EvalImport] LLM client injected: %v\n", client != nil)
}

func newEvalImportServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.EvalImport{}, vnic).
		BeforeAction(func(elem interface{}, action ifs.Action, vnic ifs.IVNic) error {
			if action == ifs.POST {
				eval := elem.(*learn.EvalImport)
				common.GenerateID(&eval.ImportId)
				eval.ProcessingStatus = learn.EvalProcessingStatus_EVAL_PROCESSING_PENDING
				fmt.Printf("[EvalImport] POST for %s, filePath=%s\n", eval.ImportId, eval.FilePath)
				// Launch async cleaning goroutine (no LLM call — clean only)
				// 500ms delay ensures the POST persistence completes before the goroutine PUTs
				go func() {
					time.Sleep(500 * time.Millisecond)
					handler.cleanEvalImport(eval, vnic)
				}()
			}
			return nil
		}).
		Require(func(v interface{}) string { return v.(*learn.EvalImport).ImportId }, "ImportId").
		Require(func(v interface{}) string { return v.(*learn.EvalImport).StudentId }, "StudentId").
		After(func(elem interface{}, action ifs.Action, vnic ifs.IVNic) error {
			eval := elem.(*learn.EvalImport)
			// Batch LLM processing: triggered by single PUT with SUBMITTED
			if action == ifs.PUT && eval.ProcessingStatus == learn.EvalProcessingStatus_EVAL_PROCESSING_SUBMITTED {
				if handler.llmClient != nil {
					go handler.processAllEvals(eval, vnic)
				} else {
					fmt.Println("[EvalImport] WARNING: llmClient is nil, skipping batch processing")
				}
				return nil
			}
			// Profile application
			if action == ifs.PUT && eval.AppliedToProfile {
				return handler.applyToProfile(eval, vnic)
			}
			return nil
		}).
		Build()
}
