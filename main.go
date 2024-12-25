package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"quiz-sanbercode/controllers"
	"quiz-sanbercode/database"
	"quiz-sanbercode/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	database.DBMigrate(db)
	database.DbConnection = db

	router := gin.Default()

	// Public routes
	router.POST("/api/users/register", controllers.RegisterUser)
	router.POST("/api/users/login", controllers.Login)

	// Protected routes
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// Categories routes
		api.GET("/categories", controllers.GetAllCategories)
		api.POST("/categories", controllers.CreateCategory)
		api.GET("/categories/:id", controllers.GetCategory)
		api.PUT("/categories/:id", controllers.UpdateCategory)
		api.DELETE("/categories/:id", controllers.DeleteCategory)
		api.GET("/categories/:id/books", controllers.GetBooksByCategory)

		// Books routes
		api.GET("/books", controllers.GetAllBooks)
		api.POST("/books", controllers.CreateBook)
		api.GET("/books/:id", controllers.GetBook)
		api.PUT("/books/:id", controllers.UpdateBook)
		api.DELETE("/books/:id", controllers.DeleteBook)
	}

	router.Run(":" + os.Getenv("PORT"))
}
