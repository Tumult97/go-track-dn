package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

// TokenType represents the type of JWT token
type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

type UserClaims struct {
	UserID    int       `json:"user_id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	TokenType TokenType `json:"token_type"`
	jwt.RegisteredClaims
}

// RefreshTokenClaims represents claims for refresh tokens
type RefreshTokenClaims struct {
	UserID    int       `json:"user_id"`
	TokenType TokenType `json:"token_type"`
	jwt.RegisteredClaims
}
