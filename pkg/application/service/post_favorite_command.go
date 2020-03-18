package service

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	PostFavoriteCommandService interface {
		Store(user *entity.User, postID int) error
		Delete(user *entity.User, postID int) error
	}

	PostFavoriteCommandServiceImpl struct {
		PostFavoriteCommandRepository repository.PostFavoriteCommandRepository
		PostFavoriteQueryRepository   repository.PostFavoriteQueryRepository
		PostQueryRepository           repository.PostQueryRepository
		PostCommandRepository         repository.PostCommandRepository
		TransactionService
	}
)

var PostFavoriteCommandServiceSet = wire.NewSet(
	wire.Struct(new(PostFavoriteCommandServiceImpl), "*"),
	wire.Bind(new(PostFavoriteCommandService), new(*PostFavoriteCommandServiceImpl)),
)

func (r *PostFavoriteCommandServiceImpl) Store(user *entity.User, postID int) error {
	existPost, err := r.PostQueryRepository.IsExist(postID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if !existPost {
		return serror.New(nil, serror.CodeNotFound, "Not found")
	}

	existFavorite, err := r.PostFavoriteQueryRepository.IsExist(user.ID, postID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if existFavorite {
		return serror.New(nil, serror.CodeInvalidParam, "already set in table")
	}

	favorite := entity.NewUserFavoritePost(user.ID, postID)

	return r.TransactionService.Do(func(c context.Context) error {
		if err := r.PostFavoriteCommandRepository.Store(c, favorite); err != nil {
			return errors.Wrap(err, "failed to store favorite")
		}

		if err := r.PostCommandRepository.IncrementFavoriteCount(c, postID); err != nil {
			return errors.Wrap(err, "failed to update post")
		}

		return nil
	})
}

func (r *PostFavoriteCommandServiceImpl) Delete(user *entity.User, postID int) error {
	existPost, err := r.PostQueryRepository.IsExist(postID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if !existPost {
		return serror.New(nil, serror.CodeNotFound, "Not found")
	}

	existFavorite, err := r.PostFavoriteQueryRepository.IsExist(user.ID, postID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if !existFavorite {
		return serror.New(nil, serror.CodeInvalidParam, "not set in table yet")
	}

	unfavorite := entity.NewUserFavoritePost(user.ID, postID)

	//TODO:lockを取る
	return r.TransactionService.Do(func(c context.Context) error {
		if err := r.PostFavoriteCommandRepository.Delete(c, unfavorite); err != nil {
			return errors.Wrap(err, "failed to delete favorite")
		}

		if err := r.PostCommandRepository.DecrementFavoriteCount(c, postID); err != nil {
			return errors.Wrap(err, "failed to update post")
		}

		return nil
	})
}
