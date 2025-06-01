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
func UpdateSettings(username string, homepageAccess, listAccess bool) {
	settings, _ := GetSettings()
	settings.Username = username
	settings.HomepageAccess = homepageAccess
	settings.ListAccess = listAccess
	database.DB.Save(&settings)
}
func UpdatePassword(password string) {
	settings, _ := GetSettings()
	settings.Password = password
	database.DB.Save(&settings)
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
