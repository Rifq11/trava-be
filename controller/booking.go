package controller

import (
	"net/http"

	config "github.com/Rifq11/Trava-be/config"
	models "github.com/Rifq11/Trava-be/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateBooking(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Status:  "error",
			Message: "Unauthorized",
		})
		return
	}

	var req models.CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "All booking fields are required",
		})
		return
	}

	userIdInt := userID.(int)

	var destination models.Destination
	if err := config.DB.First(&destination, req.DestinationID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Status:  "error",
				Message: "Destination or transportation not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to create booking",
		})
		return
	}

	var transportation models.Transportation
	if err := config.DB.First(&transportation, req.TransportationID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Status:  "error",
				Message: "Destination or transportation not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to create booking",
		})
		return
	}

	destinationPrice := destination.PricePerPerson * req.PeopleCount
	transportPrice := transportation.Price
	totalPrice := destinationPrice + transportPrice

	booking := models.Booking{
		UserID:           userIdInt,
		DestinationID:    req.DestinationID,
		TransportationID: req.TransportationID,
		PaymentMethodID:  req.PaymentMethodID,
		StatusID:         1, // Pending
		PeopleCount:      req.PeopleCount,
		StartDate:        req.StartDate,
		EndDate:          req.EndDate,
		TransportPrice:   transportPrice,
		DestinationPrice: destinationPrice,
		TotalPrice:       totalPrice,
	}

	if err := config.DB.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to create booking",
		})
		return
	}

	c.JSON(http.StatusCreated, models.SuccessResponse{
		Status:  "success",
		Message: "Booking created successfully",
		Data:    map[string]interface{}{"booking_id": booking.ID, "total_price": booking.TotalPrice},
	})
}

func GetMyBookings(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Status:  "error",
			Message: "Unauthorized",
		})
		return
	}

	userIdInt := userID.(int)

	var bookings []models.BookingResponse
	if err := config.DB.
		Table("bookings").
		Select("bookings.id as booking_id, destinations.name as destination_name, destinations.location, bookings.people_count, bookings.start_date, bookings.end_date, bookings.total_price, booking_status.name as status_name, payment_methods.name as payment_method_name").
		Joins("INNER JOIN destinations ON bookings.destination_id = destinations.id").
		Joins("INNER JOIN booking_status ON bookings.status_id = booking_status.id").
		Joins("INNER JOIN payment_methods ON bookings.payment_method_id = payment_methods.id").
		Where("bookings.user_id = ?", userIdInt).
		Order("bookings.id DESC").
		Scan(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to get bookings",
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status: "success",
		Data:   bookings,
	})
}

