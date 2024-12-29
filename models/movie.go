package models

import (
	"time"
)

type Movie struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Genre       string    `json:"genre" gorm:"not null"`
	Duration    int       `json:"duration" gorm:"not null"`
	Description *string   `json:"description" gorm:"type:text"`
	ReleaseDate string    `json:"release_date" gorm:"not null"`
	Rating      string    `json:"rating" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
