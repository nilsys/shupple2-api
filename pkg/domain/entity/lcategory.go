package entity

import "time"

type (
	Lcategory struct {
		CategoryBase
		CreatedAt time.Time `gorm:"-"`
		UpdatedAt time.Time `gorm:"-"`
	}
)

func (lc Lcategory) CategoryType() string {
	return "lcategory"
}
