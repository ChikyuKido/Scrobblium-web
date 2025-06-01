package server

import (
	"github.com/gin-gonic/gin"
	"scrobblium-web/server/middleware"
	"scrobblium-web/server/routes/auth"
	"scrobblium-web/server/routes/settings"
	"scrobblium-web/server/routes/stats"
	"scrobblium-web/server/routes/track"
	"strings"
)

func Init() {
	r := gin.Default()
	r.Static("/assets", "./dist/assets")
	r.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.JSON(404, gin.H{"error": "not found"})
			return
		}
		c.File("./dist/index.html")
	})

	apiGroup := r.Group("/api/v1")
	apiGroup.Use(middleware.AuthMiddleware())
	initTrackRoutes(apiGroup)
	initAuthRoutes(apiGroup)
	initStatRouter(apiGroup)
	initSettingsRouter(apiGroup)
	err := r.Run()
	if err != nil {
		return
	}
}
func initTrackRoutes(r *gin.RouterGroup) {
	trackGroup := r.Group("/track")
	trackGroup.POST("/add", track.Add())
}
func initAuthRoutes(r *gin.RouterGroup) {
	authGroup := r.Group("/auth")
	authGroup.POST("/login", auth.Login())
	authGroup.GET("/isAuthorized/:url", auth.IsAuthorized())
}

func initStatRouter(r *gin.RouterGroup) {
	statGroup := r.Group("/stats")
	statGroup.GET("card", stats.CardStats())
	statGroup.GET("scrobbleOverTime", stats.ScrobbleOverTime())
	statGroup.GET("top5artists", stats.Top5Artists())
	statGroup.GET("top5albums", stats.Top5Albums())
	statGroup.GET("hourly", stats.Hourly())
}

func initSettingsRouter(r *gin.RouterGroup) {
	settingsGroup := r.Group("/settings")
	settingsGroup.GET("get", settings.GetSettings())
	settingsGroup.POST("save", settings.SaveSettings())
	settingsGroup.POST("password", settings.ChangePassword())
}
