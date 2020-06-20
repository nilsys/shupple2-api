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
	ComicCommandService interface {
		ImportFromWordpressByID(wordpressComicID int) (*entity.Comic, error)
	}

	ComicCommandServiceImpl struct {
		ComicCommandRepository   repository.ComicCommandRepository
		WordpressQueryRepository repository.WordpressQueryRepository
		WordpressService
		TransactionService
	}
)

var ComicCommandServiceSet = wire.NewSet(
	wire.Struct(new(ComicCommandServiceImpl), "*"),
	wire.Bind(new(ComicCommandService), new(*ComicCommandServiceImpl)),
)

func (r *ComicCommandServiceImpl) ImportFromWordpressByID(id int) (*entity.Comic, error) {
	wpComic, err := r.WordpressQueryRepository.FindComicByID(id)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get wordpress comic(id=%d)", id)
	}

	if wpComic.Status != wordpress.StatusPublish {
		if err := r.ComicCommandRepository.DeleteByID(id); err != nil {
			return nil, errors.Wrapf(err, "failed to delete comic(id=%d)", id)
		}

		return nil, serror.New(nil, serror.CodeImportDeleted, "try to import deleted comic")
	}

	var comic *entity.Comic
	err = r.TransactionService.Do(func(c context.Context) error {
		if err := r.ComicCommandRepository.UndeleteByID(c, id); err != nil {
			return errors.Wrapf(err, "failed to undelete comic(id=%d)", id)
		}

		comic, err = r.ComicCommandRepository.Lock(c, id)
		if err != nil {
			if !serror.IsErrorCode(err, serror.CodeNotFound) {
				return errors.Wrap(err, "failed to get comic")
			}
			comic = &entity.Comic{}
		}

		if err := r.WordpressService.PatchComic(comic, wpComic); err != nil {
			return errors.Wrap(err, "failed  to patch comic")
		}

		if err := r.ComicCommandRepository.Store(c, comic); err != nil {
			return errors.Wrap(err, "failed to store comic")
		}

		return nil
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return comic, nil
}
