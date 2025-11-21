package controller

import (
	"net/http"

	config "github.com/Rifq11/Trava-be/Config"
	models "github.com/Rifq11/Trava-be/Models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Status:  "error",
			Message: "Unauthorized",
		})
		return
	}

	userIdInt := userID.(int)

	var user models.User
	if err := config.DB.First(&user, userIdInt).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Status:  "error",
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to get profile",
		})
		return
	}

	var userProfile models.UserProfile
	var adminProfile models.AdminProfile
	var profile interface{}

	if err := config.DB.Where("user_id = ?", userIdInt).First(&userProfile).Error; err == nil {
		profile = userProfile
	} else if err := config.DB.Where("user_id = ?", userIdInt).First(&adminProfile).Error; err == nil {
		profile = adminProfile
	}

	response := models.ProfileResponse{
		User:    user,
		Profile: profile,
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status: "success",
		Data:   response,
	})
}

func CompleteProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Status:  "error",
			Message: "Unauthorized",
		})
		return
	}

	var req models.CompleteProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
		return
	}

	userIdInt := userID.(int)

	var user models.User
	if err := config.DB.First(&user, userIdInt).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Status:  "error",
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to complete profile",
		})
		return
	}

	isAdmin := false
	if req.IsAdmin != nil {
		isAdmin = *req.IsAdmin
	}

	if isAdmin {
		var adminProfile models.AdminProfile
		if err := config.DB.Where("user_id = ?", userIdInt).First(&adminProfile).Error; err == nil {
			if req.Phone != nil {
				adminProfile.Phone = *req.Phone
			}
			if req.Address != nil {
				adminProfile.Address = *req.Address
			}
			if req.BirthDate != nil {
				adminProfile.BirthDate = *req.BirthDate
			}
			if req.UserPhoto != nil {
				adminProfile.UserPhoto = *req.UserPhoto
			}
			adminProfile.IsCompleted = true
			config.DB.Save(&adminProfile)
		} else {
			newProfile := models.AdminProfile{
				UserID:      userIdInt,
				Phone:       "",
				Address:     "",
				BirthDate:   "",
				UserPhoto:   "",
				IsCompleted: true,
			}
			if req.Phone != nil {
				newProfile.Phone = *req.Phone
			}
			if req.Address != nil {
				newProfile.Address = *req.Address
			}
			if req.BirthDate != nil {
				newProfile.BirthDate = *req.BirthDate
			}
			if req.UserPhoto != nil {
				newProfile.UserPhoto = *req.UserPhoto
			}
			config.DB.Create(&newProfile)
		}
	} else {
		var userProfile models.UserProfile
		if err := config.DB.Where("user_id = ?", userIdInt).First(&userProfile).Error; err == nil {
			if req.Phone != nil {
				userProfile.Phone = *req.Phone
			}
			if req.Address != nil {
				userProfile.Address = *req.Address
			}
			if req.BirthDate != nil {
				userProfile.BirthDate = *req.BirthDate
			}
			if req.UserPhoto != nil {
				userProfile.UserPhoto = *req.UserPhoto
			}
			userProfile.IsCompleted = true
			config.DB.Save(&userProfile)
		} else {
			newProfile := models.UserProfile{
				UserID:      userIdInt,
				Phone:       "",
				Address:     "",
				BirthDate:   "",
				UserPhoto:   "",
				IsCompleted: true,
			}
			if req.Phone != nil {
				newProfile.Phone = *req.Phone
			}
			if req.Address != nil {
				newProfile.Address = *req.Address
			}
			if req.BirthDate != nil {
				newProfile.BirthDate = *req.BirthDate
			}
			if req.UserPhoto != nil {
				newProfile.UserPhoto = *req.UserPhoto
			}
			config.DB.Create(&newProfile)
		}
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Profile completed successfully",
	})
}

