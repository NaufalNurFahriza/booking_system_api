package controllers

import (
	"booking_api/config"
	"booking_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBooking(c *gin.Context) {
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	booking.UserID = userID.(uint)

	var vehicle models.Vehicle
	if err := config.DB.First(&vehicle, booking.VehicleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}

	if !vehicle.Available {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Vehicle not available"})
		return
	}

	days := booking.EndDate.Sub(booking.StartDate).Hours() / 24
	booking.TotalPrice = vehicle.PricePerDay * float64(days)

	tx := config.DB.Begin()
	if err := tx.Create(&booking).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	if err := tx.Model(&vehicle).Update("available", false).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vehicle"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusCreated, booking)
}
