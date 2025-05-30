package repo

import (
	"scrobblium-web/database"
	"scrobblium-web/database/entity"
)

func GetSettings() (entity.Settings, error) {
	var settings entity.Settings
	err := database.DB.First(&settings).Error
	return settings, err
}

func GetUsername() (string, error) {
	settings, err := GetSettings()
	if err != nil {
		return "", err
	}
	return settings.Username, nil
}

func GetPassword() (string, error) {
	settings, err := GetSettings()
	if err != nil {
		return "", err
	}
	return settings.Password, nil
}
