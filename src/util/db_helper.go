package util

import (
	"log"
	"scrobblium-web/database"
	"scrobblium-web/database/entity"
	"scrobblium-web/database/repo"
)

func AddDefaultSettings() {
	_, err := repo.GetSettings()
	if err != nil {
		log.Println("settings not found -> created:", err)
		hashed, _ := HashPassword("admin")
		database.DB.Create(&entity.Settings{
			Username:       "admin",
			Password:       hashed,
			HomepageAccess: false,
			ListAccess:     false,
		})
	}
}
