package scenario

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	ReviewQueryScenario interface {
		ListByParams(query *query.ShowReviewListQuery, ouser entity.OptionalUser) (*entity.ReviewDetailWithIsFavoriteList, map[int]bool, error)
		ListFeed(query *query.FindListPaginationQuery, user entity.User) (*entity.ReviewDetailWithIsFavoriteList, map[int]bool, error)
		ListFavorite(userID int, query *query.FindListPaginationQuery, ouser entity.OptionalUser) (*entity.ReviewDetailWithIsFavoriteList, map[int]bool, error)
		Show(id int, ouser entity.OptionalUser) (*entity.ReviewDetailWithIsFavorite, map[int]bool, error)
	}

	ReviewQueryScenarioImpl struct {
		service.ReviewQueryService
		repository.UserQueryRepository
	}
)

var ReviewQueryScenarioSet = wire.NewSet(
	wire.Struct(new(ReviewQueryScenarioImpl), "*"),
	wire.Bind(new(ReviewQueryScenario), new(*ReviewQueryScenarioImpl)),
)

func (s *ReviewQueryScenarioImpl) ListByParams(query *query.ShowReviewListQuery, ouser entity.OptionalUser) (*entity.ReviewDetailWithIsFavoriteList, map[int]bool, error) {
	var idIsFollowMap map[int]bool

	// 対象のReviewを取得
	list, err := s.ReviewQueryService.ListByParams(query, ouser)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed find reviews")
	}

	if ouser.IsAuthorized() {
		// 認証されている場合、Review.Userをfollowしているかフラグを取得
		idIsFollowMap, err = s.UserQueryRepository.IsFollowing(ouser.ID, list.UserIDs())
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed find user_following")
		}
	}

	return list, idIsFollowMap, nil
}

func (s *ReviewQueryScenarioImpl) ListFeed(query *query.FindListPaginationQuery, user entity.User) (*entity.ReviewDetailWithIsFavoriteList, map[int]bool, error) {
	var idIsFollowMap map[int]bool

	list, err := s.ReviewQueryService.ListFeed(user, query)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed find reviews")
	}

	// 認証されている場合、Review.Userをfollowしているかフラグを取得
	idIsFollowMap, err = s.UserQueryRepository.IsFollowing(user.ID, list.UserIDs())
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed find user_following")
	}

	return list, idIsFollowMap, nil
}

func (s *ReviewQueryScenarioImpl) ListFavorite(userID int, query *query.FindListPaginationQuery, ouser entity.OptionalUser) (*entity.ReviewDetailWithIsFavoriteList, map[int]bool, error) {
	var idIsFollowMap map[int]bool

	list, err := s.ReviewQueryService.ListFavoriteReview(ouser, userID, query)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed find reviews")
	}

	if ouser.IsAuthorized() {
		// 認証されている場合、Review.Userをfollowしているかフラグを取得
		idIsFollowMap, err = s.UserQueryRepository.IsFollowing(ouser.ID, list.UserIDs())
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed find user_following")
		}
	}

	return list, idIsFollowMap, nil
}

func (s *ReviewQueryScenarioImpl) Show(id int, ouser entity.OptionalUser) (*entity.ReviewDetailWithIsFavorite, map[int]bool, error) {
	var idIsFollowMap map[int]bool

	review, err := s.ReviewQueryService.ShowQueryReview(id, ouser)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed find review")
	}

	if ouser.Authenticated {
		// 認証されている場合、Review.Userをfollowしているかフラグを取得
		idIsFollowMap, err = s.UserQueryRepository.IsFollowing(ouser.ID, []int{review.UserID})
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed find user_following")
		}
	}

	return review, idIsFollowMap, nil
}
