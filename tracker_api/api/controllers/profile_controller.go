package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tracker/middleware"
)

// GetProfile returns the profile information of the logged-in user
func GetProfile(c *gin.Context) {
	// Get user from context (set by the auth middleware)
	claims, exists := middleware.GetUserFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})
		return
	}

	// Return user profile data
	c.JSON(http.StatusOK, gin.H{
		"id":         claims.UserID,
		"email":      claims.Email,
		"first_name": claims.FirstName,
		"last_name":  claims.LastName,
		"message":    "Profile retrieved successfully",
	})
}

