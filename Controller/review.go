package controller

import (
	"net/http"
	"strconv"

	config "github.com/Rifq11/Trava-be/Config"
	models "github.com/Rifq11/Trava-be/Models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateReview(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Status:  "error",
			Message: "Unauthorized",
		})
		return
	}

	var req models.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Booking ID and rating are required",
		})
		return
	}

	if req.Rating < 1 || req.Rating > 5 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Rating must be between 1 and 5",
		})
		return
	}

	userIdInt := userID.(int)

	var booking models.Booking
	if err := config.DB.First(&booking, req.BookingID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Status:  "error",
				Message: "Booking not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to create review",
		})
		return
	}

	if booking.UserID != userIdInt {
		c.JSON(http.StatusForbidden, models.ErrorResponse{
			Status:  "error",
			Message: "You can only review your own bookings",
		})
		return
	}

	review := models.Review{
		BookingID:  req.BookingID,
		UserID:     userIdInt,
		Rating:     req.Rating,
		ReviewText: req.ReviewText,
	}

	if err := config.DB.Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to create review",
		})
		return
	}

	c.JSON(http.StatusCreated, models.SuccessResponse{
		Status:  "success",
		Message: "Review created successfully",
		Data:    map[string]interface{}{"review_id": review.ID},
	})
}

func GetDestinationReviews(c *gin.Context) {
	destinationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Invalid destination ID",
		})
		return
	}

	var reviews []models.ReviewResponse
	if err := config.DB.
		Table("reviews").
		Select("reviews.id, reviews.booking_id, reviews.user_id, users.full_name as user_name, reviews.rating, reviews.review_text").
		Joins("INNER JOIN bookings ON reviews.booking_id = bookings.id").
		Joins("INNER JOIN users ON reviews.user_id = users.id").
		Where("bookings.destination_id = ?", destinationID).
		Order("reviews.id DESC").
		Scan(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to get reviews",
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status: "success",
		Data:   reviews,
	})
}

