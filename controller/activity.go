package controller

import (
	"net/http"

	config "github.com/Rifq11/Trava-be/config"
	models "github.com/Rifq11/Trava-be/models"
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

	result := config.DB.Table("user_activity_log").Create(&activityLog)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Activity logged successfully",
		Data:    activityLog,
	})
}

