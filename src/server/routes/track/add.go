package track

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Add() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			ID           int64    `json:"id"`
			MaxProgress  int64    `json:"max_progress"`
			Progress     int64    `json:"progress"`
			TimeListened int64    `json:"time_listened"`
			Name         string   `json:"name"`
			Artists      []string `json:"artists"`
			ListenedAt   int64    `json:"listened_at"`
			Album        struct {
				Name    string   `json:"name"`
				Artists []string `json:"artists"`
			} `json:"album"`
		}
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json struct"})
			return
		}

		c.JSON(200, gin.H{})
	}
}
