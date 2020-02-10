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
)
