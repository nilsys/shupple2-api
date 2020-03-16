package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	VlogQueryService interface {
		Show(id int) (*entity.VlogDetailWithTouristSpots, error)
		ShowListByParams(query *query.FindVlogListQuery) ([]*entity.VlogDetail, error)
	}

	VlogQueryServiceImpl struct {
		VlogQueryRepository repository.VlogQueryRepository
	}
)

var VlogQueryServiceSet = wire.NewSet(
	wire.Struct(new(VlogQueryServiceImpl), "*"),
	wire.Bind(new(VlogQueryService), new(*VlogQueryServiceImpl)),
)

func (s *VlogQueryServiceImpl) Show(id int) (*entity.VlogDetailWithTouristSpots, error) {
	return s.VlogQueryRepository.FindWithTouristSpotsByID(id)
}

func (s *VlogQueryServiceImpl) ShowListByParams(query *query.FindVlogListQuery) ([]*entity.VlogDetail, error) {
	vlogs, err := s.VlogQueryRepository.FindListByParams(query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find vlogs")
	}

	return vlogs, nil
}
