package entity

import (
	"database/sql"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	ThemeCategory struct {
		CategoryBase
		Type       model.ThemeCategoryType
		ThemeID    int
		SubThemeID sql.NullInt64
		CreatedAt  time.Time `gorm:"-;default:current_timestamp"`
		UpdatedAt  time.Time `gorm:"-;default:current_timestamp"`
	}
)

func (tc ThemeCategory) CategoryType() string {
	return tc.Type.String()
}
