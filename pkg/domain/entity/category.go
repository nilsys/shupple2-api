package entity

import (
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	Category struct {
		ID        int `gorm:"primary_key"`
		Name      string
		Type      model.CategoryType
		CreatedAt time.Time `gorm:"-"`
		UpdatedAt time.Time `gorm:"-"`
	}
)
