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

func InitDB() *gorm.DB {
	// Load .env file dari root project
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Cek apakah DATABASE_URL ada dan log isinya
	log.Printf("DATABASE_URL: %s", os.Getenv("DATABASE_URL"))

	// Gunakan DATABASE_URL yang sudah disediakan Railway
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Jika DATABASE_URL kosong, gunakan individual credentials
		dsn = fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s?sslmode=require",
			os.Getenv("PGUSER"),
			os.Getenv("PGPASSWORD"),
			os.Getenv("PGHOST"),
			os.Getenv("PGPORT"),
			os.Getenv("PGDATABASE"),
		)
	}

	// Tampilkan DSN untuk debug (hapus password sebelum logging)
	log.Printf("Connecting to database...")

	// Buka koneksi dengan PostgreSQL
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Verifikasi koneksi
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// Test koneksi
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Printf("Database connected successfully")

	// Auto Migrate
	log.Printf("Running migrations...")
	err = DB.AutoMigrate(&models.User{}, &models.Movie{}, &models.Booking{}, &models.Schedule{})
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}
	log.Printf("Migrations completed successfully")

	return DB
}
