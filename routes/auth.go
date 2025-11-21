package routes

import (
	controller "github.com/Rifq11/Trava-be/controller"
	middleware "github.com/Rifq11/Trava-be/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
		auth.PUT("/profile", middleware.RequireAuth(), controller.UpdateProfile)
	}
}

