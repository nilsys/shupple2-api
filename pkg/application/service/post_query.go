package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	PostQueryService interface {
		FindByID(id int) (*entity.Post, error)
	}

	PostQueryServiceImpl struct {
		Repository repository.PostQueryRepository
	}
)

var PostQueryServiceSet = wire.NewSet(
	wire.Struct(new(PostQueryServiceImpl), "*"),
	wire.Bind(new(PostQueryService), new(*PostQueryServiceImpl)),
)

func (r *PostQueryServiceImpl) FindByID(id int) (*entity.Post, error) {
	post, err := r.Repository.FindByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get post")
	}

	return post, nil
}
