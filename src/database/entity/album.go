package entity

type Album struct {
	ID      int64 `gorm:"primary_key"`
	Name    string
	Artists []Artist `gorm:"many2many:album_artist;"`
}
