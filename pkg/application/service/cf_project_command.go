package service

import (
	"context"

	"github.com/google/wire"

	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CfProjectCommandService interface {
		ImportFromWordpressByID(id int) error
		Favorite(user *entity.User, projectID int) error
		Unfavorite(user *entity.User, projectID int) error
	}

	CfProjectCommandServiceImpl struct {
		repository.CfProjectCommandRepository
		repository.WordpressQueryRepository
		WordpressService
		TransactionService
	}
)

var CfProjectCommandServiceSet = wire.NewSet(
	wire.Struct(new(CfProjectCommandServiceImpl), "*"),
	wire.Bind(new(CfProjectCommandService), new(*CfProjectCommandServiceImpl)),
)

func (s *CfProjectCommandServiceImpl) ImportFromWordpressByID(id int) error {
	wpCfProject, err := s.WordpressQueryRepository.FindCfProjectByID(id)
	if err != nil {
		return errors.Wrapf(err, "failed to get wordpress cfProject(id=%d)", id)
	}

	if wpCfProject.Status != wordpress.StatusPublish {
		if err := s.CfProjectCommandRepository.DeleteByID(id); err != nil {
			return errors.Wrapf(err, "failed to delete cfProject(id=%d)", id)
		}

		return serror.New(nil, serror.CodeImportDeleted, "try to import deleted cfProject")
	}

	cfProject, err := s.WordpressService.NewCfProject(wpCfProject)
	if err != nil {
		return errors.Wrap(err, "failed  to initialize cfProject")
	}

	if err := s.CfProjectCommandRepository.Store(cfProject); err != nil {
		return errors.Wrap(err, "failed to store cfProject")
	}

	return nil
}

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
