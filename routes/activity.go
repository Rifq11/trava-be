package routes

import (
	controller "github.com/Rifq11/Trava-be/Controller"
	middleware "github.com/Rifq11/Trava-be/Middleware"
	"github.com/gin-gonic/gin"
)

func ActivityRoutes(router *gin.RouterGroup) {
	activity := router.Group("/activity")
	activity.Use(middleware.RequireAuth())
	{
		activity.POST("", controller.LogActivity)
	}
}

