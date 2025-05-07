package entity

type Track struct {
	ID           int64    `gorm:"primary_key"`
	MaxProgress  int64    `gorm:"not null"`
	Progress     int64    `gorm:"not null"`
	TimeListened int64    `gorm:"not null"`
	Name         string   `gorm:"not null"`
	Artists      []Artist `gorm:"many2many:track_artists;"`
	ListenedAt   int64    `gorm:"not null"`
	Album        Album
	AlbumID      int64 `gorm:"foreignkey:AlbumID"`
}
