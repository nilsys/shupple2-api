package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	PostCommandService interface {
		Store(post *entity.Post) error
	}

	PostCommandServiceImpl struct {
		Repository repository.PostCommandRepository
	}
)

var PostCommandServiceSet = wire.NewSet(
	wire.Struct(new(PostCommandServiceImpl), "*"),
	wire.Bind(new(PostCommandService), new(*PostCommandServiceImpl)),
)

func (r *PostCommandServiceImpl) Store(post *entity.Post) error {
	if err := r.Repository.Store(post); err != nil {
		return errors.Wrap(err, "failed to store post")
	}

	return nil
}
