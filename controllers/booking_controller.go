package controllers

import (
	"booking_api/config"
	"booking_api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateBooking(c *gin.Context) {
	// Safely extract userID
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Convert userID to float64 first (as JWT typically stores numbers as float64)
	userIDFloat, ok := userIDInterface.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Convert float64 to uint
	userID := uint(userIDFloat)

	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify that the schedule exists and preload its movie data
	var schedule models.Schedule
	if err := config.DB.Preload("Movie").First(&schedule, booking.ScheduleID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Schedule not found"})
		return
	}

	// Set booking details
	booking.UserID = userID
	booking.BookingDate = time.Now()
	booking.PaymentStatus = "unpaid"

	// Create the booking
	if err := config.DB.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	// Fetch the complete booking with all relationships
	var completeBooking models.Booking
	if err := config.DB.Preload("User").Preload("Schedule").Preload("Schedule.Movie").First(&completeBooking, booking.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch complete booking data"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Booking created successfully",
		"booking": completeBooking,
	})
}

// UpdateBooking updates a specific booking
func UpdateBooking(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var booking models.Booking

	// Fetch booking by ID
	if err := config.DB.First(&booking, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	// Ensure the user can only update their own bookings
	if booking.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized to update this booking"})
		return
	}

	// Bind the new data to the booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the updated booking
	if err := config.DB.Save(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Booking updated successfully",
		"booking": booking,
	})
}

// DeleteBooking cancels a specific booking
func DeleteBooking(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var booking models.Booking

	// Fetch booking by ID
	if err := config.DB.First(&booking, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	// Ensure the user can only delete their own bookings
	if booking.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized to delete this booking"})
		return
	}

	// Delete the booking
	if err := config.DB.Delete(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete booking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking canceled successfully"})
}

// GetAllBookings retrieves all bookings (admin only)
func GetAllBookings(c *gin.Context) {
	if role, _ := c.Get("role"); role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}

	var bookings []models.Booking
	// Fetch all bookings
	if err := config.DB.Preload("User").Preload("Schedule").Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bookings": bookings})
}
