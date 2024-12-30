package controllers

import (
	"booking_api/config"
	"booking_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateMovie handles creating a new movie (Admin only)
func CreateMovie(c *gin.Context) {
	// Ensure the user is an admin
	if role, _ := c.Get("role"); role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}

	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the new movie to the database
	if err := config.DB.Create(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create movie"})
		return
	}

	c.JSON(http.StatusCreated, movie)
}

// UpdateMovie handles updating an existing movie (Admin only)
func UpdateMovie(c *gin.Context) {
	// Ensure the user is an admin
	if role, _ := c.Get("role"); role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}

	var movie models.Movie
	if err := config.DB.First(&movie, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the updated movie to the database
	if err := config.DB.Save(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update movie"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

// GetAllMovies retrieves all movies from the database (Public access)
func GetAllMovies(c *gin.Context) {
	var movies []models.Movie

	// Fetch all movies from the database
	if err := config.DB.Find(&movies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
		return
	}

	// Return the list of movies
	c.JSON(http.StatusOK, movies)
}

// DeleteMovie handles deleting an existing movie (Admin only)
func DeleteMovie(c *gin.Context) {
	// Ensure the user is an admin
	if role, _ := c.Get("role"); role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}

	var movie models.Movie
	if err := config.DB.First(&movie, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	// Delete the movie from the database
	if err := config.DB.Delete(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete movie"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}
