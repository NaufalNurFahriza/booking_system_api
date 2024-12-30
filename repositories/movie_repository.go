package repositories

import (
	"booking_api/models"

	"gorm.io/gorm"
)

type MovieRepository interface {
	CreateMovie(movie *models.Movie) (*models.Movie, error)
	UpdateMovie(movie *models.Movie) (*models.Movie, error)
	GetAllMovies() ([]models.Movie, error)
	DeleteMovie(id uint) error
}

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{db: db}
}

// CreateMovie creates a new movie record in the database
func (r *movieRepository) CreateMovie(movie *models.Movie) (*models.Movie, error) {
	if err := r.db.Create(movie).Error; err != nil {
		return nil, err
	}
	return movie, nil
}

// UpdateMovie updates an existing movie record
func (r *movieRepository) UpdateMovie(movie *models.Movie) (*models.Movie, error) {
	if err := r.db.Save(movie).Error; err != nil {
		return nil, err
	}
	return movie, nil
}

// GetAllMovies retrieves all movies from the database
func (r *movieRepository) GetAllMovies() ([]models.Movie, error) {
	var movies []models.Movie
	if err := r.db.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

// DeleteMovie deletes a movie by its ID
func (r *movieRepository) DeleteMovie(id uint) error {
	if err := r.db.Delete(&models.Movie{}, id).Error; err != nil {
		return err
	}
	return nil
}
