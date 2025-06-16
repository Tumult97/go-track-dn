package jwt

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	entities "tracker/models/entities"
	jwtModels "tracker/models/jwt"
)

var (
	secretKey            []byte
	accessTokenLifetime  time.Duration = 15 * time.Minute
	refreshTokenLifetime time.Duration = 7 * 24 * time.Hour // 7 days
	issuerName           string        = "tracker-api"
)

func Init() {
	secretString := os.Getenv("JWT_SECRET")
	if secretString == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}
	secretKey = []byte(secretString)

	// Check for custom token lifetimes from environment
	if accessExp := os.Getenv("ACCESS_TOKEN_EXPIRATION_MINUTES"); accessExp != "" {
		if minutes, err := strconv.Atoi(accessExp); err == nil {
			accessTokenLifetime = time.Duration(minutes) * time.Minute
		}
	}

	if refreshExp := os.Getenv("REFRESH_TOKEN_EXPIRATION_DAYS"); refreshExp != "" {
		if days, err := strconv.Atoi(refreshExp); err == nil {
			refreshTokenLifetime = time.Duration(days) * 24 * time.Hour
		}
	}

	issuerName = os.Getenv("TOKEN_ISSUER")
	if issuerName == "" {
		issuerName = "tracker-api" // default value
	}
}

// GenerateAccessToken creates a short-lived access token
func GenerateAccessToken(user entities.User) (string, error) {
	if len(secretKey) == 0 {
		Init()
	}

	expirationTime := time.Now().Add(accessTokenLifetime)

	claims := &jwtModels.UserClaims{
		UserID:    user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		TokenType: jwtModels.AccessToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    issuerName,
			Subject:   fmt.Sprintf("%d", user.ID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GenerateRefreshToken creates a long-lived refresh token
func GenerateRefreshToken(user entities.User) (string, error) {
	if len(secretKey) == 0 {
		Init()
	}

	expirationTime := time.Now().Add(refreshTokenLifetime)

	claims := &jwtModels.RefreshTokenClaims{
		UserID:    user.ID,
		TokenType: jwtModels.RefreshToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    issuerName,
			Subject:   fmt.Sprintf("%d", user.ID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GenerateTokenPair creates both access and refresh tokens
func GenerateTokenPair(user entities.User) (accessToken, refreshToken string, err error) {
	accessToken, err = GenerateAccessToken(user)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = GenerateRefreshToken(user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// GenerateToken (legacy function) - maintains backward compatibility
func GenerateToken(user entities.User) (string, error) {
	return GenerateAccessToken(user)
}

func ValidateToken(tokenString string) (*jwtModels.UserClaims, error) {
	if len(secretKey) == 0 {
		Init()
	}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwtModels.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*jwtModels.UserClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func GetUserTokenFromContext(context *gin.Context) (*string, error) {
	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("auth header is empty")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, fmt.Errorf("authorization header format must be Bearer {token}")
	}

	token := &parts[1]

	return token, nil
}

func ExtractUserID(tokenString string) (int, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return 0, err
	}
	return claims.UserID, nil
}

func ExtractUserIDFromContext(c *gin.Context) (int, error) {
	tokenString, err := GetUserTokenFromContext(c)

	if err != nil || tokenString == nil {
		return 0, err
	}

	// * dereferences the pointer from tokenstring
	// Using * in a function definition makes it a nullable pointer variable
	// adding & to a non pointer variable makes it a pointer variable
	return ExtractUserID(*tokenString)
}

// ValidateRefreshToken validates a refresh token and returns the claims
func ValidateRefreshToken(tokenString string) (*jwtModels.RefreshTokenClaims, error) {
	if len(secretKey) == 0 {
		Init()
	}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwtModels.RefreshTokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(*jwtModels.RefreshTokenClaims)
	if !ok {
		return nil, errors.New("invalid refresh token claims")
	}

	// Verify this is actually a refresh token
	if claims.TokenType != jwtModels.RefreshToken {
		return nil, errors.New("token is not a refresh token")
	}

	return claims, nil
}

// ValidateAccessToken specifically validates access tokens
func ValidateAccessToken(tokenString string) (*jwtModels.UserClaims, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	// Verify this is actually an access token
	if claims.TokenType != jwtModels.AccessToken {
		return nil, errors.New("token is not an access token")
	}

	return claims, nil
}
