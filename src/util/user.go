package util

import (
	"github.com/gin-gonic/gin"
	"log"
)

// DenyGuestUser return true if the user was guest and return Unauthorized
func DenyGuestUser(c *gin.Context) bool {
	guestBool := IsGuestUser(c)
	if guestBool {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return true
	}
	return false
}
func IsGuestUser(c *gin.Context) bool {
	guest, exists := c.Get("isGuest")
	if !exists {
		log.Fatalln("No isGuest found")
		return false
	}
	guestBool := guest.(bool)
	return guestBool
}
