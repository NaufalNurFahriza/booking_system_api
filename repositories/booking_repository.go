package repositories

import (
	"booking_api/models"
	"errors"

	"gorm.io/gorm"
)

type BookingRepository interface {
	CreateBooking(booking *models.Booking) (*models.Booking, error)
	GetBookingByID(id uint) (*models.Booking, error)
	GetUserBookings(userID uint) ([]models.Booking, error)
	UpdateBooking(booking *models.Booking) (*models.Booking, error)
	DeleteBooking(id uint) error
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

// CreateBooking creates a new booking record in the database
func (r *bookingRepository) CreateBooking(booking *models.Booking) (*models.Booking, error) {
	if err := r.db.Create(booking).Error; err != nil {
		return nil, err
	}
	return booking, nil
}

// GetBookingByID retrieves a booking by its ID
func (r *bookingRepository) GetBookingByID(id uint) (*models.Booking, error) {
	var booking models.Booking
	if err := r.db.First(&booking, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &booking, nil
}

// GetUserBookings retrieves all bookings made by a specific user
func (r *bookingRepository) GetUserBookings(userID uint) ([]models.Booking, error) {
	var bookings []models.Booking
	if err := r.db.Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

// UpdateBooking updates an existing booking record
func (r *bookingRepository) UpdateBooking(booking *models.Booking) (*models.Booking, error) {
	if err := r.db.Save(booking).Error; err != nil {
		return nil, err
	}
	return booking, nil
}

// DeleteBooking deletes a booking by its ID
func (r *bookingRepository) DeleteBooking(id uint) error {
	if err := r.db.Delete(&models.Booking{}, id).Error; err != nil {
		return err
	}
	return nil
}
