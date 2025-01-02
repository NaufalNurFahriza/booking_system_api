package main

import (
	"booking_api/config"
	"booking_api/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection and get the DB instance
	db := config.InitDB()

	// Create a new Gin router
	r := gin.Default()

	// Setup routes with the database connection
	routes.SetupRoutes(r, db)

	// Get port from environment variable, default to "8080" if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Log and run the server on the specified port
	log.Printf("Server is running on port %s", port)
	r.Run(":" + port)
}
