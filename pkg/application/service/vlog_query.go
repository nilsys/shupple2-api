package service

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/factory"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	VlogQueryService interface {
		Show(id int, ouser *entity.OptionalUser) (*entity.VlogDetail, error)
		ListByParams(query *query.FindVlogListQuery, ouser *entity.OptionalUser) (*entity.VlogList, error)
	}

	VlogQueryServiceImpl struct {
		VlogQueryRepository  repository.VlogQueryRepository
		CategoryIDMapFactory factory.CategoryIDMapFactory
	}
)

var VlogQueryServiceSet = wire.NewSet(
	wire.Struct(new(VlogQueryServiceImpl), "*"),
	wire.Bind(new(VlogQueryService), new(*VlogQueryServiceImpl)),
)

func (s *VlogQueryServiceImpl) Show(id int, ouser *entity.OptionalUser) (*entity.VlogDetail, error) {
	if ouser.Authenticated {
		return s.VlogQueryRepository.FindDetailWithIsFavoriteByID(id, ouser.ID)
	}
	return s.VlogQueryRepository.FindDetailByID(id)
}

func (s *VlogQueryServiceImpl) ListByParams(query *query.FindVlogListQuery, ouser *entity.OptionalUser) (*entity.VlogList, error) {
	if ouser.Authenticated {
		return s.VlogQueryRepository.FindWithIsFavoriteListByParams(query, ouser.ID)
	}
	return s.VlogQueryRepository.FindListByParams(query)
}
