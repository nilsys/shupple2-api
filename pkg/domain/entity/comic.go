package entity

import "time"

type (
	Comic struct {
		ID        int `gorm:"primary_key"`
		UserID    int
		Slug      string
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

// TODO: サムネイル生成ロジック
func (comic *Comic) GenerateThumbnailURL() string {
	return "thumbnailURL"
}

// テーブル名
func (queryComic *QueryComic) TableName() string {
	return "comic"
}
