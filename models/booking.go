package models

import (
	"time"
)

type Booking struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	UserID        uint      `json:"user_id" gorm:"not null"`
	ScheduleID    uint      `json:"schedule_id" gorm:"not null"`
	BookingDate   time.Time `json:"booking_date"`
	TotalPrice    float64   `json:"total_price" gorm:"not null;check:total_price > 0"`
	PaymentStatus string    `json:"payment_status" gorm:"default:unpaid;not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`

	User     User     `json:"user" gorm:"foreignKey:UserID"`
	Schedule Schedule `json:"Schedule" gorm:"foreignKey:ScheduleID"`
}
