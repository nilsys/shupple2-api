package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
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
		WordpressService         WordpressService
	}
)

var ComicCommandServiceSet = wire.NewSet(
	wire.Struct(new(ComicCommandServiceImpl), "*"),
	wire.Bind(new(ComicCommandService), new(*ComicCommandServiceImpl)),
)

func (r *ComicCommandServiceImpl) ImportFromWordpressByID(id int) (*entity.Comic, error) {
	wpComics, err := r.WordpressQueryRepository.FindComicsByIDs([]int{id})
	if err != nil || len(wpComics) == 0 {
		return nil, serror.NewResourcesNotFoundError(err, "wordpress comic(id=%d)", id)
	}

	comic, err := r.WordpressService.ConvertComic(wpComics[0])
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert comic")
	}

	if err := r.ComicCommandRepository.Store(comic); err != nil {
		return nil, errors.Wrap(err, "failed to store comic")
	}

	return comic, nil
}
