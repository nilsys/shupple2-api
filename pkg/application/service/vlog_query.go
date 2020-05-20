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
		Show(id int) (*entity.VlogDetail, error)
		ListByParams(query *query.FindVlogListQuery) (*entity.VlogList, error)
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

func (s *VlogQueryServiceImpl) Show(id int) (*entity.VlogDetail, error) {
	return s.VlogQueryRepository.FindDetailByID(id)
}

func (s *VlogQueryServiceImpl) ListByParams(query *query.FindVlogListQuery) (*entity.VlogList, error) {
	return s.VlogQueryRepository.FindListByParams(query)
}
