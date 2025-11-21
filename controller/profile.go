package controller

import (
	"net/http"
	"strconv"

	config "github.com/Rifq11/Trava-be/config"
	helper "github.com/Rifq11/Trava-be/helper"
	models "github.com/Rifq11/Trava-be/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	userIdInt := userID.(int)

	var user models.User
	result := config.DB.First(&user, userIdInt)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	var userProfile models.UserProfile
	var adminProfile models.AdminProfile
	var profile interface{}

	if result := config.DB.Where("user_id = ?", userIdInt).First(&userProfile); result.Error == nil {
		profile = userProfile
	} else if result := config.DB.Where("user_id = ?", userIdInt).First(&adminProfile); result.Error == nil {
		profile = adminProfile
	}

	response := models.ProfileResponse{
		User:    user,
		Profile: profile,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
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

	userIdInt := userID.(int)

	var userPhoto string
	if uploadedFile, exists := c.Get("uploaded_file"); exists {
		if filename, ok := uploadedFile.(string); ok {
			// get url
			userPhoto = helper.GetFileUrl(filename)
		}
	}
	if userPhoto == "" {
		userPhoto = c.PostForm("user_photo")
		if userPhoto == "" {
			userPhoto = c.PostForm("userPhoto")
		}
	}

	phone := c.PostForm("phone")
	address := c.PostForm("address")
	birthDate := c.PostForm("birth_date")
	if birthDate == "" {
		birthDate = c.PostForm("birthDate")
	}
	isAdminStr := c.PostForm("is_admin")
	if isAdminStr == "" {
		isAdminStr = c.PostForm("isAdmin")
	}

	var req models.CompleteProfileRequest
	if phone != "" {
		req.Phone = &phone
	}
	if address != "" {
		req.Address = &address
	}
	if birthDate != "" {
		req.BirthDate = &birthDate
	}
	if userPhoto != "" {
		req.UserPhoto = &userPhoto
	}
	if isAdminStr != "" {
		isAdmin, err := strconv.ParseBool(isAdminStr)
		if err == nil {
			req.IsAdmin = &isAdmin
		}
	}

	var user models.User
	result := config.DB.First(&user, userIdInt)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	isAdmin := false
	if req.IsAdmin != nil {
		isAdmin = *req.IsAdmin
	}

	if isAdmin {
		var adminProfile models.AdminProfile
		result := config.DB.Where("user_id = ?", userIdInt).First(&adminProfile)
		if result.Error == nil {
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
			result = config.DB.Save(&adminProfile)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": result.Error.Error(),
				})
				return
			}
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
			result = config.DB.Create(&newProfile)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": result.Error.Error(),
				})
				return
			}
		}
	} else {
		var userProfile models.UserProfile
		result := config.DB.Where("user_id = ?", userIdInt).First(&userProfile)
		if result.Error == nil {
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
			result = config.DB.Save(&userProfile)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": result.Error.Error(),
				})
				return
			}
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
			result = config.DB.Create(&newProfile)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": result.Error.Error(),
				})
				return
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile completed successfully",
	})
}
