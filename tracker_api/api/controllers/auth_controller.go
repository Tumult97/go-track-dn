package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"tracker/models/entities"
	"tracker/models/requests"
	"tracker/models/responses"
	"tracker/repository/repos"
	"tracker/utils/jwt"
)

var userRepo = repos.NewUserRepository()

func Register(c *gin.Context) {
	var request requests.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Invalid request data: %s", err.Error()),
		})
		return
	}

	user := entities.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
		CreatedAt: time.Now(),
	}

	id, err := userRepo.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to create user: %s", err.Error()),
		})
		return
	}

	user.ID = id

	c.JSON(http.StatusCreated, responses.AuthResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		Message:   "User registered successfully",
	})
}

func Login(c *gin.Context) {
	var request requests.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Invalid request data: %s", err.Error()),
		})
		return
	}

	user, err := userRepo.ValidateCredentials(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	accessToken, refreshToken, err := jwt.GenerateTokenPair(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate authentication tokens",
		})
		return
	}

	c.JSON(http.StatusOK, responses.AuthResponse{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		CreatedAt:    user.CreatedAt,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Message:      "Login successful",
	})
}

// RefreshToken handles token refresh requests
func RefreshToken(c *gin.Context) {
	var request requests.RefreshTokenRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Invalid request data: %s", err.Error()),
		})
		return
	}

	// Validate the refresh token
	refreshClaims, err := jwt.ValidateRefreshToken(request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid or expired refresh token",
		})
		return
	}

	// Get user from database using the user ID from refresh token
	user, err := userRepo.GetUserByID(refreshClaims.UserID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		return
	}

	// Generate new access token
	newAccessToken, err := jwt.GenerateAccessToken(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate new access token",
		})
		return
	}

	// Optionally generate new refresh token (token rotation)
	newRefreshToken, err := jwt.GenerateRefreshToken(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate new refresh token",
		})
		return
	}

	c.JSON(http.StatusOK, responses.RefreshTokenResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
		Message:      "Token refreshed successfully",
	})
}
