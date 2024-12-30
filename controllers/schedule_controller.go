package controllers

import (
	"booking_api/config"
	"booking_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateSchedule handles creating a new schedule (Admin only)
func CreateSchedule(c *gin.Context) {
	if role, _ := c.Get("role"); role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}

	var schedule models.Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the new schedule to the database
	if err := config.DB.Create(&schedule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create schedule"})
		return
	}

	c.JSON(http.StatusCreated, schedule)
}

// UpdateSchedule handles updating an existing schedule (Admin only)
func UpdateSchedule(c *gin.Context) {
	if role, _ := c.Get("role"); role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}

	var schedule models.Schedule
	if err := config.DB.First(&schedule, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the updated schedule to the database
	if err := config.DB.Save(&schedule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update schedule"})
		return
	}

	c.JSON(http.StatusOK, schedule)
}

// GetAllSchedules retrieves all schedules (Public access)
func GetAllSchedules(c *gin.Context) {
	var schedules []models.Schedule

	// Fetch all schedules from the database
	if err := config.DB.Preload("Movie").Find(&schedules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schedules"})
		return
	}

	// Return the list of schedules
	c.JSON(http.StatusOK, schedules)
}

// DeleteSchedule handles deleting an existing schedule (Admin only)
func DeleteSchedule(c *gin.Context) {
	if role, _ := c.Get("role"); role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}

	var schedule models.Schedule
	if err := config.DB.First(&schedule, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	// Delete the schedule from the database
	if err := config.DB.Delete(&schedule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete schedule"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule deleted successfully"})
}
