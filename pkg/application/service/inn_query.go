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
		repository.AreaCategoryQueryRepository
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
		area, err := s.AreaCategoryQueryRepository.FindByIDAndType(areaID, model.AreaCategoryTypeArea)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get area")
		}
		q.SetMetaserachID(area.MetasearchAreaID, area.MetasearchSubAreaID, area.MetasearchSubSubAreaID)
	}
	if subAreaID != 0 {
		subArea, err := s.AreaCategoryQueryRepository.FindByIDAndType(subAreaID, model.AreaCategoryTypeSubArea)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get sub area")
		}
		q.SetMetaserachID(subArea.MetasearchAreaID, subArea.MetasearchSubAreaID, subArea.MetasearchSubSubAreaID)
	}
	if subSubAreaID != 0 {
		subSubArea, err := s.AreaCategoryQueryRepository.FindByIDAndType(subSubAreaID, model.AreaCategoryTypeSubSubArea)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get sub sub area")
		}
		q.SetMetaserachID(subSubArea.MetasearchAreaID, subSubArea.MetasearchSubAreaID, subSubArea.MetasearchSubSubAreaID)
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
