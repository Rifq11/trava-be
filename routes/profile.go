package routes

import (
	controller "github.com/Rifq11/Trava-be/controller"
	middleware "github.com/Rifq11/Trava-be/middleware"
	"github.com/gin-gonic/gin"
)

func ProfileRoutes(router *gin.RouterGroup) {
	profile := router.Group("/profile")
	profile.Use(middleware.RequireAuth())
	{
		profile.GET("", controller.GetProfile)
		profile.POST("/complete", controller.CompleteProfile)
	}
}

