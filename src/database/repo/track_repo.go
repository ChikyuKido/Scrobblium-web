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
func GetTracksWithAll() ([]entity.Track, error) {
	var tracks []entity.Track
	err := database.DB.Preload("Album").Find(&tracks).Error
	if err != nil {
		return nil, err
	}

	for i := range tracks {
		var artists []entity.Artist
		err = database.DB.Model(&tracks[i]).Association("Artists").Find(&artists)
		if err != nil {
			return nil, err
		}
		tracks[i].Artists = artists
	}

	return tracks, nil
}

func GetEarliestTrack() (entity.Track, error) {
	var track entity.Track
	err := database.DB.Order("listened_at asc").Limit(1).First(&track).Error
	return track, err
}
