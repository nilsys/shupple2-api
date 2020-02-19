package entity

type Hashtag struct {
	ID   int `gorm:"primary_key"`
	Name string
}
