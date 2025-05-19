package server

import (
	"github.com/gin-gonic/gin"
	"scrobblium-web/server/middleware"
	"scrobblium-web/server/routes/auth"
	"scrobblium-web/server/routes/track"
)

func Init() {
	r := gin.Default()
	apiGroup := r.Group("/api/v1")
	apiGroup.Use(middleware.AuthMiddleware())
	initTrackRoutes(apiGroup)
	initAuthRoutes(apiGroup)
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
}
