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
		// Admin routes
		admin := protected.Group("/")
		admin.Use(middleware.AdminOnly())
		{
			admin.POST("/vehicles", controllers.CreateVehicle)
			admin.PUT("/vehicles/:id", controllers.UpdateVehicle)
			admin.DELETE("/vehicles/:id", controllers.DeleteVehicle)

			// Admin payment routes
			admin.PUT("/payments/:id", controllers.UpdatePayment)
			admin.GET("/payments", controllers.GetAllPayments)
		}

		// Customer routes
		protected.POST("/bookings", controllers.CreateBooking)
		protected.GET("/bookings", controllers.GetUserBookings)
		protected.PUT("/bookings/:id", controllers.UpdateBooking)
		protected.DELETE("/bookings/:id", controllers.CancelBooking)

		// Payment routes
		protected.POST("/bookings/:id/payments", controllers.CreatePayment)
		protected.GET("/bookings/:id/payments", controllers.GetBookingPayments)
	}
}
