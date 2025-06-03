package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	jwtModels "tracker/models/jwt"
	"tracker/utils/jwt"
)

const UserKey = "user"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is required",
			})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header format must be Bearer {token}",
			})
			return
		}

		tokenString := parts[1]

		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			return
		}

		c.Set(UserKey, claims)

		c.Next()
	}
}

func GetUserFromContext(c *gin.Context) (*jwtModels.UserClaims, bool) {
	user, exists := c.Get(UserKey)
	if !exists {
		return nil, false
	}

	claims, ok := user.(*jwtModels.UserClaims)
	if !ok {
		return nil, false
	}
	return claims, true
}
