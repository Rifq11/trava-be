package controller

import (
	"net/http"
	"strconv"
	"strings"

	config "github.com/Rifq11/Trava-be/config"
	models "github.com/Rifq11/Trava-be/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetAllUsers(c *gin.Context) {
	var userResponses []models.UserResponse

	if err := config.DB.
		Table("users").
		Select("users.id, users.full_name, users.email, users.role_id, roles.name as role_name").
		Joins("INNER JOIN roles ON users.role_id = roles.id").
		Order("users.id").
		Scan(&userResponses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to get users",
		})
		return
	}

	if userResponses == nil {
		userResponses = []models.UserResponse{}
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status: "success",
		Data:   userResponses,
	})
}

func GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to get user",
		})
		return
	}

	var userResponse models.UserResponse
	if err := config.DB.
		Table("users").
		Select("users.id, users.full_name, users.email, users.role_id, roles.name as role_name").
		Joins("INNER JOIN roles ON users.role_id = roles.id").
		Where("users.id = ?", id).
		First(&userResponse).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Status:  "error",
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to get user",
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status: "success",
		Data:   userResponse,
	})
}

func CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Full name, email, and password are required",
		})
		return
	}

	var existingUser models.User
	if err := config.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, models.ErrorResponse{
			Status:  "error",
			Message: "User with this email already exists",
		})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to create user",
		})
		return
	}

	roleID := int(2)
	if req.RoleID != nil {
		roleID = *req.RoleID
	}

	user := models.User{
		FullName: req.FullName,
		Email:    req.Email,
		Password: string(hashedPassword),
		RoleID:   roleID,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to create user",
		})
		return
	}

	// Get role name for response
	var role models.Role
	config.DB.First(&role, user.RoleID)

	userResponse := models.UserResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		RoleID:   user.RoleID,
		RoleName: role.Name,
	}

	c.JSON(http.StatusCreated, models.SuccessResponse{
		Status:  "success",
		Message: "User created successfully",
		Data:    userResponse,
	})
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to update user",
		})
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to update user",
		})
		return
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Status:  "error",
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to update user",
		})
		return
	}

	if req.Email != nil && *req.Email != user.Email {
		var existingUser models.User
		if err := config.DB.Where("email = ?", *req.Email).First(&existingUser).Error; err == nil {
			c.JSON(http.StatusConflict, models.ErrorResponse{
				Status:  "error",
				Message: "User with this email already exists",
			})
			return
		}
		user.Email = *req.Email
	}

	if req.FullName != nil {
		user.FullName = *req.FullName
	}
	if req.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Status:  "error",
				Message: "Failed to update user",
			})
			return
		}
		user.Password = string(hashedPassword)
	}
	if req.RoleID != nil {
		user.RoleID = *req.RoleID
	}

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to update user",
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "User updated successfully",
	})
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete user",
		})
		return
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Status:  "error",
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete user",
		})
		return
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		errStr := err.Error()
		if strings.Contains(errStr, "1451") || strings.Contains(errStr, "foreign key constraint") {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Status:  "error",
				Message: "Cannot delete user: user is referenced by other records",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete user",
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "User deleted successfully",
	})
}
