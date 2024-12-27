package models

import (
	"time"
)

type Vehicle struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"not null"`
	Type         string    `json:"type" gorm:"not null"`
	LicensePlate string    `json:"license_plate" gorm:"unique;not null"`
	PricePerDay  float64   `json:"price_per_day" gorm:"not null;check:price_per_day > 0"`
	Available    bool      `json:"available" gorm:"default:true"`
	Description  string    `json:"description"`
	ImageURL     string    `json:"image_url"`
	CreatedAt    time.Time `json:"created_at"`
}
