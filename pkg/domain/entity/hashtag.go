package entity

type (
	Hashtag struct {
		ID          int `gorm:"primary_key"`
		Name        string
		PostCount   int
		ReviewCount int
		Score       int
	}
)
