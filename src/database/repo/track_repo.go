package repo

import (
	"scrobblium-web/database"
	"scrobblium-web/database/entity"
)

func SaveTrack(track entity.Track) (entity.Track, error) {
	err := database.DB.Create(&track).Error
	return track, err
}
