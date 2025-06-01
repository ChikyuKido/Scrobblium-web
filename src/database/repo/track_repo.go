package repo

import (
	"scrobblium-web/database"
	"scrobblium-web/database/entity"
)

func SaveTrack(track entity.Track) (entity.Track, error) {
	err := database.DB.Create(&track).Error
	return track, err
}

func GetTracks() ([]entity.Track, error) {
	var tracks []entity.Track
	err := database.DB.Find(&tracks).Error
	return tracks, err
}

func GetEarliestTrack() (entity.Track, error) {
	var track entity.Track
	err := database.DB.Order("listened_at asc").Limit(1).First(&track).Error
	return track, err
}
