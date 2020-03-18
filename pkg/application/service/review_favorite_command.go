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
	ReviewFavoriteCommandService interface {
		Store(user *entity.User, reviewID int) error
		Delete(user *entity.User, reviewID int) error
	}

	ReviewFavoriteCommandServiceImpl struct {
		ReviewFavoriteCommandRepository repository.ReviewFavoriteCommandRepository
		ReviewFavoriteQueryRepository   repository.ReviewFavoriteQueryRepository
		ReviewQueryRepository           repository.ReviewQueryRepository
		ReviewCommandRepository         repository.ReviewCommandRepository
		TransactionService
	}
)

var ReviewFavoriteCommandServiceSet = wire.NewSet(
	wire.Struct(new(ReviewFavoriteCommandServiceImpl), "*"),
	wire.Bind(new(ReviewFavoriteCommandService), new(*ReviewFavoriteCommandServiceImpl)),
)

func (r *ReviewFavoriteCommandServiceImpl) Store(user *entity.User, reviewID int) error {
	existReview, err := r.ReviewQueryRepository.IsExist(reviewID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if !existReview {
		return serror.New(nil, serror.CodeNotFound, "Not found")
	}

	existFavorite, err := r.ReviewFavoriteQueryRepository.IsExist(user.ID, reviewID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if existFavorite {
		return serror.New(nil, serror.CodeInvalidParam, "already set in table")
	}

	favorite := entity.NewUserFavoriteReview(user.ID, reviewID)

	return r.TransactionService.Do(func(c context.Context) error {
		if err := r.ReviewFavoriteCommandRepository.Store(c, favorite); err != nil {
			return errors.Wrap(err, "failed to store")
		}

		if err := r.ReviewCommandRepository.IncrementFavoriteCount(c, reviewID); err != nil {
			return errors.Wrap(err, "failed to update post")
		}

		return nil
	})
}

func (r *ReviewFavoriteCommandServiceImpl) Delete(user *entity.User, reviewID int) error {
	existReview, err := r.ReviewQueryRepository.IsExist(reviewID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if !existReview {
		return serror.New(nil, serror.CodeNotFound, "Not found")
	}

	existFavorite, err := r.ReviewFavoriteQueryRepository.IsExist(user.ID, reviewID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if !existFavorite {
		return serror.New(nil, serror.CodeInvalidParam, "not set in table yet")
	}

	unfavorite := entity.NewUserFavoriteReview(user.ID, reviewID)

	//TODO:lockを取る
	return r.TransactionService.Do(func(c context.Context) error {
		if err := r.ReviewFavoriteCommandRepository.Delete(c, unfavorite); err != nil {
			return errors.Wrap(err, "failed to delete")
		}

		if err := r.ReviewCommandRepository.DecrementFavoriteCount(c, reviewID); err != nil {
			return errors.Wrap(err, "failed to update post")
		}

		return nil
	})
}
