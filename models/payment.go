package models

import (
	"time"
)

type Payment struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	BookingID     uint      `json:"booking_id" gorm:"not null"`
	Amount        float64   `json:"amount" gorm:"not null;check:amount > 0"`
	PaymentDate   time.Time `json:"payment_date"`
	PaymentMethod string    `json:"payment_method" gorm:"not null"`        // cash, credit_card, transfer
	PaymentStatus string    `json:"payment_status" gorm:"default:pending"` // pending, completed, failed, refunded
	TransactionID string    `json:"transaction_id"`
	Notes         string    `json:"notes"`
	Booking       Booking   `json:"booking" gorm:"foreignKey:BookingID"`
}
