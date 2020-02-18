package service

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	// HashTag参照系サービス
	HashTagQueryService interface {
		ShowRecommendList(areID, subAreaID, subSubAreaID int) ([]*entity.HashTag, error)
	}

	// HashTag参照系サービス実装
	HashTagQueryServiceImpl struct {
		repository.HashTagQueryRepository
	}
)

var HashTagQueryServiceSet = wire.NewSet(
	wire.Struct(new(HashTagQueryServiceImpl), "*"),
	wire.Bind(new(HashTagQueryService), new(*HashTagQueryServiceImpl)),
)

// おすすめHashTag一覧参照
func (s *HashTagQueryServiceImpl) ShowRecommendList(areaID, subAreaID, subSubAreaID int) ([]*entity.HashTag, error) {
	return s.HashTagQueryRepository.FindRecommendList(areaID, subAreaID, subSubAreaID)
}
