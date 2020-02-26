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
