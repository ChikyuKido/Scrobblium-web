package server

import (
	"github.com/gin-gonic/gin"
	"scrobblium-web/server/routes/track"
)

func Init() {
	r := gin.Default()
	initTrackRoutes(r)
	err := r.Run()
	if err != nil {
		return
	}
}
func initTrackRoutes(r *gin.Engine) {
	trackGroup := r.Group("/api/v1/track")
	trackGroup.POST("/add", track.Add())
}
