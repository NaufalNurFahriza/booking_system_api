package main

import (
	"booking_api/config"
	"booking_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection and get the DB instance
	db := config.InitDB()

	// Create a new Gin router
	r := gin.Default()

	// Setup routes with the database connection
	routes.SetupRoutes(r, db)

	// Run the server on port 8080
	r.Run(":8080")
}
