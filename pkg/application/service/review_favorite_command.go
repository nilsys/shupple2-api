package service

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/service"

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
		FavoriteReviewCommentReply(user *entity.User, reviewCommentReplyID int) error
		UnFavoriteReviewCommentReply(user *entity.User, reviewCommentReplyID int) error
	}

	ReviewFavoriteCommandServiceImpl struct {
		ReviewFavoriteCommandRepository repository.ReviewFavoriteCommandRepository
		ReviewFavoriteQueryRepository   repository.ReviewFavoriteQueryRepository
		ReviewQueryRepository           repository.ReviewQueryRepository
		ReviewCommandRepository         repository.ReviewCommandRepository
		service.NoticeDomainService
		TransactionService
	}
)

var ReviewFavoriteCommandServiceSet = wire.NewSet(
	wire.Struct(new(ReviewFavoriteCommandServiceImpl), "*"),
	wire.Bind(new(ReviewFavoriteCommandService), new(*ReviewFavoriteCommandServiceImpl)),
)

func (s *ReviewFavoriteCommandServiceImpl) Store(user *entity.User, reviewID int) error {
	isExistReview, err := s.ReviewQueryRepository.IsExist(reviewID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if !isExistReview {
		return serror.New(nil, serror.CodeNotFound, "Not found")
	}

	isExistFavorite, err := s.ReviewFavoriteQueryRepository.IsExist(user.ID, reviewID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if isExistFavorite {
		return serror.New(nil, serror.CodeInvalidParam, "already set in table")
	}

	favorite := entity.NewUserFavoriteReview(user.ID, reviewID)

	return s.TransactionService.Do(func(c context.Context) error {
		if err := s.ReviewFavoriteCommandRepository.Store(c, favorite); err != nil {
			return errors.Wrap(err, "failed to store")
		}

		if err := s.ReviewCommandRepository.IncrementFavoriteCount(c, reviewID); err != nil {
			return errors.Wrap(err, "failed to update post")
		}

		review, err := s.ReviewQueryRepository.FindByID(reviewID)
		if err != nil {
			return errors.Wrap(err, "failed to find review by id")
		}

		return s.NoticeDomainService.FavoriteReview(c, favorite, review, user)
	})
}

func (s *ReviewFavoriteCommandServiceImpl) Delete(user *entity.User, reviewID int) error {
	isExistReview, err := s.ReviewQueryRepository.IsExist(reviewID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if !isExistReview {
		return serror.New(nil, serror.CodeNotFound, "Not found")
	}

	isExistFavorite, err := s.ReviewFavoriteQueryRepository.IsExist(user.ID, reviewID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if !isExistFavorite {
		return serror.New(nil, serror.CodeInvalidParam, "not set in table yet")
	}

	unfavorite := entity.NewUserFavoriteReview(user.ID, reviewID)

	//TODO:lockを取る
	return s.TransactionService.Do(func(c context.Context) error {
		if err := s.ReviewFavoriteCommandRepository.Delete(c, unfavorite); err != nil {
			return errors.Wrap(err, "failed to delete")
		}

		if err := s.ReviewCommandRepository.DecrementFavoriteCount(c, reviewID); err != nil {
			return errors.Wrap(err, "failed to update post")
		}

		return nil
	})
}

func (s *ReviewFavoriteCommandServiceImpl) FavoriteReviewCommentReply(user *entity.User, reviewCommentReplyID int) error {
	isExistReply, err := s.ReviewQueryRepository.IsExistReviewCommentReply(reviewCommentReplyID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if !isExistReply {
		return serror.New(nil, serror.CodeNotFound, "Not found")
	}

	isExistFavorite, err := s.ReviewFavoriteQueryRepository.IsExistReviewCommentReply(user.ID, reviewCommentReplyID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if isExistFavorite {
		return serror.New(nil, serror.CodeInvalidParam, "already set in the table")
	}

	favorite := entity.NewUserFavoriteReviewCommentReply(user.ID, reviewCommentReplyID)

	return s.TransactionService.Do(func(c context.Context) error {
		if err := s.ReviewFavoriteCommandRepository.StoreReviewCommentReply(c, favorite); err != nil {
			return errors.Wrap(err, "failed to store")
		}

		if err := s.ReviewCommandRepository.IncrementReviewCommentReplyFavoriteCount(c, reviewCommentReplyID); err != nil {
			return errors.Wrap(err, "failed to update post")
		}

		reviewCommentReply, err := s.ReviewQueryRepository.FindReviewCommentReplyByID(reviewCommentReplyID)
		if err != nil {
			return errors.Wrap(err, "failed to find review by id")
		}

		reviewComment, err := s.ReviewQueryRepository.FindReviewCommentByID(reviewCommentReply.ReviewCommentID)
		if err != nil {
			return errors.Wrap(err, "failed find review_comment")
		}

		review, err := s.ReviewQueryRepository.FindByID(reviewComment.ReviewID)
		if err != nil {
			return errors.Wrap(err, "failed find review")
		}

		return s.NoticeDomainService.FavoriteReviewCommentReply(c, favorite, reviewCommentReply, review, user)
	})
}

func (s *ReviewFavoriteCommandServiceImpl) UnFavoriteReviewCommentReply(user *entity.User, reviewCommentReplyID int) error {
	isExistReply, err := s.ReviewQueryRepository.IsExistReviewCommentReply(reviewCommentReplyID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if !isExistReply {
		return serror.New(nil, serror.CodeNotFound, "Not found")
	}

	isExistFavorite, err := s.ReviewFavoriteQueryRepository.IsExistReviewCommentReply(user.ID, reviewCommentReplyID)
	if err != nil {
		return errors.Wrap(err, "failed to IsExist")
	}
	if !isExistFavorite {
		return serror.New(nil, serror.CodeInvalidParam, "not set in the table yet")
	}

	unfavorite := entity.NewUserFavoriteReviewCommentReply(user.ID, reviewCommentReplyID)

	return s.TransactionService.Do(func(c context.Context) error {
		if err := s.ReviewFavoriteCommandRepository.DeleteReviewCommentReply(c, unfavorite); err != nil {
			return errors.Wrap(err, "failed to delete")
		}

		if err := s.ReviewCommandRepository.DecrementReviewCommentReplyFavoriteCount(c, reviewCommentReplyID); err != nil {
			return errors.Wrap(err, "failed to update post")
		}

		return nil
	})
}
