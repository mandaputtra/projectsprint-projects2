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

// func (s *DepartmentService) Create(departmentReqDTO *dtos.ActivityRequestDTO, ctx *gin.Context) (*dtos.ActivityResponseDTO, error) {

// 	newDepartment := mappers.MapRequestToActivityModel(departmentReqDTO)

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

// 	department, err := s.repo.Create(newDepartment)
// 	if err != nil {
// 		return nil, err
// 	}

// 	departmentResponseDTO := mappers.MapActivityModelToResponse(department)

// 	return departmentResponseDTO, nil
// }

func (s *ActivityService) GetAll(limit, offset int) ([]*dtos.ActivityResponseDTO, error) {
	departments, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	// Map models to response DTOs
	var departmentDTOs []*dtos.ActivityResponseDTO
	for _, department := range departments {
		departmentDTOs = append(departmentDTOs, mappers.MapActivityModelToResponse(department))
	}

	return departmentDTOs, nil
}

// func (s *DepartmentService) GetOne(id string) (*dtos.ActivityResponseDTO, error) {
// 	department, err := s.repo.GetOne(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return mappers.MapActivityModelToResponse(department), nil
// }

// func (s *DepartmentService) UpdateOneDepartment(id string, departmentDTO *dtos.ActivityRequestDTO) (*dtos.ActivityResponseDTO, error) {
// 	// Ambil data department berdasarkan ID
// 	department, err := s.GetOne(id)
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, errors.New("department not found")
// 		}
// 		return nil, err // Error lain
// 	}

// 	// Gabungkan data baru dengan data lama
// 	updatedModel := mappers.MapRequestToActivityModel(departmentDTO)
// 	updatedModel.ID = department.ID // Tetap gunakan ID lama

// 	// Perbarui data di repository
// 	updatedData, err := s.repo.UpdateDepartment(updatedModel)
// 	if err != nil {
// 		return nil, err // Error saat update
// 	}

// 	// Map hasil ke DTO respons
// 	return mappers.MapActivityModelToResponse(updatedData), nil
// }

// func (s *DepartmentService) UpdateMassDepartmentByEmail(oldEmail string, newEmail string) (string, error) {
// 	departments, err := s.repo.GetAllWithoutPaginationById(oldEmail)
// 	if err != nil {
// 		return "", err
// 	}

// 	for _, department := range departments {
// 		randomString := utils.GenerateRandomString(5)
// 		oldId := department.ID
// 		department.ID = newEmail + "-" + randomString
// 		_, err := s.repo.UpdateDepartmentId(oldId, department)
// 		if err != nil {
// 			return "", fmt.Errorf("failed to update department with old ID %s: %w", oldId, err)
// 		}
// 	}
// 	return "All departments updated successfully", nil
// }

// func (s *DepartmentService) DeleteById(id string) error {
// 	return s.repo.DeleteById(id)
// }
