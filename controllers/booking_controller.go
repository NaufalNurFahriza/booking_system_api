package controllers

import (
	"booking_api/config"
	"booking_api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateBooking creates a new booking
func CreateBooking(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var booking models.Booking

	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign user ID from the authenticated user
	booking.UserID = userID.(uint)
	booking.BookingDate = time.Now()
	booking.PaymentStatus = "unpaid"

	// Save booking to the database
	if err := config.DB.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Booking created successfully",
		"booking": booking,
	})
}

// GetMyBookings retrieves bookings for the authenticated user
func GetMyBookings(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var bookings []models.Booking

	// Fetch bookings for the logged-in user
	if err := config.DB.Preload("Schedule").Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bookings": bookings})
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
