package routes

import (
	controller "github.com/Rifq11/Trava-be/controller"
	middleware "github.com/Rifq11/Trava-be/middleware"
	"github.com/gin-gonic/gin"
)

func BookingRoutes(router *gin.RouterGroup) {
	bookings := router.Group("/bookings")
	bookings.Use(middleware.RequireAuth())
	{
		bookings.POST("", controller.CreateBooking)
		bookings.GET("/my", controller.GetMyBookings)
	}
}

