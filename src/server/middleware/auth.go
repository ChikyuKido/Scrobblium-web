package middleware

import (
	"github.com/gin-gonic/gin"
	"scrobblium-web/util"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.Set("isGuest", true)
			c.Next()
			return
		}
		if !strings.HasPrefix(authorizationHeader, "Bearer ") {
			c.JSON(401, gin.H{
				"error": "Invalid authorization format",
			})
			c.Abort()
			return
		}
		tokenString := authorizationHeader[7:]
		token, err := util.GetToken(tokenString)
		if err != nil || !token.Valid {
			c.Set("isGuest", true)
			c.Next()
			return
		}
		c.Set("isGuest", false)
		c.Next()
	}
}
