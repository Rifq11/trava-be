package controller

import (
	"net/http"
	"strconv"

	config "github.com/Rifq11/Trava-be/Config"
	models "github.com/Rifq11/Trava-be/Models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDestinations(c *gin.Context) {
	categoryIDStr := c.Query("category_id")
	var destinations []models.DestinationResponse

	query := config.DB.
		Table("destinations").
		Select("destinations.id, destinations.name, destinations.description, destinations.location, destinations.price_per_person, destinations.image, destinations.category_id, destinations.created_by, destination_categories.name as category_name").
		Joins("INNER JOIN destination_categories ON destinations.category_id = destination_categories.id").
		Order("destinations.id DESC")

	if categoryIDStr != "" {
		categoryID, err := strconv.Atoi(categoryIDStr)
		if err == nil {
			query = query.Where("destinations.category_id = ?", categoryID)
		}
	}

	if err := query.Scan(&destinations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to get destinations",
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status: "success",
		Data:   destinations,
	})
}

func GetDestinationById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Invalid destination ID",
		})
		return
	}

	var destination models.DestinationDetailResponse
	if err := config.DB.
		Table("destinations").
		Select("destinations.id, destinations.name, destinations.description, destinations.location, destinations.price_per_person, destinations.image, destinations.category_id, destinations.created_by, destination_categories.name as category_name, users.full_name as creator_name").
		Joins("INNER JOIN destination_categories ON destinations.category_id = destination_categories.id").
		Joins("INNER JOIN users ON destinations.created_by = users.id").
		Where("destinations.id = ?", id).
		First(&destination).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Status:  "error",
				Message: "Destination not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to get destination",
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status: "success",
		Data:   destination,
	})
}

func CreateDestination(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Status:  "error",
			Message: "Unauthorized",
		})
		return
	}

	var req models.CreateDestinationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Category ID, name, location, and price per person are required",
		})
		return
	}

	userIdInt := userID.(int)
	destination := models.Destination{
		CategoryID:     req.CategoryID,
		CreatedBy:      userIdInt,
		Name:           req.Name,
		Description:    req.Description,
		Location:       req.Location,
		PricePerPerson: req.PricePerPerson,
		Image:          req.Image,
	}

	if err := config.DB.Create(&destination).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to create destination",
		})
		return
	}

	c.JSON(http.StatusCreated, models.SuccessResponse{
		Status:  "success",
		Message: "Destination created successfully",
		Data:    map[string]interface{}{"id": destination.ID},
	})
}

func UpdateDestination(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Invalid destination ID",
		})
		return
	}

	var req models.UpdateDestinationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
		return
	}

	var destination models.Destination
	if err := config.DB.First(&destination, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Status:  "error",
				Message: "Destination not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to update destination",
		})
		return
	}

	if req.CategoryID != nil {
		destination.CategoryID = *req.CategoryID
	}
	if req.Name != nil {
		destination.Name = *req.Name
	}
	if req.Description != nil {
		destination.Description = *req.Description
	}
	if req.Location != nil {
		destination.Location = *req.Location
	}
	if req.PricePerPerson != nil {
		destination.PricePerPerson = *req.PricePerPerson
	}
	if req.Image != nil {
		destination.Image = *req.Image
	}

	if err := config.DB.Save(&destination).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to update destination",
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Destination updated successfully",
	})
}

func DeleteDestination(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Invalid destination ID",
		})
		return
	}

	var destination models.Destination
	if err := config.DB.First(&destination, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Status:  "error",
				Message: "Destination not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete destination",
		})
		return
	}

	if err := config.DB.Delete(&destination).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete destination",
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Destination deleted successfully",
	})
}
