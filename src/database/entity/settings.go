package entity

type Settings struct {
	Id             int    `gorm:"primary_key"`
	Username       string `gorm:"not null"`
	Password       string `gorm:"not null"`
	HomepageAccess bool   `gorm:"not null"`
	ListAccess     bool   `gorm:"not null"`
}
