package entity

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"gopkg.in/guregu/null.v3"
)

type (
	SpotCategory struct {
		CategoryBase
		Type              model.SpotCategoryType
		SpotCategoryID    int
		SubSpotCategoryID null.Int
		TimesWithoutDeletedAt
	}
)

func (lc SpotCategory) CategoryType() string {
	return lc.Type.String()
}
