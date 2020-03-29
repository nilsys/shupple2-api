package entity

import (
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"gopkg.in/guregu/null.v3"
)

type (
	ThemeCategory struct {
		CategoryBase
		Type       model.ThemeCategoryType
		ThemeID    int
		SubThemeID null.Int
		CreatedAt  time.Time `gorm:"-;default:current_timestamp"`
		UpdatedAt  time.Time `gorm:"-;default:current_timestamp"`
	}
)

func (tc ThemeCategory) CategoryType() string {
	return tc.Type.String()
}
