package entity

import "time"

type (
	Times struct {
		CreatedAt time.Time  `gorm:"-;default:current_timestamp"`
		UpdatedAt time.Time  `gorm:"-;default:current_timestamp"`
		DeletedAt *time.Time `json:"-"`
	}

	TimesWithoutDeletedAt struct {
		CreatedAt time.Time `gorm:"-;default:current_timestamp"`
		UpdatedAt time.Time `gorm:"-;default:current_timestamp"`
	}
)
