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
	secretKey     []byte
	tokenLifetime time.Duration = 24 * time.Hour
	issuerName    string        = "tracker-api"
)

func Init() {
	secretString := os.Getenv("JWT_SECRET")
	if secretString == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}
	secretKey = []byte(secretString)

	expireyHours, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_HOURS"))

	if err != nil {
		log.Fatal("TOKEN_EXPIRATION_HOURS environment variable is not set")
	}

	tokenLifetime = time.Duration(expireyHours) * time.Hour

	issuerName = os.Getenv("TOKEN_ISSUER")

	if issuerName == "" {
		log.Fatal("TOKEN_ISSUER environment variable is not set")
	}
}

func GenerateToken(user entities.User) (string, error) {
	if len(secretKey) == 0 {
		Init()
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
