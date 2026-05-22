/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * Service helper functions for loading/saving StudentProfile and EvalImport.
 * Uses common.GetEntity/PutEntity/PostEntity for inter-service calls.
 */
package evalimports

import (
	"fmt"
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
	"time"
)

const (
	profileServiceName = "Profile"
	profileServiceArea = byte(20)
)

// loadCurrentProfile fetches the StudentProfile for a given studentId.
func loadCurrentProfile(studentId string, vnic ifs.IVNic) *learn.StudentProfile {
	result, err := common.GetEntity(profileServiceName, profileServiceArea,
		&learn.StudentProfile{StudentId: studentId}, vnic)
	if err != nil || result == nil {
		return nil
	}
	return result.(*learn.StudentProfile)
}

// loadOrCreateProfile fetches the StudentProfile or creates an empty one.
func loadOrCreateProfile(studentId string, vnic ifs.IVNic) *learn.StudentProfile {
	profile := loadCurrentProfile(studentId, vnic)
	if profile != nil {
		return profile
	}

	now := time.Now().Unix()
	profile = &learn.StudentProfile{
		ProfileId:   fmt.Sprintf("PROF-%s", studentId),
		StudentId:   studentId,
		CreatedDate: now,
		LastUpdated: now,
	}

	_, err := common.PostEntity(profileServiceName, profileServiceArea, profile, vnic)
	if err != nil {
		vnic.Resources().Logger().Error("Failed to create profile for student ", studentId, ": ", err.Error())
	}
	return profile
}

// saveProfile updates a StudentProfile via PUT.
func saveProfile(profile *learn.StudentProfile, vnic ifs.IVNic) error {
	return common.PutEntity(profileServiceName, profileServiceArea, profile, vnic)
}

// saveEvalImport updates an EvalImport record via PUT.
func saveEvalImport(eval *learn.EvalImport, vnic ifs.IVNic) {
	err := common.PutEntity(ServiceName, ServiceArea, eval, vnic)
	if err != nil {
		vnic.Resources().Logger().Error("Failed to save EvalImport ", eval.ImportId, ": ", err.Error())
	}
}

// loadEvalsForStudent fetches all EvalImports for a given studentId.
func loadEvalsForStudent(studentId string, vnic ifs.IVNic) []*learn.EvalImport {
	results, err := common.GetEntities(ServiceName, ServiceArea,
		&learn.EvalImport{StudentId: studentId}, vnic)
	if err != nil || len(results) == 0 {
		return nil
	}
	var evals []*learn.EvalImport
	for _, r := range results {
		if ev, ok := r.(*learn.EvalImport); ok {
			evals = append(evals, ev)
		}
	}
	return evals
}

// applyToProfile applies accepted findings to the StudentProfile.
func (h *evalImportHandler) applyToProfile(eval *learn.EvalImport, vnic ifs.IVNic) error {
	profile := loadCurrentProfile(eval.StudentId, vnic)
	if profile == nil {
		return fmt.Errorf("no profile found for student %s", eval.StudentId)
	}

	if err := ApplyFindingsToProfile(eval, profile); err != nil {
		return err
	}

	profile.LastUpdated = time.Now().Unix()
	return saveProfile(profile, vnic)
}
