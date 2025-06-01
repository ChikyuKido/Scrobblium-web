package stats

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrobblium-web/database/repo"
)

func CardStats() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, _, from, dayCount, done := ParseRange(c)
		if done {
			return
		}

		tracks, err := repo.GetTracks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tracks"})
			return
		}

		var totalScrobbles int64
		var totalListeningTime int64
		var totalMaxProgress int64

		for _, track := range tracks {
			if track.ListenedAt >= from {
				totalScrobbles++
				totalListeningTime += track.TimeListened
				totalMaxProgress += track.MaxProgress
			}
		}

		var ratio float64
		if totalMaxProgress > 0 {
			ratio = float64(totalListeningTime) / float64(totalMaxProgress)
		}
		averageScrobbles := float64(totalScrobbles) / float64(dayCount)
		c.JSON(http.StatusOK, gin.H{
			"totalScrobbles":   totalScrobbles,
			"averageScrobbles": averageScrobbles,
			"ratio":            ratio,
			"listeningTime":    totalListeningTime,
		})
	}
}
