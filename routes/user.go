package routes

import (
	controller "github.com/Rifq11/Trava-be/Controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	users := router.Group("/users")
	{
		users.GET("", controller.GetAllUsers)
		users.GET("/:id", controller.GetUserById)
		users.POST("", controller.CreateUser)
		users.PUT("/:id", controller.UpdateUser)
		users.DELETE("/:id", controller.DeleteUser)
	}
}
