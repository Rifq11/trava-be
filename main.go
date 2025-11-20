package main

import (
	config "github.com/Rifq11/Trava-be/Config"
	routes "github.com/Rifq11/Trava-be/Routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
