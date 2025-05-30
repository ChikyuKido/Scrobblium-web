package repo

import (
	"scrobblium-web/database"
	"scrobblium-web/database/entity"
)

func SaveArtist(artist entity.Artist) (entity.Artist, error) {
	err := database.DB.Create(&artist).Error
	return artist, err
}

func GetArtistByName(name string) (entity.Artist, error) {
	var artist entity.Artist
	err := database.DB.Where("name = ?", name).First(&artist).Error
	return artist, err
}
