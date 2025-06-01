package stats

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrobblium-web/database/repo"
	"time"
)

func ParseRange(c *gin.Context) (string, int64, int64, int64, bool) {
	rangeParam := c.DefaultQuery("range", "all")
	now := time.Now().Unix()

	var from int64
	var dayCount int64

	switch rangeParam {
	case "last7":
		from = now - 7*24*60*60
		dayCount = 7
	case "last30":
		from = now - 30*24*60*60
		dayCount = 30
	case "lastY":
		from = now - 365*24*60*60
		dayCount = 365
	case "all":
		track, _ := repo.GetEarliestTrack()
		from = track.ListenedAt
		dayCount = (now - from) / (60 * 60 * 24)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid range param"})
		return "", 0, 0, 0, true
	}
	return rangeParam, now, from, dayCount, false
}
