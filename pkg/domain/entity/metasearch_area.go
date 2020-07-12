package entity

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	MetasearchArea struct {
		MetasearchAreaID   int
		MetasearchAreaType model.AreaCategoryType
		AreaCategoryID     int
		TimesWithoutDeletedAt
	}
)
