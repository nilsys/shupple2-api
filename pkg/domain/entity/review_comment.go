package entity

import "time"

type (
	ReviewComment struct {
		ID        int `gorm:"column:id"`
		UserID    int
		ReviewID  int
		User      *User     `gorm:"foreignkey:UserID"`
		Body      string    `gorm:"column:body"`
		CreatedAt time.Time `gorm:"-;default:current_timestamp"`
		UpdatedAt time.Time `gorm:"-;default:current_timestamp"`
	}
)
