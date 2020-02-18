package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

type (
	HashTagQueryRepository interface {
		FindRecommendList(areaID, subAreaID, subSubAreaID int) ([]*entity.HashTag, error)
	}
)
