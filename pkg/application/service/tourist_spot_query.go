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
	// TouristSpot参照系サービス
	TouristSpotQueryService interface {
		Show(id int) (*entity.TouristSpotDetail, error)
		List(query *query.FindTouristSpotListQuery) (*entity.TouristSpotList, error)
		ListRecommend(query *query.FindRecommendTouristSpotListQuery) (*entity.TouristSpotList, error)
	}

	// TouristSpot参照系サービス実装
	TouristSpotQueryServiceImpl struct {
		repository.TouristSpotQueryRepository
		repository.ReviewQueryRepository
		factory.CategoryIDMapFactory
	}
)

var TouristSpotQueryServiceSet = wire.NewSet(
	wire.Struct(new(TouristSpotQueryServiceImpl), "*"),
	wire.Bind(new(TouristSpotQueryService), new(*TouristSpotQueryServiceImpl)),
)

// TouristSpot参照
func (s *TouristSpotQueryServiceImpl) Show(id int) (*entity.TouristSpotDetail, error) {
	touristSpot, err := s.TouristSpotQueryRepository.FindDetailByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed find tourist_spot")
	}

	if err := s.setAlternativeThumbnail(&touristSpot.TouristSpotTiny); err != nil {
		return nil, errors.Wrap(err, "failed set alternative thumbnail")
	}

	return touristSpot, nil
}

// TouristSpot一覧参照
func (s *TouristSpotQueryServiceImpl) List(query *query.FindTouristSpotListQuery) (*entity.TouristSpotList, error) {
	touristSpots, err := s.TouristSpotQueryRepository.FindListByParams(query)
	if err != nil {
		return nil, errors.Wrap(err, "failed find tourist_spot list")
	}

	if err := s.setAlternativeThumbnailForList(touristSpots); err != nil {
		return nil, errors.Wrap(err, "failed set alternative thumbnail")
	}

	return touristSpots, nil
}

// おすすめTouristSpot一覧取得
func (s *TouristSpotQueryServiceImpl) ListRecommend(query *query.FindRecommendTouristSpotListQuery) (*entity.TouristSpotList, error) {
	touristSpots, err := s.TouristSpotQueryRepository.FindRecommendListByParams(query)
	if err != nil {
		return nil, errors.Wrap(err, "failed find tourist_spot list")
	}

	if err := s.setAlternativeThumbnailForList(touristSpots); err != nil {
		return nil, errors.Wrap(err, "failed set alternative thumbnail")
	}

	return touristSpots, nil
}

// サムネイルを持たない場合,代替サムネイルをセットする
func (s *TouristSpotQueryServiceImpl) setAlternativeThumbnail(touristSpot *entity.TouristSpotTiny) error {
	if touristSpot.HasThumbnail() {
		return nil
	}

	reviews, err := s.ReviewQueryRepository.FindLatestHasMediaReviewByTouristSpotIDs([]int{touristSpot.ID})
	if err != nil {
		return errors.Wrap(err, "failed find tourist_spot latest review")
	}

	touristSpot.Thumbnail = reviews.TouristSpotAlternativeImage(touristSpot.ID)

	return nil
}

// サムネイルを持たない場合,代替サムネイルをセットする
func (s *TouristSpotQueryServiceImpl) setAlternativeThumbnailForList(touristSpots *entity.TouristSpotList) error {
	reviews, err := s.ReviewQueryRepository.FindLatestHasMediaReviewByTouristSpotIDs(touristSpots.IDs())
	if err != nil {
		return errors.Wrap(err, "failed find tourist_spot latest review")
	}

	for _, tiny := range touristSpots.TouristSpots {
		if tiny.HasThumbnail() {
			continue
		}
		tiny.Thumbnail = reviews.TouristSpotAlternativeImage(tiny.ID)
	}

	return nil
}
