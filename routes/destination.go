package routes

import (
	controller "github.com/Rifq11/Trava-be/controller"
	middleware "github.com/Rifq11/Trava-be/middleware"
	"github.com/gin-gonic/gin"
)

func DestinationRoutes(router *gin.RouterGroup) {
	destinations := router.Group("/destinations")
	{
		destinations.GET("", controller.GetDestinations)
		destinations.GET("/:id", controller.GetDestinationById)
		destinations.POST("", middleware.RequireAuth(), controller.CreateDestination)
		destinations.PUT("/:id", controller.UpdateDestination)
		destinations.DELETE("/:id", controller.DeleteDestination)
	}
}
