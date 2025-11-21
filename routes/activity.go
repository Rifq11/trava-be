package routes

import (
	controller "github.com/Rifq11/Trava-be/controller"
	middleware "github.com/Rifq11/Trava-be/middleware"
	"github.com/gin-gonic/gin"
)

func ActivityRoutes(router *gin.RouterGroup) {
	activity := router.Group("/activity")
	activity.Use(middleware.RequireAuth())
	{
		activity.POST("", controller.LogActivity)
	}
}

