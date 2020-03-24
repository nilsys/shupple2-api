package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	InnQueryService interface {
		ListInnByParams(areaID, subAreaID, subSubAreaID, touristSpotID int) (*entity.Inns, error)
	}

	InnQueryServiceImpl struct {
		repository.InnQueryRepository
		repository.CategoryQueryRepository
		repository.TouristSpotQueryRepository
	}
)

var InnQueryServiceSet = wire.NewSet(
	wire.Struct(new(InnQueryServiceImpl), "*"),
	wire.Bind(new(InnQueryService), new(*InnQueryServiceImpl)),
)

// TODO: スマートじゃない
func (s *InnQueryServiceImpl) ListInnByParams(areaID, subAreaID, subSubAreaID, touristSpotID int) (*entity.Inns, error) {
	q := &query.FindInn{}

	if areaID != 0 {
		area, err := s.CategoryQueryRepository.FindByID(areaID)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		q.SetMetaserachID(area.MetasearchAreaID, area.MetasearchSubAreaID, area.MetasearchSubSubAreaID)
	}
	if subAreaID != 0 {
		subArea, err := s.CategoryQueryRepository.FindByID(subAreaID)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		q.SetMetaserachID(subArea.MetasearchAreaID, subArea.MetasearchSubAreaID, subArea.MetasearchSubSubAreaID)
	}
	if subSubAreaID != 0 {
		subSubArea, err := s.CategoryQueryRepository.FindByID(subSubAreaID)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		q.SetMetaserachID(subSubArea.MetasearchAreaID, subSubArea.MetasearchSubAreaID, subSubArea.MetasearchSubSubAreaID)
	}
	if touristSpotID != 0 {
		touristSpot, err := s.TouristSpotQueryRepository.FindByID(touristSpotID)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		q.Longitude = touristSpot.Lng
		q.Latitude = touristSpot.Lat
	}

	return s.InnQueryRepository.FindByParams(q)
}
