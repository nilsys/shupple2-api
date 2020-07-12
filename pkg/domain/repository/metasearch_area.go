package repository

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	MetasearchAreaCommandRepository interface {
		Store(*entity.MetasearchArea) error
	}

	MetasearchAreaQueryRepository interface {
		FindByMetasearchAreaID(metasearchAreaID int, metasearchAreaType model.AreaCategoryType) (*entity.MetasearchArea, error)
		FindByAreaCategoryID(areaCategoryID int, areaCategoryType model.AreaCategoryType) ([]*entity.MetasearchArea, error)
	}
)
