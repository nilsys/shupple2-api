package entity

import "time"

type (
	Comic struct {
		ID        int `gorm:"primary_key"`
		UserID    int
		Slug      string
		Thumbnail string
		Title     string
		Body      string
		CreatedAt time.Time `gorm:"default:current_timestamp"`
		UpdatedAt time.Time `gorm:"default:current_timestamp"`
		DeletedAt *time.Time
	}

	QueryComic struct {
		Comic
		User *User `gorm:"foreignkey:UserID"`
	}
)

// テーブル名
func (queryComic *QueryComic) TableName() string {
	return "comic"
}
