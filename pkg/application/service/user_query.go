package service

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	UserQueryService interface {
		ShowUserRanking(query *query.FindUserRankingListQuery) ([]*entity.QueryRankingUser, error)
		ListFollowee(query *query.FindFollowUser) ([]*entity.User, error)
		ListFollower(query *query.FindFollowUser) ([]*entity.User, error)
	}

	UserQueryServiceImpl struct {
		repository.UserQueryRepository
	}
)

var UserQueryServiceSet = wire.NewSet(
	wire.Struct(new(UserQueryServiceImpl), "*"),
	wire.Bind(new(UserQueryService), new(*UserQueryServiceImpl)),
)

func (s *UserQueryServiceImpl) ShowUserRanking(query *query.FindUserRankingListQuery) ([]*entity.QueryRankingUser, error) {
	return s.UserQueryRepository.FindUserRankingListByParams(query)
}

func (s *UserQueryServiceImpl) ListFollowee(query *query.FindFollowUser) ([]*entity.User, error) {
	return s.UserQueryRepository.FindFolloweeByID(query)
}

func (s *UserQueryServiceImpl) ListFollower(query *query.FindFollowUser) ([]*entity.User, error) {
	return s.UserQueryRepository.FindFollowerByID(query)
}
