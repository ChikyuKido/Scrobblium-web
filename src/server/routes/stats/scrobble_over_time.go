package stats

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrobblium-web/database/repo"
)

func ScrobbleOverTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, now, from, dayCount, done := ParseRange(c)
		if done {
			return
		}
		tracks, err := repo.GetTracks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tracks"})
			return
		}

		counts := make(map[int64]int64)

		secondsPerDay := int64(24 * 60 * 60)

		for _, track := range tracks {
			if track.ListenedAt >= from && track.ListenedAt <= now {
				dayIndex := (track.ListenedAt - from) / secondsPerDay
				counts[dayIndex]++
			}
		}

		// Prepare result array with length = dayCount
		data := make([]int64, int(dayCount))
		for i := int64(0); i < dayCount; i++ {
			if val, exists := counts[i]; exists {
				data[i] = val
			} else {
				data[i] = 0
			}
		}
		start := 0
		for start < len(data) && data[start] == 0 {
			start++
		}

		// Trim trailing zeros
		end := len(data) - 1
		for end >= 0 && data[end] == 0 {
			end--
		}

		// If all zeros, return empty array
		if start > end {
			data = []int64{}
		} else {
			data = data[start : end+1]
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
