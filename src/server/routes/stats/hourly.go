package stats

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrobblium-web/database/repo"
	"time"
)

func Hourly() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, _, from, _, done := ParseRange(c)
		if done {
			return
		}

		tracks, err := repo.GetTracks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tracks"})
			return
		}

		hourCounts := make([]int, 24)

		for _, track := range tracks {
			if track.ListenedAt < from {
				continue
			}
			hour := time.Unix(track.ListenedAt, 0).UTC().Hour()
			hourCounts[hour]++
		}

		c.JSON(http.StatusOK, gin.H{
			"hourly": hourCounts,
		})

	}
}
