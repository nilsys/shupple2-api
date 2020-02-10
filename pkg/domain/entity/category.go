package entity

import "time"

type (
	Category struct {
		ID        int `gorm:"primary_key"`
		Name      string
		CreatedAt time.Time `gorm:"-"`
		UpdatedAt time.Time `gorm:"-"`
	}
)
