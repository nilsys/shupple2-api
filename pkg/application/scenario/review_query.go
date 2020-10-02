package scenario

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	ReviewQueryScenario interface {
		ListByParams(query *query.ShowReviewListQuery, ouser entity.OptionalUser) (*entity.ReviewDetailWithIsFavoriteList, *entity.UserRelationFlgMap, error)
		ListFeed(query *query.FindListPaginationQuery, user entity.User) (*entity.ReviewDetailWithIsFavoriteList, *entity.UserRelationFlgMap, error)
		ListFavorite(userID int, query *query.FindListPaginationQuery, ouser entity.OptionalUser) (*entity.ReviewDetailWithIsFavoriteList, *entity.UserRelationFlgMap, error)
		Show(id int, ouser entity.OptionalUser) (*entity.ReviewDetailWithIsFavorite, *entity.UserRelationFlgMap, error)
	}

	ReviewQueryScenarioImpl struct {
		service.ReviewQueryService
		service.UserQueryService
	}
)

var ReviewQueryScenarioSet = wire.NewSet(
	wire.Struct(new(ReviewQueryScenarioImpl), "*"),
	wire.Bind(new(ReviewQueryScenario), new(*ReviewQueryScenarioImpl)),
)

func (s *ReviewQueryScenarioImpl) ListByParams(query *query.ShowReviewListQuery, ouser entity.OptionalUser) (*entity.ReviewDetailWithIsFavoriteList, *entity.UserRelationFlgMap, error) {
	idRelationFlgMap := &entity.UserRelationFlgMap{}

	// 対象のReviewを取得
	list, err := s.ReviewQueryService.ListByParams(query, ouser)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed find reviews")
	}

	if ouser.IsAuthorized() {
		// 認証されている場合、Review.Userをfollowしているかフラグを取得
		idRelationFlgMap, err = s.UserQueryService.RelationFlgMaps(ouser.ID, list.UserIDs())
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed find is doing flg")
		}
	}

	return list, idRelationFlgMap, nil
}

func (s *ReviewQueryScenarioImpl) ListFeed(query *query.FindListPaginationQuery, user entity.User) (*entity.ReviewDetailWithIsFavoriteList, *entity.UserRelationFlgMap, error) {
	list, err := s.ReviewQueryService.ListFeed(user, query)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed find reviews")
	}

	// 認証されている場合、Review.Userをfollow, blockしているかフラグを取得
	idRelationFlgMap, err := s.UserQueryService.RelationFlgMaps(user.ID, list.UserIDs())
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed find is doing flg")
	}

	return list, idRelationFlgMap, nil
}

func (s *ReviewQueryScenarioImpl) ListFavorite(userID int, query *query.FindListPaginationQuery, ouser entity.OptionalUser) (*entity.ReviewDetailWithIsFavoriteList, *entity.UserRelationFlgMap, error) {
	idRelationFlgMap := &entity.UserRelationFlgMap{}

	list, err := s.ReviewQueryService.ListFavoriteReview(ouser, userID, query)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed find reviews")
	}

	if ouser.IsAuthorized() {
		// 認証されている場合、Review.Userをfollow, blockしているかフラグを取得
		idRelationFlgMap, err = s.UserQueryService.RelationFlgMaps(ouser.ID, list.UserIDs())
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed find is doing flg")
		}
	}

	return list, idRelationFlgMap, nil
}

func (s *ReviewQueryScenarioImpl) Show(id int, ouser entity.OptionalUser) (*entity.ReviewDetailWithIsFavorite, *entity.UserRelationFlgMap, error) {
	idRelationFlgMap := &entity.UserRelationFlgMap{}

	review, err := s.ReviewQueryService.ShowQueryReview(id, ouser)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed find review")
	}

	if ouser.Authenticated {
		// 認証されている場合、Review.Userをfollow, blockしているかフラグを取得
		idRelationFlgMap, err = s.UserQueryService.RelationFlgMaps(ouser.ID, []int{review.UserID})
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed find is doing flg")
		}
	}

	return review, idRelationFlgMap, nil
}
