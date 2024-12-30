package controllers

import (
	"booking_api/config"
	"booking_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateCustomer(c *gin.Context) {
	// Ensure only the customer themselves or an admin can update.
	userID := c.GetUint("user_id") // Get user_id from middleware.
	role, _ := c.Get("role")

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Allow admin or the user themselves to update.
	if role != "admin" && user.ID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
		return
	}

	// Bind the JSON body to the user model.
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the updated user.
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User profile updated successfully", "user": user})
}

func GetAllCustomers(c *gin.Context) {
	// Ensure the request is from an admin.
	if role, _ := c.Get("role"); role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}

	var customers []models.User
	// Fetch all users with the role 'customer'.
	if err := config.DB.Where("role = ?", "customer").Find(&customers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve customers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"customers": customers})
}

func DeleteCustomer(c *gin.Context) {
	// Ensure only the customer themselves or an admin can delete.
	userID := c.GetUint("user_id") // Get user_id from middleware.
	role, _ := c.Get("role")

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if role != "admin" && user.ID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
		return
	}

	// Delete the customer.
	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}
