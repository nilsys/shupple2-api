package repository

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

func newCategory(id int) *entity.Category {
	category := &entity.Category{ID: id}
	util.FillDymmyString(category, id)
	return category
}
