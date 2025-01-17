package mappers

import (
	"github.com/mandaputtra/projectsprint-projects2/services/ms-activity-svc/dtos"
	"github.com/mandaputtra/projectsprint-projects2/services/ms-activity-svc/models"
)

func MapActivityModelToResponse(activityModel *models.Activity) *dtos.ActivityResponseDTO {
	return &dtos.ActivityResponseDTO{
		ActivityId:        activityModel.ID,
		ActivityType:      activityModel.ActivityTypeName,
		DoneAt:            activityModel.DoneAt,
		DurationInMinutes: activityModel.DurationInMinutes,
		CaloriesBurned:    activityModel.CaloriesBurned,
		CreatedAt:         activityModel.CreatedAt.String(),
		UpdatedAt:         activityModel.UpdatedAt.String(),
	}
}
