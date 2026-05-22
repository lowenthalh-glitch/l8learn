/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package schedules

import (
	"fmt"
	"github.com/saichler/l8learn/go/learn/adaptive/engine"
	"github.com/saichler/l8learn/go/types/learn"
)

type llmScheduleResponse struct {
	Days []llmDay `json:"days"`
}

type llmDay struct {
	DayName string     `json:"dayName"`
	Blocks  []llmBlock `json:"blocks"`
}

type llmBlock struct {
	BlockId        string `json:"blockId"`
	StartMinute    int32  `json:"startMinute"`
	DurationMinutes int32 `json:"durationMinutes"`
	ActivityType   string `json:"activityType"`
	Subject        string `json:"subject"`
	Description    string `json:"description"`
	ParentRole     string `json:"parentRole"`
	RequiresParent bool   `json:"requiresParent"`
}

func parseScheduleResponse(jsonResponse string) ([]*learn.ScheduleBlock, int32) {
	var resp llmScheduleResponse
	if err := engine.ParseLLMResponse(jsonResponse, &resp); err != nil {
		fmt.Printf("[ScheduleParser] %s\n", err.Error())
		return nil, 0
	}

	var blocks []*learn.ScheduleBlock
	academicCount := int32(0)
	for _, day := range resp.Days {
		for _, b := range day.Blocks {
			block := &learn.ScheduleBlock{
				BlockId:         b.BlockId,
				StartMinute:     b.StartMinute,
				DurationMinutes: b.DurationMinutes,
				ActivityType:    b.ActivityType,
				Description:     b.Description,
				ParentRole:      b.ParentRole,
				RequiresParent:  b.RequiresParent,
			}
			blocks = append(blocks, block)
			if b.ActivityType == "academic" || b.ActivityType == "therapy" {
				academicCount++
			}
		}
	}

	return blocks, academicCount
}
