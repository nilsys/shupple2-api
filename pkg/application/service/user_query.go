package service

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	UserQueryService interface {
		Show(id int) (*entity.UserDetailWithMediaCount, error)
		ShowUserRanking(query *query.FindUserRankingListQuery) ([]*entity.UserDetail, error)
		ListFollowing(query *query.FindFollowUser) ([]*entity.User, error)
		ListFollowed(query *query.FindFollowUser) ([]*entity.User, error)
		ListFavoritePostUser(postID int, user *entity.OptionalUser, query *query.FindListPaginationQuery) ([]*entity.User, error)
		ListFavoriteReviewUser(reviewID int, user *entity.OptionalUser, query *query.FindListPaginationQuery) ([]*entity.User, error)
	}

	UserQueryServiceImpl struct {
		repository.UserQueryRepository
	}
)

var UserQueryServiceSet = wire.NewSet(
	wire.Struct(new(UserQueryServiceImpl), "*"),
	wire.Bind(new(UserQueryService), new(*UserQueryServiceImpl)),
)

func (s *UserQueryServiceImpl) Show(id int) (*entity.UserDetailWithMediaCount, error) {
	return s.UserQueryRepository.FindUserDetailWithCountByID(id)
}

func (s *UserQueryServiceImpl) ShowUserRanking(query *query.FindUserRankingListQuery) ([]*entity.UserDetail, error) {
	return s.UserQueryRepository.FindUserRankingListByParams(query)
}

func (s *UserQueryServiceImpl) ListFollowing(query *query.FindFollowUser) ([]*entity.User, error) {
	return s.UserQueryRepository.FindFollowingByID(query)
}

func (s *UserQueryServiceImpl) ListFollowed(query *query.FindFollowUser) ([]*entity.User, error) {
	return s.UserQueryRepository.FindFollowedByID(query)
}

func (s *UserQueryServiceImpl) ListFavoritePostUser(postID int, user *entity.OptionalUser, query *query.FindListPaginationQuery) ([]*entity.User, error) {
	if user.IsAuthorized() {
		return s.UserQueryRepository.FindFavoritePostUserByUserID(postID, user.ID, query)
	}

	return s.UserQueryRepository.FindFavoritePostUser(postID, query)
}

func (s *UserQueryServiceImpl) ListFavoriteReviewUser(reviewID int, user *entity.OptionalUser, query *query.FindListPaginationQuery) ([]*entity.User, error) {
	if user.IsAuthorized() {
		return s.UserQueryRepository.FindFavoriteReviewUserByUserID(reviewID, user.ID, query)
	}

	return s.UserQueryRepository.FindFavoriteReviewUser(reviewID, query)
}
