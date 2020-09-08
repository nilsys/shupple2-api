package service

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/service"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	ComicFavoriteCommandService interface {
		Store(user *entity.User, comicID int) error
		Delete(user *entity.User, comicID int) error
	}

	ComicFavoriteCommandServiceImpl struct {
		repository.ComicFavoriteCommandRepository
		repository.ComicFavoriteQueryRepository
		repository.ComicCommandRepository
		repository.ComicQueryRepository
		service.NoticeDomainService
		TransactionService
	}
)

var ComicFavoriteCommandServiceSet = wire.NewSet(
	wire.Struct(new(ComicFavoriteCommandServiceImpl), "*"),
	wire.Bind(new(ComicFavoriteCommandService), new(*ComicFavoriteCommandServiceImpl)),
)

func (s *ComicFavoriteCommandServiceImpl) Store(user *entity.User, comicID int) error {
	isExistComic, err := s.ComicQueryRepository.IsExist(comicID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if !isExistComic {
		return serror.New(nil, serror.CodeNotFound, "Not found")
	}

	isExistFavorite, err := s.ComicFavoriteQueryRepository.IsExist(user.ID, comicID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if isExistFavorite {
		return serror.New(nil, serror.CodeInvalidParam, "already setted in the table")
	}

	favorite := entity.NewUserFavoriteComic(user.ID, comicID)

	return s.TransactionService.Do(func(c context.Context) error {
		if err := s.ComicFavoriteCommandRepository.Store(c, favorite); err != nil {
			return errors.Wrap(err, "failed to store favorite")
		}

		if err := s.ComicCommandRepository.IncrementFavoriteCount(c, comicID); err != nil {
			return errors.Wrap(err, "failed to update post")
		}

		queryComic, err := s.ComicQueryRepository.FindByID(comicID)
		if err != nil {
			return errors.Wrap(err, "failed to find post by id")
		}

		return s.NoticeDomainService.FavoriteComic(c, favorite, &queryComic.Comic, user)
	})
}

func (s *ComicFavoriteCommandServiceImpl) Delete(user *entity.User, comicID int) error {
	isExistComic, err := s.ComicQueryRepository.IsExist(comicID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if !isExistComic {
		return serror.New(nil, serror.CodeNotFound, "Not found")
	}

	isExistFavorite, err := s.ComicFavoriteQueryRepository.IsExist(user.ID, comicID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if !isExistFavorite {
		return serror.New(nil, serror.CodeInvalidParam, "not setted in the table yet")
	}

	unfavorite := entity.NewUserFavoriteComic(user.ID, comicID)

	return s.TransactionService.Do(func(c context.Context) error {
		if err := s.ComicFavoriteCommandRepository.Delete(c, unfavorite); err != nil {
			return errors.Wrap(err, "failed to delete favorite")
		}

		if err := s.ComicCommandRepository.DecrementFavoriteCount(c, comicID); err != nil {
			return errors.Wrap(err, "failed to update post")
		}

		return nil
	})
}
