package settings

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrobblium-web/database/repo"
)

func GetSettings() gin.HandlerFunc {
	return func(c *gin.Context) {
		settings, _ := repo.GetSettings()
		c.JSON(http.StatusOK, settings)
	}
}
