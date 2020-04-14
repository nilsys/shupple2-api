package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	TouristSpotCommandRepository interface {
		Lock(c context.Context, id int) (*entity.TouristSpot, error)
		Store(c context.Context, touristSpot *entity.TouristSpot) error
		// レビューの平均値を更新
		UpdateScoreByID(c context.Context, id int) error
		DeleteByID(id int) error
	}

	TouristSpotQueryRepository interface {
		FindAll() ([]*entity.TouristSpot, error)
		FindByID(id int) (*entity.TouristSpot, error)
		FindDetailByID(id int) (*entity.TouristSpotDetail, error)
		FindListByParams(query *query.FindTouristSpotListQuery) ([]*entity.TouristSpot, error)
		FindRecommendListByParams(query *query.FindRecommendTouristSpotListQuery) ([]*entity.TouristSpot, error)
		// name部分一致検索
		SearchByName(name string) ([]*entity.TouristSpot, error)
	}
)
