package entity

type HashTag struct {
	ID   int `gorm:"primary_key"`
	Name string
}

func (hastTag *HashTag) TableName() string {
	return "hashtag"
}
