package jwt

import (
	"errors"
	"fmt"
	"time"

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

func ExtractUserID(tokenString string) (int, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return 0, err
	}
	return claims.UserID, nil
}
