package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"scrobblium-web/database/entity"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(sqlite.Open("scrob.db"), &gorm.Config{})
	_ = db.Exec(`
		PRAGMA journal_mode=WAL;
		PRAGMA synchronous=OFF;
		PRAGMA temp_store=MEMORY;
		PRAGMA cache_size=10000;
	`)
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	err = db.AutoMigrate(&entity.Album{}, &entity.Artist{}, &entity.Track{}, &entity.Settings{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
		return
	}

	DB = db
}
