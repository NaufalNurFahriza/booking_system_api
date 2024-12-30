package config

import (
	"booking_api/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB { // Mengembalikan *gorm.DB
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// dsn := os.Getenv("DATABASE_URL")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGPORT"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// Auto Migrate
	err = DB.AutoMigrate(&models.User{}, &models.Movie{}, &models.Booking{}, &models.Schedule{})
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}

	return DB // Mengembalikan objek DB
}
