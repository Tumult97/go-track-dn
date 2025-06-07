package jwt

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	entities "tracker/models/entities"
	jwtModels "tracker/models/jwt"
)

var (
	secretKey     []byte
	tokenLifetime time.Duration = 24 * time.Hour
	issuerName    string        = "tracker-api"
)

func Init(config jwtModels.Config) {
	if config.Secret != "" {
		secretKey = []byte(config.Secret)
	} else {
		secretKey = []byte("your_jwt_secret_key_change_this_in_production")
	}

	if config.Expiration > 0 {
		tokenLifetime = config.Expiration
	}

	if config.Issuer != "" {
		issuerName = config.Issuer
	}
}

func GenerateToken(user entities.User) (string, error) {
	if len(secretKey) == 0 {
		Init(jwtModels.Config{})
	}

	expirationTime := time.Now().Add(tokenLifetime)

	claims := &jwtModels.UserClaims{
		UserID:    user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
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

func ValidateToken(tokenString string) (*jwtModels.UserClaims, error) {
	if len(secretKey) == 0 {
		Init(jwtModels.Config{})
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
