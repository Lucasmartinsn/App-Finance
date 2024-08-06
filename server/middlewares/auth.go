package middlewares

import (
	jwttoken "development/application/fiance/server/services/jwtToken"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth_default() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"mensage": "token authentication not provided",
			})
			c.Abort()
			return
		}

		token := header[len(BearerSchema):]
		if valid, err := jwttoken.NewJWTService_Default().ValidateToken_Default(token); err != nil || !valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"mensage": "invalid token authentication",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}