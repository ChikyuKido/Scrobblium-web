package entity

type Artist struct {
	ID   int64 `gorm:"primary_key"`
	Name string
}
