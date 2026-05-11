/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package mocks

import (
	"fmt"
	"github.com/saichler/l8learn/go/types/learn"
	"math/rand"
	"time"
)

func generateEvalImports(store *MockDataStore) []*learn.EvalImport {
	var list []*learn.EvalImport
	now := time.Now().Unix()

	docTypes := []learn.EvalDocumentType{
		learn.EvalDocumentType_EVAL_DOCUMENT_TYPE_SPEECH,
		learn.EvalDocumentType_EVAL_DOCUMENT_TYPE_OCCUPATIONAL_THERAPY,
		learn.EvalDocumentType_EVAL_DOCUMENT_TYPE_PSYCHOLOGICAL,
		learn.EvalDocumentType_EVAL_DOCUMENT_TYPE_IEP,
		learn.EvalDocumentType_EVAL_DOCUMENT_TYPE_READING_SPECIALIST,
	}

	professionals := []string{
		"Dr. Sarah Miller, SLP", "Dr. James Chen, OTR/L",
		"Dr. Patricia Evans, PhD", "Ms. Rachel Kim, M.Ed.",
		"Dr. David Park, Reading Specialist",
	}

	for i := 0; i < 10; i++ {
		docType := docTypes[i%len(docTypes)]
		accepted := int32(2 + rand.Intn(3))
		rejected := int32(rand.Intn(2))

		findings := []*learn.EvalFinding{
			{
				FindingId:      fmt.Sprintf("EF-%04d-1", i+1),
				ProfileSection: "speech",
				ProfileField:   "expressive_language",
				CurrentValue:   "developing",
				NewValue:       "age_appropriate",
				SourceText:     "Expressive language skills are now within normal limits",
				Confidence:     0.92,
				Status:         learn.EvalFindingStatus_EVAL_FINDING_STATUS_ACCEPTED,
			},
			{
				FindingId:      fmt.Sprintf("EF-%04d-2", i+1),
				ProfileSection: "attention",
				ProfileField:   "focus_academic_task_minutes",
				CurrentValue:   "8",
				NewValue:       "12",
				SourceText:     "Sustained attention improved to 12 minutes during structured tasks",
				Confidence:     0.88,
				Status:         learn.EvalFindingStatus_EVAL_FINDING_STATUS_PENDING,
			},
		}

		list = append(list, &learn.EvalImport{
			ImportId:         fmt.Sprintf("EVAL-%03d", i+1),
			StudentId:        pickRef(store.StudentIDs, i*3),
			UploadedBy:       pickRef(store.GuardianIDs, i),
			DocumentType:     docType,
			ProfessionalName: professionals[i%len(professionals)],
			EvaluationDate:   now - int64(rand.Intn(180*24*3600)),
			Findings:         findings,
			AllReviewed:      i < 5,
			AcceptedCount:    accepted,
			RejectedCount:    rejected,
			AppliedToProfile: i < 3,
		})
	}
	return list
}
