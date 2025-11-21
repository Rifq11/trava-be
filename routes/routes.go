package routes

import "../Routes/github.com/gin-gonic/gin"

func SetupRoutes(app *gin.Engine) {
	api := app.Group("/api")
	{
		AuthRoutes(api)
		ProfileRoutes(api)
		UserRoutes(api)
		DestinationRoutes(api)
		BookingRoutes(api)
		PaymentRoutes(api)
		ReviewRoutes(api)
		ActivityRoutes(api)
	}
}
