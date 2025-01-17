package services

import (
	// "errors"
	// "fmt"

	// "github.com/gin-gonic/gin"
	"github.com/mandaputtra/projectsprint-projects2/services/ms-activity-svc/dtos"
	"github.com/mandaputtra/projectsprint-projects2/services/ms-activity-svc/mappers"
	"github.com/mandaputtra/projectsprint-projects2/services/ms-activity-svc/repositories"
	// "gorm.io/gorm"
)

type ActivityService struct {
	repo *repositories.ActivityRepository
}

func NewActivityService(repo *repositories.ActivityRepository) *ActivityService {
	return &ActivityService{
		repo: repo,
	}
}

// func (s *DepartmentService) Create(activityReqDTO *dtos.ActivityRequestDTO, ctx *gin.Context) (*dtos.ActivityResponseDTO, error) {

// 	newDepartment := mappers.MapRequestToActivityModel(activityReqDTO)

// 	email, ok := ctx.Get("email")
// 	if !ok {
// 		return nil, errors.New("email not found in context")
// 	}

// 	emailStr, ok := email.(string)
// 	if !ok {
// 		return nil, errors.New("email in context is not a valid string")
// 	}

// 	randomString := utils.GenerateRandomString(5)
// 	newDepartment.ID = emailStr + "-" + randomString

// 	activity, err := s.repo.Create(newDepartment)
// 	if err != nil {
// 		return nil, err
// 	}

// 	activityResponseDTO := mappers.MapActivityModelToResponse(activity)

// 	return activityResponseDTO, nil
// }

func (s *ActivityService) GetAll(limit, offset int) ([]*dtos.ActivityResponseDTO, error) {
	activities, err := s.repo.GetAll(limit, offset)
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

// func (s *DepartmentService) GetOne(id string) (*dtos.ActivityResponseDTO, error) {
// 	activity, err := s.repo.GetOne(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return mappers.MapActivityModelToResponse(activity), nil
// }

// func (s *DepartmentService) UpdateOneDepartment(id string, activityDTO *dtos.ActivityRequestDTO) (*dtos.ActivityResponseDTO, error) {
// 	// Ambil data activity berdasarkan ID
// 	activity, err := s.GetOne(id)
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, errors.New("activity not found")
// 		}
// 		return nil, err // Error lain
// 	}

// 	// Gabungkan data baru dengan data lama
// 	updatedModel := mappers.MapRequestToActivityModel(activityDTO)
// 	updatedModel.ID = activity.ID // Tetap gunakan ID lama

// 	// Perbarui data di repository
// 	updatedData, err := s.repo.UpdateDepartment(updatedModel)
// 	if err != nil {
// 		return nil, err // Error saat update
// 	}

// 	// Map hasil ke DTO respons
// 	return mappers.MapActivityModelToResponse(updatedData), nil
// }

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
