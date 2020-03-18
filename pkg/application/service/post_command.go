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
	PostCommandService interface {
		ImportFromWordpressByID(wordpressPostID int) (*entity.Post, error)
		Store(post *entity.Post) error
	}

	PostCommandServiceImpl struct {
		repository.PostCommandRepository
		repository.HashtagCommandRepository
		repository.WordpressQueryRepository
		WordpressService
		TransactionService
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

	if wpPosts[0].Status != wordpress.StatusPublish {
		if err := r.PostCommandRepository.DeleteByID(context.TODO(), id); err != nil {
			return nil, errors.Wrapf(err, "failed to delete post(id=%d)", id)
		}

		return nil, serror.New(nil, serror.CodeImportDeleted, "try to import deleted post")
	}

	post, err := r.WordpressService.ConvertPost(wpPosts[0])
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert post")
	}

	if err := r.Store(post); err != nil {
		return nil, errors.Wrap(err, "failed to store post")
	}

	return post, nil
}

func (r *PostCommandServiceImpl) Store(post *entity.Post) error {
	return r.TransactionService.Do(func(c context.Context) error {
		if err := r.HashtagCommandRepository.DecrementPostCountByPostID(c, post.ID); err != nil {
			return errors.Wrap(err, "failed to decrement post_count")
		}

		if err := r.PostCommandRepository.Store(c, post); err != nil {
			return errors.Wrap(err, "failed to store post")
		}

		if err := r.HashtagCommandRepository.IncrementPostCountByPostID(c, post.ID); err != nil {
			return errors.Wrap(err, "failed to increment post_count")
		}

		return nil
	})
}
