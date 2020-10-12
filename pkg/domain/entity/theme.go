package entity

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	null "gopkg.in/guregu/null.v3"
)

type (
	ThemeCategory struct {
		CategoryBase
		Type       model.ThemeCategoryType
		ThemeID    int
		SubThemeID null.Int
		TimesWithoutDeletedAt
	}

	ThemeCategoryWithPostCount struct {
		CategoryBase
		Type      model.ThemeCategoryType
		PostCount int
	}
)

func (tc ThemeCategory) CategoryType() string {
	return tc.Type.String()
}

func (t *ThemeCategoryWithPostCount) TableName() string {
	return "theme_category"
}
