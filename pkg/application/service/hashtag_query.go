package service

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	// HashTag参照系サービス
	HashtagQueryService interface {
		ShowRecommendList(areID, subAreaID, subSubAreaID int) ([]*entity.Hashtag, error)
	}

	// HashTag参照系サービス実装
	HashtagQueryServiceImpl struct {
		repository.HashtagQueryRepository
	}
)

var HashtagQueryServiceSet = wire.NewSet(
	wire.Struct(new(HashtagQueryServiceImpl), "*"),
	wire.Bind(new(HashtagQueryService), new(*HashtagQueryServiceImpl)),
)

// おすすめHashTag一覧参照
func (s *HashtagQueryServiceImpl) ShowRecommendList(areaID, subAreaID, subSubAreaID int) ([]*entity.Hashtag, error) {
	return s.HashtagQueryRepository.FindRecommendList(areaID, subAreaID, subSubAreaID)
}
