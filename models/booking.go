package models

import (
	"time"
)

type Booking struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	VehicleID   uint      `json:"vehicle_id" gorm:"not null"`
	BookingDate time.Time `json:"booking_date"`
	StartDate   time.Time `json:"start_date" gorm:"not null"`
	EndDate     time.Time `json:"end_date" gorm:"not null"`
	TotalPrice  float64   `json:"total_price" gorm:"not null;check:total_price > 0"`
	Status      string    `json:"status" gorm:"default:pending"`
	Notes       string    `json:"notes"`
	User        User      `json:"user" gorm:"foreignKey:UserID"`
	Vehicle     Vehicle   `json:"vehicle" gorm:"foreignKey:VehicleID"`
	Payments    []Payment `json:"payments" gorm:"foreignKey:BookingID"`
}
