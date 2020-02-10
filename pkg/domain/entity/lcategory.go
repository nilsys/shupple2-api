package entity

import "time"

type (
	Lcategory struct {
		ID        int `gorm:"primary_key"`
		Name      string
		CreatedAt time.Time `gorm:"-"`
		UpdatedAt time.Time `gorm:"-"`
	}
)
