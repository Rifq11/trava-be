package controller

import (
	"net/http"

	config "github.com/Rifq11/Trava-be/Config"
	models "github.com/Rifq11/Trava-be/Models"
	"github.com/gin-gonic/gin"
)

func LogActivity(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Status:  "error",
			Message: "Unauthorized",
		})
		return
	}

	var req models.LogActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Activity type is required",
		})
		return
	}

	userIdInt := userID.(int)

	activityLog := models.ActivityLog{
		UserID:        userIdInt,
		DestinationID: req.DestinationID,
		ActivityType:  req.ActivityType,
	}

	if err := config.DB.Table("user_activity_log").Create(&activityLog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to log activity: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.SuccessResponse{
		Status:  "success",
		Message: "Activity logged successfully",
		Data:    map[string]interface{}{"activity_id": activityLog.ID},
	})
}

