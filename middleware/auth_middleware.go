package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	jwtModels "tracker/models/jwt"
	"tracker/utils/jwt"
)

const UserKey = "user"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := jwt.GetUserTokenFromContext(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		if tokenString == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		claims, err := jwt.ValidateToken(*tokenString)
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
