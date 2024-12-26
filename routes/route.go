package routes

import (
	"booking_api/controllers"
	"booking_api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// Public routes
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)
	api.GET("/vehicles", controllers.GetVehicles)

	// Protected routes
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/vehicles", controllers.CreateVehicle)
		protected.POST("/bookings", controllers.CreateBooking)
	}
}
