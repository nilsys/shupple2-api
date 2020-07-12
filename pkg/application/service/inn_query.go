package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	InnQueryService interface {
		ListInnByParams(areaID, subAreaID, subSubAreaID, touristSpotID int) (*entity.Inns, error)
	}

	InnQueryServiceImpl struct {
		repository.InnQueryRepository
		repository.MetasearchAreaQueryRepository
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
		metasearchAreas, err := s.MetasearchAreaQueryRepository.FindByAreaCategoryID(areaID, model.AreaCategoryTypeArea)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get area")
		}
		q.SetMetaserachID(metasearchAreas)
	}
	if subAreaID != 0 {
		metasearchAreas, err := s.MetasearchAreaQueryRepository.FindByAreaCategoryID(subAreaID, model.AreaCategoryTypeSubArea)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get sub area")
		}
		q.SetMetaserachID(metasearchAreas)
	}
	if subSubAreaID != 0 {
		metasearchAreas, err := s.MetasearchAreaQueryRepository.FindByAreaCategoryID(subSubAreaID, model.AreaCategoryTypeSubSubArea)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get sub sub area")
		}
		q.SetMetaserachID(metasearchAreas)
	}
	if touristSpotID != 0 {
		touristSpot, err := s.TouristSpotQueryRepository.FindByID(touristSpotID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get tourist spot")
		}
		if !touristSpot.Lat.Valid || !touristSpot.Lng.Valid {
			return nil, serror.New(nil, serror.CodeInvalidParam, "tourist_spot(id=%d) lacks lat or lng", touristSpotID)
		}
		q.Longitude = touristSpot.Lng.Float64
		q.Latitude = touristSpot.Lat.Float64
	}

	return s.InnQueryRepository.FindByParams(q)
}
