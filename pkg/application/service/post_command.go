package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	PostCommandService interface {
		ImportFromWordpressByID(wordpressPostID int) (*entity.Post, error)
	}

	PostCommandServiceImpl struct {
		PostCommandRepository    repository.PostCommandRepository
		WordpressQueryRepository repository.WordpressQueryRepository
		WordpressService         WordpressService
	}
)

var PostCommandServiceSet = wire.NewSet(
	wire.Struct(new(PostCommandServiceImpl), "*"),
	wire.Bind(new(PostCommandService), new(*PostCommandServiceImpl)),
)

func (r *PostCommandServiceImpl) ImportFromWordpressByID(id int) (*entity.Post, error) {
	wpPosts, err := r.WordpressQueryRepository.FindPostsByIDs([]int{id})
	if err != nil || len(wpPosts) == 0 {
		return nil, serror.NewResourcesNotFoundError(err, "wordpress post(id=%d)", id)
	}

	post, err := r.WordpressService.ConvertPost(wpPosts[0])
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert post")
	}

	if err := r.PostCommandRepository.Store(post); err != nil {
		return nil, errors.Wrap(err, "failed to store post")
	}

	return post, nil
}
