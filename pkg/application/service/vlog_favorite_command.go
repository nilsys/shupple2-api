package service

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/service"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	VlogFavoriteCommandService interface {
		Store(user *entity.User, vlogID int) error
		Delete(user *entity.User, vlogID int) error
	}

	VlogFavoriteCommandServiceImpl struct {
		repository.VlogFavoriteCommandRepository
		repository.VlogFavoriteQueryRepository
		repository.VlogQueryRepository
		repository.VlogCommandRepository
		service.NoticeDomainService
		TransactionService
	}
)

var VlogFavoriteCommandServiceSet = wire.NewSet(
	wire.Struct(new(VlogFavoriteCommandServiceImpl), "*"),
	wire.Bind(new(VlogFavoriteCommandService), new(*VlogFavoriteCommandServiceImpl)),
)

func (s *VlogFavoriteCommandServiceImpl) Store(user *entity.User, vlogID int) error {
	isVlogExist, err := s.VlogQueryRepository.IsExist(vlogID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if !isVlogExist {
		return serror.New(nil, serror.CodeNotFound, "Not found")
	}

	existFavorite, err := s.VlogFavoriteQueryRepository.IsExist(user.ID, vlogID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if existFavorite {
		return serror.New(nil, serror.CodeInvalidParam, "already setted in the table")
	}

	favorite := entity.NewUserFavoriteVlog(user.ID, vlogID)

	return s.TransactionService.Do(func(c context.Context) error {
		if err := s.VlogFavoriteCommandRepository.Store(c, favorite); err != nil {
			return errors.Wrap(err, "failed to store favorite")
		}

		if err := s.VlogCommandRepository.IncrementFavoriteCount(c, vlogID); err != nil {
			return errors.Wrap(err, "failed to update vlog")
		}

		vlog, err := s.VlogQueryRepository.FindByID(vlogID)
		if err != nil {
			return errors.Wrap(err, "failed to find vlog by id")
		}

		return s.NoticeDomainService.FavoriteVlog(c, favorite, vlog, user)
	})
}

func (s *VlogFavoriteCommandServiceImpl) Delete(user *entity.User, vlogID int) error {
	isVlogExist, err := s.VlogQueryRepository.IsExist(vlogID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if !isVlogExist {
		return serror.New(nil, serror.CodeNotFound, "Not found")
	}

	existFavorite, err := s.VlogFavoriteQueryRepository.IsExist(user.ID, vlogID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if !existFavorite {
		return serror.New(nil, serror.CodeInvalidParam, "not setted in the table yet")
	}

	unfavorite := entity.NewUserFavoriteVlog(user.ID, vlogID)

	return s.TransactionService.Do(func(c context.Context) error {
		if err := s.VlogFavoriteCommandRepository.Delete(c, unfavorite); err != nil {
			return errors.Wrap(err, "failed to delete favorite")
		}

		if err := s.VlogCommandRepository.DecrementFavoriteCount(c, vlogID); err != nil {
			return errors.Wrap(err, "failed to update post")
		}

		return nil
	})
}
