package repositories

import (
	"booking_api/models"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	CreateSchedule(schedule *models.Schedule) (*models.Schedule, error)
	UpdateSchedule(schedule *models.Schedule) (*models.Schedule, error)
	GetAllSchedules() ([]models.Schedule, error)
	DeleteSchedule(id uint) error
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db: db}
}

// CreateSchedule creates a new schedule record in the database
func (r *scheduleRepository) CreateSchedule(schedule *models.Schedule) (*models.Schedule, error) {
	if err := r.db.Create(schedule).Error; err != nil {
		return nil, err
	}
	return schedule, nil
}

// UpdateSchedule updates an existing schedule record
func (r *scheduleRepository) UpdateSchedule(schedule *models.Schedule) (*models.Schedule, error) {
	if err := r.db.Save(schedule).Error; err != nil {
		return nil, err
	}
	return schedule, nil
}

// GetAllSchedules retrieves all schedules from the database
func (r *scheduleRepository) GetAllSchedules() ([]models.Schedule, error) {
	var schedules []models.Schedule
	if err := r.db.Preload("Movie").Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

// DeleteSchedule deletes a schedule by its ID
func (r *scheduleRepository) DeleteSchedule(id uint) error {
	if err := r.db.Delete(&models.Schedule{}, id).Error; err != nil {
		return err
	}
	return nil
}
