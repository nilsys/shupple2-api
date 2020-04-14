package service

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	// TouristSpot参照系サービス
	TouristSpotQueryService interface {
		Show(id int) (*entity.TouristSpotDetail, error)
		List(query *query.FindTouristSpotListQuery) ([]*entity.TouristSpot, error)
		ListRecommend(query *query.FindRecommendTouristSpotListQuery) ([]*entity.TouristSpot, error)
	}

	// TouristSpot参照系サービス実装
	TouristSpotQueryServiceImpl struct {
		repository.TouristSpotQueryRepository
	}
)

var TouristSpotQueryServiceSet = wire.NewSet(
	wire.Struct(new(TouristSpotQueryServiceImpl), "*"),
	wire.Bind(new(TouristSpotQueryService), new(*TouristSpotQueryServiceImpl)),
)

// TouristSpot参照
func (s *TouristSpotQueryServiceImpl) Show(id int) (*entity.TouristSpotDetail, error) {
	return s.TouristSpotQueryRepository.FindDetailByID(id)
}

// TouristSpot一覧参照
func (s *TouristSpotQueryServiceImpl) List(query *query.FindTouristSpotListQuery) ([]*entity.TouristSpot, error) {
	return s.TouristSpotQueryRepository.FindListByParams(query)
}

// おすすめTouristSpot一覧取得
func (s *TouristSpotQueryServiceImpl) ListRecommend(query *query.FindRecommendTouristSpotListQuery) ([]*entity.TouristSpot, error) {
	return s.TouristSpotQueryRepository.FindRecommendListByParams(query)
}
