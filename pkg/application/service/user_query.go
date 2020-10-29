package service

import (
	"github.com/google/wire"
	"github.com/uma-co82/shupple2-api/pkg/domain/entity"
	"github.com/uma-co82/shupple2-api/pkg/domain/model/serror"
	"github.com/uma-co82/shupple2-api/pkg/domain/repository"
)

type (
	UserQueryService interface {
		Show(user *entity.UserTiny) (*entity.User, error)
		ShowByID(id int) (*entity.User, error)
		ShowMatchingUser(user *entity.UserTiny) (*entity.User, error)
		ListNotReviewMainMatchingMatchingUser(user *entity.UserTiny) ([]*entity.User, error)
		ListMainMatchingUser(user *entity.UserTiny) ([]*entity.User, error)
	}

	UserQueryServiceImpl struct {
		repository.UserQueryRepository
	}
)

var UserQueryServiceSet = wire.NewSet(
	wire.Struct(new(UserQueryServiceImpl), "*"),
	wire.Bind(new(UserQueryService), new(*UserQueryServiceImpl)),
)

func (s *UserQueryServiceImpl) Show(user *entity.UserTiny) (*entity.User, error) {
	return s.UserQueryRepository.FindByID(user.ID)
}

func (s *UserQueryServiceImpl) ShowByID(id int) (*entity.User, error) {
	return s.UserQueryRepository.FindByID(id)
}

/*
	マッチングしているユーザーを取得
	マッチングしていない場合はCodeNotMatching
*/
func (s *UserQueryServiceImpl) ShowMatchingUser(user *entity.UserTiny) (*entity.User, error) {
	if !user.IsMatching() {
		return nil, serror.New(nil, serror.CodeNotMatching, "not matching")
	}
	return s.UserQueryRepository.FindMatchingUserByID(user.ID)
}

/*
	マッチング後の評価をしていないユーザー一覧
*/
func (s *UserQueryServiceImpl) ListNotReviewMainMatchingMatchingUser(user *entity.UserTiny) ([]*entity.User, error) {
	return s.UserQueryRepository.FindNotMainMatchingReviewMatchingUsersByID(user.ID)
}

/*
	マッチング後お互いがいいねしたユーザー一覧
*/
func (s *UserQueryServiceImpl) ListMainMatchingUser(user *entity.UserTiny) ([]*entity.User, error) {
	return s.UserQueryRepository.FindMainMatchingUserByID(user.ID)
}
