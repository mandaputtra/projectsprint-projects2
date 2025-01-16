package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mandaputtra/projectsprint-projects2/services/ms-activity-svc/services"
)

type ActivityController struct {
	service *services.ActivityService
}

func NewActivityController(service *services.ActivityService) *ActivityController {
	return &ActivityController{
		service: service,
	}
}

// Handler untuk membuat department baru
// func (c *ActivityController) Create(ctx *gin.Context) {
// 	var activityDTO dtos.ActivityRequestDTO

// 	// Bind JSON request body ke struct department
// 	if err := ctx.ShouldBindJSON(&activityDTO); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Invalid request body",
// 		})
// 		return
// 	}

// 	// Panggil service untuk membuat department
// 	createdDepartment, err := c.service.Create(&activityDTO, ctx)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Failed to create department",
// 		})
// 		return
// 	}

// 	// Kirim response dengan data department yang berhasil dibuat
// 	ctx.JSON(http.StatusCreated, createdDepartment)
// }

func (c *ActivityController) GetAllDepartments(ctx *gin.Context) {
	limitStr := ctx.DefaultQuery("limit", "10")
	offsetStr := ctx.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		// Jika terjadi error, beri nilai default
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		// Jika terjadi error, beri nilai default
		offset = 0
	}

	activities, err := c.service.GetAll(limit, offset)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch activities"})
		return
	}

	if len(activities) <= 0 {
		ctx.JSON(http.StatusNoContent, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, activities)
}

// func (c *ActivityController) GetOneDepartment(ctx *gin.Context) {
// 	checkIdWithCredential(ctx)
// 	if ctx.IsAborted() {
// 		return
// 	}

// 	department, err := c.service.GetOne(id)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "ID is not found"})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, department)
// }

// func (c *ActivityController) UpdateOneDepartment(ctx *gin.Context) {
// 	checkIdWithCredential(ctx)
// 	if ctx.IsAborted() {
// 		return
// 	}

// 	// Bind input dari request body ke DTO
// 	var activityDTO dtos.ActivityRequestDTO
// 	if err := ctx.ShouldBindJSON(&activityDTO); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	// Panggil service untuk update
// 	updatedDepartment, err := c.service.UpdateOneDepartment(id, &activityDTO)
// 	if err != nil {
// 		if err.Error() == "department not found" {
// 			ctx.JSON(http.StatusNotFound, gin.H{"error": "Department with the given ID not found"})
// 			return
// 		}
// 		// Error lain saat update
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update department"})
// 		return
// 	}

// 	// Berhasil update
// 	ctx.JSON(http.StatusOK, updatedDepartment)
// }

// func (c *ActivityController) UpdateDepartmentId(ctx *gin.Context) {
// 	var body map[string]interface{}

// 	// Parse JSON ke map
// 	if err := ctx.ShouldBindJSON(&body); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
// 		return
// 	}

// 	_, err := c.service.UpdateMassDepartmentByEmail(body["oldEmail"].(string), body["newEmail"].(string))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update department"})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, body["oldEmail"])
// }

// func (c *ActivityController) DeleteOneDepartment(ctx *gin.Context) {
// 	checkIdWithCredential(ctx)
// 	if ctx.IsAborted() {
// 		return
// 	}

// 	department, err := c.service.GetOne(id)
// 	if err != nil {
// 		ctx.JSON(http.StatusNotFound, gin.H{"error": "Department with the given ID not found"})
// 	}

// 	err = c.service.DeleteById(department.ID)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "delete is not successful"})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{"status": "delete is successful"})
// }
