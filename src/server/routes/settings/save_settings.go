package settings

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrobblium-web/database/repo"
)

func SaveSettings() gin.HandlerFunc {
	return func(c *gin.Context) {
		var settings = struct {
			Username       string `json:"username"`
			HomepageAccess bool   `json:"homepageAccess"`
			ListAccess     bool   `json:"listAccess"`
		}{}
		if err := c.ShouldBindJSON(&settings); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}
		repo.UpdateSettings(settings.Username, settings.HomepageAccess, settings.ListAccess)
		c.JSON(http.StatusOK, gin.H{"message": "settings saved"})
	}
}
