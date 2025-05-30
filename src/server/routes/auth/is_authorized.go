package auth

import (
	"github.com/gin-gonic/gin"
	"scrobblium-web/database/repo"
	"scrobblium-web/util"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Param("url")
		isGuest := util.IsGuestUser(c)
		if !isGuest {
			c.JSON(200, gin.H{})
		}
		settings, err := repo.GetSettings()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}

		if url == "homepage" && !settings.HomepageAccess {
			if util.DenyGuestUser(c) {
				return
			}
		} else if url == "list" && !settings.ListAccess {
			if util.DenyGuestUser(c) {
				return
			}
		} else if url == "sources" || url == "settings" {
			if util.DenyGuestUser(c) {
				return
			}
		}
		c.JSON(200, gin.H{})
	}
}
