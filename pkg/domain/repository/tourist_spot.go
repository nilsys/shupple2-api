package repository

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	TouristSpotCommandRepository interface {
		Store(touristSpot *entity.TouristSpot) error
		DeleteByID(id int) error
	}

	TouristSpotQueryRepository interface {
		FindByID(id int) (*entity.QueryTouristSpot, error)
		FindListByParams(query *query.FindTouristSpotListQuery) ([]*entity.TouristSpot, error)
		FindRecommendListByParams(query *query.FindRecommendTouristSpotListQuery) ([]*entity.TouristSpot, error)
		// name部分一致検索
		SearchByName(name string) ([]*entity.TouristSpot, error)
	}
)
