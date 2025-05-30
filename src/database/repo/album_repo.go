package repo

import (
	"scrobblium-web/database"
	"scrobblium-web/database/entity"
)

func SaveAlbum(album entity.Album) (entity.Album, error) {
	err := database.DB.Create(&album).Error
	return album, err
}

func GetAlbumByName(name string) (entity.Album, error) {
	var album entity.Album
	err := database.DB.Where("name = ?", name).First(&album).Error
	return album, err
}
