package stats

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrobblium-web/database"
)

func Top5Albums() gin.HandlerFunc {
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
		SELECT al.name AS name, COUNT(DISTINCT t.id) AS count
		FROM tracks t
		JOIN albums al ON t.album_id = al.id
		WHERE t.listened_at >= ?
		GROUP BY al.name
		ORDER BY count DESC
		LIMIT 5
		`

		err := database.DB.Raw(query, from).Scan(&results).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch top albums"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"topAlbums": results,
		})
	}
}
