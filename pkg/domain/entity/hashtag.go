package entity

type (
	Hashtag struct {
		ID          int `gorm:"primary_key"`
		Name        string
		PostCount   int
		ReviewCount int
		Score       int
	}

	HashtagCategory struct {
		HashtagID  int `gorm:"primary_key"`
		CategoryID int `gorm:"primary_key"`
	}
)

func NewHashtagCategory(hashtagID, categoryID int) *HashtagCategory {
	return &HashtagCategory{
		HashtagID:  hashtagID,
		CategoryID: categoryID,
	}
}
