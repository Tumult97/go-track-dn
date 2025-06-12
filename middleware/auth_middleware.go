package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	jwtModels "tracker/models/jwt"
	"tracker/utils/jwt"
	tc "tracker/utils/trycatch"
)

const UserKey = "user"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer tc.Catch(func(err error) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		})

		tokenString := tc.Try(jwt.GetUserTokenFromContext(c))

		if tokenString == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		claims := tc.Try(jwt.ValidateAccessToken(*tokenString))

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
