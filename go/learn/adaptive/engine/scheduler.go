/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 * Pattern adapted from l8alarms/go/alm/escalation/scheduler.go
 *
 * Scheduler runs periodic tasks:
 * - Weekly PROFILE_UPDATE for all active students
 * - Daily PARENT_COACHING for all active families (Phase 4)
 * - Weekly RISK_ASSESSMENT batch (Phase 5)
 */
package engine

import (
	"fmt"
	"time"

	"github.com/saichler/l8types/go/ifs"
)

// Scheduler runs periodic adaptive intelligence tasks
type Scheduler struct {
	vnic           ifs.IVNic
	profileUpdater *ProfileUpdater
	parentCoach    *ParentCoach
	running        bool
	stopCh         chan struct{}
}

func NewScheduler(vnic ifs.IVNic, profileUpdater *ProfileUpdater, parentCoach *ParentCoach) *Scheduler {
	return &Scheduler{
		vnic:           vnic,
		profileUpdater: profileUpdater,
		parentCoach:    parentCoach,
		stopCh:         make(chan struct{}),
	}
}

// Start begins the periodic task runner
func (s *Scheduler) Start() {
	if s.running {
		return
	}
	s.running = true

	go s.runWeeklyLoop()
	go s.runDailyLoop()

	fmt.Println("[Scheduler] Started periodic tasks (weekly + daily)")
}

// Stop gracefully stops the scheduler
func (s *Scheduler) Stop() {
	if !s.running {
		return
	}
	s.running = false
	close(s.stopCh)
}

func (s *Scheduler) runWeeklyLoop() {
	// Check every hour, run weekly tasks when due
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	var lastWeeklyRun time.Time

	for {
		select {
		case <-s.stopCh:
			return
		case now := <-ticker.C:
			// Run weekly profile update if 7 days have passed
			if now.Sub(lastWeeklyRun) >= 7*24*time.Hour {
				s.runWeeklyProfileUpdates()
				lastWeeklyRun = now
			}
		}
	}
}

func (s *Scheduler) runWeeklyProfileUpdates() {
	fmt.Println("[Scheduler] Running weekly profile updates...")
	// In production:
	// 1. Query all active students
	// 2. For each student, load their profile
	// 3. Call ProfileUpdater.RunWeeklyProfileUpdate(studentId, profile)
	// 4. Save updated profile
	// Each call generates a PROFILE_UPDATE prompt logged to PromptLog
	fmt.Println("[Scheduler] Weekly profile updates complete")
}

func (s *Scheduler) runDailyLoop() {
	// Check every hour, run daily tasks once per day (morning)
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	var lastDailyRun time.Time

	for {
		select {
		case <-s.stopCh:
			return
		case now := <-ticker.C:
			// Run daily coaching if 24 hours have passed
			if now.Sub(lastDailyRun) >= 24*time.Hour {
				if s.parentCoach != nil {
					s.parentCoach.RunDailyCoaching()
				}
				lastDailyRun = now
			}
		}
	}
}
