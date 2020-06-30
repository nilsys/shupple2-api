package service

import (
	"context"

	"github.com/google/wire"

	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CfProjectCommandService interface {
		Favorite(user *entity.User, projectID int) error
		Unfavorite(user *entity.User, projectID int) error
	}

	CfProjectCommandServiceImpl struct {
		repository.CfProjectCommandRepository
		TransactionService
	}
)

var CfProjectCommandServiceSet = wire.NewSet(
	wire.Struct(new(CfProjectCommandServiceImpl), "*"),
	wire.Bind(new(CfProjectCommandService), new(*CfProjectCommandServiceImpl)),
)

func (s *CfProjectCommandServiceImpl) Favorite(user *entity.User, projectID int) error {
	return s.TransactionService.Do(func(c context.Context) error {
		if err := s.CfProjectCommandRepository.StoreUserFavoriteCfProject(c, entity.NewUserFavoriteCfProject(user.ID, projectID)); err != nil {
			return errors.Wrap(err, "failed store user_favorite_cf_project")
		}

		if err := s.CfProjectCommandRepository.IncrementFavoriteCountByID(c, projectID); err != nil {
			return errors.Wrap(err, "failed increment favorite_count")
		}

		// MEMO: 通知を足すならここで
		return nil
	})
}

func (s *CfProjectCommandServiceImpl) Unfavorite(user *entity.User, projectID int) error {
	return s.TransactionService.Do(func(c context.Context) error {
		if err := s.CfProjectCommandRepository.DeleteUserFavoriteCfProject(c, entity.NewUserFavoriteCfProject(user.ID, projectID)); err != nil {
			return errors.Wrap(err, "failed delete user_favorite_cf_project")
		}

		if err := s.CfProjectCommandRepository.DecrementFavoriteCountByID(c, projectID); err != nil {
			return errors.Wrap(err, "failed decrement favorite_count")
		}

		return nil
	})
}
