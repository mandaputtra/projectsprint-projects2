package dtos

type ActivityRequestDTO struct {
	ActivityType      string `json:"activityType"`
	DoneAt            string `json:"doneAt"`
	DurationInMinutes int    `json:"durationInMinutes"`
}
