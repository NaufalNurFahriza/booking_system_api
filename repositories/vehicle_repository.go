package repositories

import (
	"booking_api/models"

	"gorm.io/gorm"
)

type VehicleRepository interface {
	CreateVehicle(vehicle *models.Vehicle) (*models.Vehicle, error)
	GetAllVehicles() ([]models.Vehicle, error)
}

type vehicleRepository struct {
	db *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) VehicleRepository {
	return &vehicleRepository{db: db}
}

func (r *vehicleRepository) CreateVehicle(vehicle *models.Vehicle) (*models.Vehicle, error) {
	if err := r.db.Create(vehicle).Error; err != nil {
		return nil, err
	}
	return vehicle, nil
}

func (r *vehicleRepository) GetAllVehicles() ([]models.Vehicle, error) {
	var vehicles []models.Vehicle
	if err := r.db.Find(&vehicles).Error; err != nil {
		return nil, err
	}
	return vehicles, nil
}
