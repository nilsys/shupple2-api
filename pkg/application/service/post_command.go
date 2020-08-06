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
		Store(c context.Context, post *entity.Post) error
	}

	PostCommandServiceImpl struct {
		repository.PostCommandRepository
		repository.HashtagCommandRepository
		repository.WordpressQueryRepository
		repository.CfProjectCommandRepository
		WordpressService
		TransactionService
	}
)

var PostCommandServiceSet = wire.NewSet(
	wire.Struct(new(PostCommandServiceImpl), "*"),
	wire.Bind(new(PostCommandService), new(*PostCommandServiceImpl)),
)

func (r *PostCommandServiceImpl) ImportFromWordpressByID(id int) (*entity.Post, error) {
	wpPost, err := r.WordpressQueryRepository.FindPostByID(id)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get wordpress post(id=%d)", id)
	}

	if wpPost.Status != wordpress.StatusPublish {
		if err := r.PostCommandRepository.DeleteByID(context.TODO(), id); err != nil {
			return nil, errors.Wrapf(err, "failed to delete post(id=%d)", id)
		}

		return nil, serror.New(nil, serror.CodeImportDeleted, "try to import deleted post")
	}

	var post *entity.Post
	err = r.TransactionService.Do(func(c context.Context) error {
		if err := r.PostCommandRepository.UndeleteByID(c, id); err != nil {
			return errors.Wrapf(err, "failed to undelete post(id=%d)", id)
		}

		post, err = r.PostCommandRepository.Lock(c, id)
		if err != nil {
			if !serror.IsErrorCode(err, serror.CodeNotFound) {
				return errors.Wrap(err, "failed to get post")
			}
			post = &entity.Post{}
		}

		if err := r.WordpressService.PatchPost(post, wpPost); err != nil {
			return errors.Wrap(err, "failed to patch post")
		}

		if err := r.Store(c, post); err != nil {
			return errors.Wrap(err, "failed to store post")
		}

		if post.CfProjectID.Valid {
			if err := r.CfProjectCommandRepository.UpdateLatestPostID(c, int(post.CfProjectID.Int64), post.ID); err != nil {
				return errors.Wrap(err, "failed to update cf_project.latest_post_id, is_sent_new_post_email")
			}
		}

		return nil
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return post, nil
}

func (r *PostCommandServiceImpl) Store(c context.Context, post *entity.Post) error {
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
}
