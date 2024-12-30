package routes

import (
	"booking_api/controllers"
	"booking_api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")

	//Auth routes
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	// Public routes
	api.GET("/movies", controllers.GetAllMovies)
	api.GET("/schedules", controllers.GetAllSchedules)

	// Protected routes
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// Customer profile management
		api.PUT("/customers/:id", controllers.UpdateCustomer)
		api.DELETE("/customers/:id", controllers.DeleteCustomer)

		// Booking routes for all authenticated users
		api.POST("/bookings", controllers.CreateBooking)
		api.PUT("/bookings/:id", controllers.UpdateBooking)
		api.DELETE("/bookings/:id", controllers.DeleteBooking)

		// Admin routes
		admin := protected.Group("/")
		admin.Use(middleware.AdminOnly())
		{
			// Customer management
			admin.GET("/customers", controllers.GetAllCustomers)

			// Movie management
			admin.POST("/movies", controllers.CreateMovie)
			admin.PUT("/movies/:id", controllers.UpdateMovie)
			admin.DELETE("/movies/:id", controllers.DeleteMovie)

			// Schedule management
			admin.POST("/schedules", controllers.CreateSchedule)
			admin.PUT("/schedules/:id", controllers.UpdateSchedule)
			admin.DELETE("/schedules/:id", controllers.DeleteSchedule)

			// Booking management
			admin.GET("/bookings", controllers.GetAllBookings)
		}

	}
}
