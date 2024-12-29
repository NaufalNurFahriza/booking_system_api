package models

import (
	"time"
)

type Schedule struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	MovieID     uint      `json:"movie_id" gorm:"not null"`
	ShowTime    time.Time `json:"show_time" gorm:"not null"`
	TicketPrice float64   `json:"ticket_price"  gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`

	Movie Movie `json:"movie" gorm:"foreignKey:MovieID"`
}
