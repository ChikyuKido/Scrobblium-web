package stats

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrobblium-web/database"
)

func Top5Artists() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, _, from, _, done := ParseRange(c)
		if done {
			return
		}

		type result struct {
			Name  string
			Count int
		}
		var results []result

		query := `
		SELECT a.name AS name, COUNT(DISTINCT t.id) AS count
		FROM track_artists ta
		JOIN artists a ON a.id = ta.artist_id
		JOIN tracks t ON t.id = ta.track_id
		WHERE t.listened_at >= ?
		GROUP BY a.name
		ORDER BY count DESC
		LIMIT 5
		`

		err := database.DB.Raw(query, from).Scan(&results).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch top artists"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"topArtists": results,
		})
	}
}
