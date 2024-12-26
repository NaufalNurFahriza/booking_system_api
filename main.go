package main

import (
	"booking_api/config"
	"booking_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
