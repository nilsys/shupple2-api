package service

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	FeatureQueryService interface {
		ShowQuery(id int) (*entity.FeatureDetailWithPosts, error)
		ShowList(query *query.FindListPaginationQuery) (*entity.FeatureList, error)
	}

	FeatureQueryServiceImpl struct {
		FeatureQueryRepository repository.FeatureQueryRepository
	}
)

var FeatureQueryServiceSet = wire.NewSet(
	wire.Struct(new(FeatureQueryServiceImpl), "*"),
	wire.Bind(new(FeatureQueryService), new(*FeatureQueryServiceImpl)),
)

// QueryFeature参照
func (s *FeatureQueryServiceImpl) ShowQuery(id int) (*entity.FeatureDetailWithPosts, error) {
	return s.FeatureQueryRepository.FindQueryFeatureByID(id)
}

// Feature一覧取得
func (s *FeatureQueryServiceImpl) ShowList(query *query.FindListPaginationQuery) (*entity.FeatureList, error) {
	return s.FeatureQueryRepository.FindList(query)
}
