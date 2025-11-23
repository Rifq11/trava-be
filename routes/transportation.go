package routes

import (
	controller "github.com/Rifq11/Trava-be/controller"
	"github.com/gin-gonic/gin"
)

func TransportationRoutes(router *gin.RouterGroup) {
	transportations := router.Group("/transportations")
	{
		transportations.GET("/destination/:id", controller.GetTransportationsByDestination)
		transportations.GET("/all", controller.GetAllAccommodations)
		transportations.POST("", controller.CreateTransportation)
		transportations.PUT("/:id", controller.UpdateTransportation)
		transportations.DELETE("/:id", controller.DeleteTransportation)
		transportations.GET("/transport-types", controller.GetTransportTypes)
		transportations.POST("/transport-types", controller.CreateTransportType)
		transportations.PUT("/transport-types/:id", controller.UpdateTransportType)
		transportations.DELETE("/transport-types/:id", controller.DeleteTransportType)
	}
}