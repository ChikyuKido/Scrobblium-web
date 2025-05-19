package util

import (
	"github.com/gin-gonic/gin"
	"log"
)

func AllowsGuestUser(c *gin.Context, allows bool) bool {
	guest, exists := c.Get("isGuest")
	if !exists {
		log.Fatalln("No isGuest found")
		return false
	}
	guestBool := guest.(bool)
	if guestBool && !allows {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return false
	}
	return true
}
