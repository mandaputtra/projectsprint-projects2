package mappers

import (
	"github.com/mandaputtra/projectsprint-projects2/services/ms-activity-svc/dtos"
	"github.com/mandaputtra/projectsprint-projects2/services/ms-activity-svc/models"
)

func MapRequestToActivityModel(dto *dtos.ActivityRequestDTO) *models.Activity {
	return &models.Activity{}
}

func MapActivityModelToResponse(activityModel *models.Activity) *dtos.ActivityResponseDTO {
	return &dtos.ActivityResponseDTO{
		ID: activityModel.ID,
	}
}
