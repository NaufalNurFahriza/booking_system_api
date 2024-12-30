package repositories

import (
	"booking_api/models"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	CreatePayment(payment *models.Payment) (*models.Payment, error)
	GetPaymentByID(id uint) (*models.Payment, error)
	GetPaymentsByBookingID(bookingID uint) ([]models.Payment, error)
	UpdatePayment(payment *models.Payment) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) CreatePayment(payment *models.Payment) (*models.Payment, error) {
	if err := r.db.Create(payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *paymentRepository) GetPaymentByID(id uint) (*models.Payment, error) {
	var payment models.Payment
	if err := r.db.First(&payment, id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) GetPaymentsByBookingID(bookingID uint) ([]models.Payment, error) {
	var payments []models.Payment
	if err := r.db.Where("booking_id = ?", bookingID).Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *paymentRepository) UpdatePayment(payment *models.Payment) error {
	return r.db.Save(payment).Error
}
