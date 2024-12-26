package repositories

import (
	"booking_api/models"

	"gorm.io/gorm"
)

type BookingRepository interface {
	CreateBooking(booking *models.Booking) (*models.Booking, error)
	GetBookingByID(id uint) (*models.Booking, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) CreateBooking(booking *models.Booking) (*models.Booking, error) {
	if err := r.db.Create(booking).Error; err != nil {
		return nil, err
	}
	return booking, nil
}

func (r *bookingRepository) GetBookingByID(id uint) (*models.Booking, error) {
	var booking models.Booking
	if err := r.db.First(&booking, id).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}
