package services

import (
	// "errors"
	// "fmt"

	// "github.com/gin-gonic/gin"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/mandaputtra/projectsprint-projects2/services/ms-activity-svc/dtos"
	"github.com/mandaputtra/projectsprint-projects2/services/ms-activity-svc/mappers"
	"github.com/mandaputtra/projectsprint-projects2/services/ms-activity-svc/models"
	"github.com/mandaputtra/projectsprint-projects2/services/ms-activity-svc/repositories"
	"gorm.io/gorm"
	// "gorm.io/gorm"
)

type ActivityService struct {
	repo             *repositories.ActivityRepository
	activityTypeRepo *repositories.ActivityTypeRepository
}

func NewActivityService(repo *repositories.ActivityRepository, activityTypeRepo *repositories.ActivityTypeRepository) *ActivityService {
	return &ActivityService{
		repo:             repo,
		activityTypeRepo: activityTypeRepo,
	}
}

func (s *ActivityService) Create(activityReqDTO *dtos.ActivityRequestDTO, ctx *gin.Context) (*dtos.ActivityResponseDTO, error) {

	activityType, err := s.activityTypeRepo.GetOneByName(activityReqDTO.ActivityType)

	if err != nil {
		return nil, err
	}

	userId, _ := ctx.Get("userId")

	newActivityModel := &models.Activity{
		UserID:            userId.(string),
		ActivityTypeID:    activityType.ID,
		ActivityTypeName:  activityType.ActivityType,
		CaloriesBurned:    activityType.Calories * activityReqDTO.DurationInMinutes,
		DurationInMinutes: activityReqDTO.DurationInMinutes,
		DoneAt:            activityReqDTO.DoneAt,
	}

	activity, err := s.repo.Create(newActivityModel)
	if err != nil {
		return nil, err
	}

	activityResponseDTO := &dtos.ActivityResponseDTO{
		ActivityId:        activity.ID,
		ActivityType:      activity.ActivityTypeName,
		DoneAt:            activity.DoneAt,
		DurationInMinutes: activity.DurationInMinutes,
		CaloriesBurned:    activity.CaloriesBurned,
		CreatedAt:         activity.CreatedAt.String(),
		UpdatedAt:         activity.UpdatedAt.String(),
	}

	return activityResponseDTO, nil
}

func (s *ActivityService) GetAll(limit, offset int, userId string) ([]*dtos.ActivityResponseDTO, error) {
	activities, err := s.repo.GetAll(limit, offset, userId)
	if err != nil {
		return nil, err
	}

	// Map models to response DTOs
	var activityDTOs []*dtos.ActivityResponseDTO
	for _, activity := range activities {

		activityDTOs = append(activityDTOs, mappers.MapActivityModelToResponse(activity))
	}

	return activityDTOs, nil
}

func (s *ActivityService) GetOne(id, userId string) (*dtos.ActivityResponseDTO, error) {
	activity, err := s.repo.GetOne(id, userId)
	if err != nil {
		return nil, err
	}

	return mappers.MapActivityModelToResponse(activity), nil
}

func (s *ActivityService) UpdateActivity(id, userId string, activityDTO *dtos.ActivityRequestDTO) (*dtos.ActivityResponseDTO, error) {
	// Ambil data activity berdasarkan ID
	activity, err := s.GetOne(id,userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("activity not found")
		}
		return nil, err // Error lain
	}

	

	// Perbarui data di repository
	updatedData, err := s.repo.UpdateDepartment(updatedModel)
	if err != nil {
		return nil, err // Error saat update
	}

	// Map hasil ke DTO respons
	return mappers.MapActivityModelToResponse(updatedData), nil
}

// func (s *DepartmentService) UpdateMassDepartmentByEmail(oldEmail string, newEmail string) (string, error) {
// 	activities, err := s.repo.GetAllWithoutPaginationById(oldEmail)
// 	if err != nil {
// 		return "", err
// 	}

// 	for _, activity := range activities {
// 		randomString := utils.GenerateRandomString(5)
// 		oldId := activity.ID
// 		activity.ID = newEmail + "-" + randomString
// 		_, err := s.repo.UpdateDepartmentId(oldId, activity)
// 		if err != nil {
// 			return "", fmt.Errorf("failed to update activity with old ID %s: %w", oldId, err)
// 		}
// 	}
// 	return "All activities updated successfully", nil
// }

// func (s *DepartmentService) DeleteById(id string) error {
// 	return s.repo.DeleteById(id)
// }
