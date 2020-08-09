package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/factory"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	VlogQueryService interface {
		Show(id int, ouser *entity.OptionalUser) (*entity.VlogDetail, *entity.TouristSpotReviewCountList, error)
		ListByParams(query *query.FindVlogListQuery, ouser *entity.OptionalUser) (*entity.VlogList, error)
	}

	VlogQueryServiceImpl struct {
		VlogQueryRepository        repository.VlogQueryRepository
		TouristSpotQueryRepository repository.TouristSpotQueryRepository
		CategoryIDMapFactory       factory.CategoryIDMapFactory
	}
)

var VlogQueryServiceSet = wire.NewSet(
	wire.Struct(new(VlogQueryServiceImpl), "*"),
	wire.Bind(new(VlogQueryService), new(*VlogQueryServiceImpl)),
)

// WARN:
// TODO: review_countをtourist_spotテーブルに追加して、Review投稿時にIncrementする様にする、その際にscriptを書いて既存のReviewの数を含める
func (s *VlogQueryServiceImpl) Show(id int, ouser *entity.OptionalUser) (*entity.VlogDetail, *entity.TouristSpotReviewCountList, error) {
	var vlog *entity.VlogDetail
	var err error

	if ouser.Authenticated {
		vlog, err = s.VlogQueryRepository.FindDetailWithIsFavoriteByID(id, ouser.ID)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed find vlog by id")
		}
	} else {
		vlog, err = s.VlogQueryRepository.FindDetailByID(id)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed find vlog by id")
		}
	}

	touristSpotReviewCountList, err := s.TouristSpotQueryRepository.FindReviewCountByIDs(vlog.TouristSpotIDs())
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed find tourist_spot and review_count")
	}
	return vlog, touristSpotReviewCountList, nil
}

func (s *VlogQueryServiceImpl) ListByParams(query *query.FindVlogListQuery, ouser *entity.OptionalUser) (*entity.VlogList, error) {
	if ouser.Authenticated {
		return s.VlogQueryRepository.FindWithIsFavoriteListByParams(query, ouser.ID)
	}
	return s.VlogQueryRepository.FindListByParams(query)
}
