package routes

import (
	controller "github.com/Rifq11/Trava-be/Controller"
	middleware "github.com/Rifq11/Trava-be/Middleware"
	"github.com/gin-gonic/gin"
)

func ReviewRoutes(router *gin.RouterGroup) {
	reviews := router.Group("/reviews")
	{
		reviews.POST("", middleware.RequireAuth(), controller.CreateReview)
		reviews.GET("/destination/:id", controller.GetDestinationReviews)
	}
}

