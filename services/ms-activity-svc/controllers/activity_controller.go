package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mandaputtra/projectsprint-projects2/services/ms-activity-svc/dtos"
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

func (c *ActivityController) Create(ctx *gin.Context) {
	var activityDTO dtos.ActivityRequestDTO

	// Bind JSON request body ke struct department
	if err := ctx.ShouldBindJSON(&activityDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	createdActivity, err := c.service.Create(&activityDTO, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create activity",
		})
		return
	}

	// Kirim response dengan data department yang berhasil dibuat
	ctx.JSON(http.StatusCreated, createdActivity)
}

func (c *ActivityController) GetAllActivities(ctx *gin.Context) {
	limitStr := ctx.DefaultQuery("limit", "10")
	offsetStr := ctx.DefaultQuery("offset", "0")
	userId, _ := ctx.Get("userId")

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

	activities, err := c.service.GetAll(limit, offset, userId.(string))

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

func (c *ActivityController) GetOneActivity(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, _ := ctx.Get("userId")

	activity, err := c.service.GetOne(id, userId.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "ID is not found"})
		return
	}
	ctx.JSON(http.StatusOK, activity)
}

func (c *ActivityController) UpdateActivity(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, _ := ctx.Get("userId")
	// Bind input dari request body ke DTO
	var activityDTO dtos.ActivityRequestDTO
	if err := ctx.ShouldBindJSON(&activityDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Panggil service untuk update
	updatedActivity, err := c.service.UpdateActivity(id, userId.(string), &activityDTO)
	if err != nil {
		if err.Error() == "activity not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Activity with the given ID not found"})
			return
		}
		// Error lain saat update
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Activity"})
		return
	}

	// Berhasil update
	ctx.JSON(http.StatusOK, updatedActivity)
}

func (c *ActivityController) DeleteOneActivity(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, _ := ctx.Get("userId")

	err := c.service.DeleteById(id, userId.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "delete is not successful"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "delete is successful"})
}
