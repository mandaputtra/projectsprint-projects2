package dtos

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

type ActivityRequestDTO struct {
	ActivityType      string `json:"activityType"`
	DoneAt            string `json:"doneAt"`
	DurationInMinutes int    `json:"durationInMinutes"`
}

var activityTypes = []string{"Walking", "Yoga", "Stretching", "Cycling", "Swimming", "Dancing", "Hiking", "Running", "HIIT", "JumpRope"}

func ValidateActivityRequest(dto ActivityRequestDTO) error {
	// Validasi activityType
	if dto.ActivityType == "" {
		return errors.New("activityType is required")
	}
	if !contains(activityTypes, dto.ActivityType) {
		return fmt.Errorf("activityType must be one of %v", activityTypes)
	}

	// Validasi doneAt
	if dto.DoneAt == "" {
		return errors.New("doneAt is required")
	}
	if !isValidISODate(dto.DoneAt) {
		return errors.New("doneAt must be a valid ISO date")
	}

	// Validasi durationInMinutes
	if dto.DurationInMinutes < 1 {
		return errors.New("durationInMinutes must be at least 1")
	}

	return nil
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func isValidISODate(date string) bool {
	isoDatePattern := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?Z$`
	matched, _ := regexp.MatchString(isoDatePattern, date)
	if !matched {
		return false
	}
	_, err := time.Parse(time.RFC3339, date)
	return err == nil
}
