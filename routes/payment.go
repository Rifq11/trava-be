package routes

import (
	controller "github.com/Rifq11/Trava-be/controller"
	"github.com/gin-gonic/gin"
)

func PaymentRoutes(router *gin.RouterGroup) {
	payments := router.Group("/payments")
	{
		payments.POST("", controller.InitiatePayment)
		payments.PUT("/:id", controller.UpdatePayment)
	}
}

