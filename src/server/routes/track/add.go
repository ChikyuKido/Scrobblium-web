package track

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrobblium-web/database/entity"
	"scrobblium-web/database/repo"
	"scrobblium-web/util"
)

func Add() gin.HandlerFunc {
	return func(c *gin.Context) {
		if util.DenyGuestUser(c) {
			return
		}
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

		var trackArtists []entity.Artist
		for _, artistName := range request.Artists {
			artist, err := repo.GetArtistByName(artistName)
			if err != nil {
				artist = entity.Artist{Name: artistName}
				artist, _ = repo.SaveArtist(artist)
			}
			trackArtists = append(trackArtists, artist)
		}

		// Handle album artists
		var albumArtists []entity.Artist
		for _, artistName := range request.Album.Artists {
			artist, err := repo.GetArtistByName(artistName)
			if err != nil {
				artist = entity.Artist{Name: artistName}
				artist, _ = repo.SaveArtist(artist)
			}
			albumArtists = append(albumArtists, artist)
		}

		// Handle album
		album, err := repo.GetAlbumByName(request.Album.Name)
		if err != nil {
			album = entity.Album{
				Name:    request.Album.Name,
				Artists: albumArtists,
			}
			album, _ = repo.SaveAlbum(album)
		}

		// Create and save track
		track := entity.Track{
			ID:           request.ID,
			MaxProgress:  request.MaxProgress,
			Progress:     request.Progress,
			TimeListened: request.TimeListened,
			Name:         request.Name,
			Artists:      trackArtists,
			ListenedAt:   request.ListenedAt,
			Album:        album,
			AlbumID:      album.ID,
		}

		_, err = repo.SaveTrack(track)
		if err != nil {
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "failed to save track",
					"details": err.Error(),
				})
				return
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "track saved"})
	}
}
