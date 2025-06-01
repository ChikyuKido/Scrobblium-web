package entity

type Settings struct {
	Id             int    `gorm:"primary_key" json:"id"`
	Username       string `gorm:"not null" json:"username"`
	Password       string `gorm:"not null" json:"-"`
	HomepageAccess bool   `gorm:"not null" json:"homepageAccess"`
	ListAccess     bool   `gorm:"not null" json:"listAccess"`
}
