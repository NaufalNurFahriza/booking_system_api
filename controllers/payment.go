package controllers

import (
	"booking_api/database"
	"booking_api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get booking to verify ownership and status
	var booking models.Booking
	if err := database.DB.First(&booking, payment.BookingID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	// Verify user owns the booking
	userID, _ := c.Get("user_id")
	if booking.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to pay for this booking"})
		return
	}

	payment.PaymentDate = time.Now()

	tx := database.DB.Begin()
	if err := tx.Create(&payment).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
		return
	}

	// Update booking status if payment is completed
	if payment.PaymentStatus == "completed" {
		booking.Status = "confirmed"
		if err := tx.Save(&booking).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking status"})
			return
		}
	}

	tx.Commit()
	c.JSON(http.StatusCreated, payment)
}

func GetBookingPayments(c *gin.Context) {
	bookingID := c.Param("bookingID")
	var payments []models.Payment

	// Verify booking ownership
	var booking models.Booking
	if err := database.DB.First(&booking, bookingID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	userID, _ := c.Get("user_id")
	if booking.UserID != userID.(uint) && c.GetString("role") != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to view these payments"})
		return
	}

	if err := database.DB.Where("booking_id = ?", bookingID).Find(&payments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch payments"})
		return
	}

	c.JSON(http.StatusOK, payments)
}

func UpdatePayment(c *gin.Context) {
	if c.GetString("role") != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}

	var payment models.Payment
	if err := database.DB.First(&payment, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := database.DB.Begin()
	if err := tx.Save(&payment).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update payment"})
		return
	}

	// Update booking status if payment status changes
	if payment.PaymentStatus == "completed" {
		var booking models.Booking
		if err := tx.First(&booking, payment.BookingID).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find booking"})
			return
		}

		booking.Status = "confirmed"
		if err := tx.Save(&booking).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking status"})
			return
		}
	}

	tx.Commit()
	c.JSON(http.StatusOK, payment)
}

func GetAllPayments(c *gin.Context) {
	var payments []models.Payment

	if err := database.DB.Find(&payments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch payments"})
		return
	}

	c.JSON(http.StatusOK, payments)
}
