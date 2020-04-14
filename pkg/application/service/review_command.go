package service

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/service"
)

type (
	// Reviewコマンド系サービス
	ReviewCommandService interface {
		// 以下2つは処理が特別複雑なので、別途シナリオクラスを作成している
		//*************************************************
		StoreTouristSpotReview(review *entity.Review) error
		StoreInnReview(review *entity.Review) error
		//*************************************************
		UpdateReview(review *entity.Review) error
		DeleteReview(review *entity.Review) error
		CreateReviewCommentReply(user *entity.User, cmd *command.CreateReviewCommentReply) error
		CreateReviewComment(user *entity.User, reviewID int, body string) error
		DeleteReviewComment(user *entity.User, commentID int) error
		FavoriteReviewComment(user *entity.User, reviewCommentID int) error
		UnfavoriteReviewComment(user *entity.User, reviewCommentID int) error
	}

	// Reviewコマンド系サービス実装
	ReviewCommandServiceImpl struct {
		repository.ReviewQueryRepository
		repository.ReviewCommandRepository
		repository.HashtagCommandRepository
		// repository.CategoryQueryRepository
		repository.InnQueryRepository
		repository.TouristSpotCommandRepository
		service.NoticeDomainService
		TransactionService
	}
)

var ReviewCommandServiceSet = wire.NewSet(
	wire.Struct(new(ReviewCommandServiceImpl), "*"),
	wire.Bind(new(ReviewCommandService), new(*ReviewCommandServiceImpl)),
)

// touristSpotと紐付くレビューの場合
func (s *ReviewCommandServiceImpl) StoreTouristSpotReview(review *entity.Review) error {
	// TODO: lock時間長くなるのが気になる
	return s.TransactionService.Do(func(c context.Context) error {
		if err := s.ReviewCommandRepository.StoreReview(c, review); err != nil {
			return errors.Wrap(err, "failed store review")
		}

		if err := s.persistReviewMedia(review.Medias); err != nil {
			return errors.Wrap(err, "failed to persist media")
		}

		// 紐づくtourist_spotの平均値を更新
		if err := s.TouristSpotCommandRepository.UpdateScoreByID(c, review.TouristSpotID); err != nil {
			return errors.Wrap(err, "failed increment hashtag.score")
		}

		return s.NoticeDomainService.Review(c, review)
	})
}

// innと紐付くレビューの場合
func (s *ReviewCommandServiceImpl) StoreInnReview(review *entity.Review) error {
	return s.TransactionService.Do(func(c context.Context) error {
		if err := s.ReviewCommandRepository.StoreReview(c, review); err != nil {
			return errors.Wrap(err, "failed store review")
		}

		if err := s.persistReviewMedia(review.Medias); err != nil {
			return errors.Wrap(err, "failed to persist media")
		}

		return s.NoticeDomainService.Review(c, review)
	})
}

func (s *ReviewCommandServiceImpl) UpdateReview(review *entity.Review) error {
	return s.TransactionService.Do(func(c context.Context) error {

		if err := s.ReviewCommandRepository.StoreReview(c, review); err != nil {
			return errors.Wrap(err, "failed store review")
		}

		if err := s.persistReviewMedia(review.Medias); err != nil {
			return errors.Wrap(err, "failed to persist media")
		}

		return nil
	})
}

func (s *ReviewCommandServiceImpl) CreateReviewCommentReply(user *entity.User, cmd *command.CreateReviewCommentReply) error {
	reply := s.convertCreateReviewCommentReplyToEntity(user, cmd)

	return s.TransactionService.Do(func(c context.Context) error {
		if err := s.ReviewCommandRepository.StoreReviewCommentReply(c, reply); err != nil {
			return errors.Wrap(err, "failed to store review_comment_reply")
		}

		if err := s.ReviewCommandRepository.IncrementReviewCommentReplyCount(c, reply.ReviewCommentID); err != nil {
			return errors.Wrap(err, "failed to increment review_comment.favorite_count")
		}

		comment, err := s.ReviewQueryRepository.FindReviewCommentByID(reply.ReviewCommentID)
		if err != nil {
			return errors.Wrap(err, "failed to find review_comment by id")
		}

		return s.NoticeDomainService.ReviewCommentReply(c, reply, comment)
	})
}

func (s *ReviewCommandServiceImpl) FavoriteReviewComment(user *entity.User, reviewCommentID int) error {
	isExist, err := s.ReviewQueryRepository.IsExistReviewComment(reviewCommentID)
	if err != nil {
		return errors.Wrap(err, "failed to is exist")
	}
	if !isExist {
		return serror.New(nil, serror.CodeNotFound, "review_comment not found")
	}

	favorite := entity.NewUserFavoriteReviewComment(user.ID, reviewCommentID)

	return s.TransactionService.Do(func(c context.Context) error {
		// レビューコメントへのいいねを永続化
		if err := s.ReviewCommandRepository.StoreReviewCommentFavorite(c, favorite); err != nil {
			return errors.Wrap(err, "failed to store user_favorite_review_comment")
		}

		// レビューコメントのいいね数を1インクリメント
		if err := s.ReviewCommandRepository.IncrementReviewCommentFavoriteCount(c, reviewCommentID); err != nil {
			return errors.Wrap(err, "failed to increment review_comment.favorite_count")
		}

		comment, err := s.ReviewQueryRepository.FindReviewCommentByID(reviewCommentID)
		if err != nil {
			return errors.Wrap(err, "failed to find review_comment by id")
		}

		return s.NoticeDomainService.FavoriteReviewComment(c, favorite, comment)
	})
}

func (s *ReviewCommandServiceImpl) DeleteReview(review *entity.Review) error {
	if err := s.ReviewCommandRepository.DeleteReview(context.TODO(), review); err != nil {
		return errors.Wrap(err, "failed to delete delete")
	}
	return nil
}

// TODO: lock取る
func (s *ReviewCommandServiceImpl) UnfavoriteReviewComment(user *entity.User, reviewCommentID int) error {
	isExist, err := s.ReviewQueryRepository.IsExistReviewCommentFavorite(user.ID, reviewCommentID)
	if err != nil {
		return errors.Wrap(err, "failed to is exist")
	}
	if !isExist {
		return serror.New(nil, serror.CodeNotFound, "review_comment_favorite not found")
	}

	return s.TransactionService.Do(func(c context.Context) error {
		// レビューコメントへのいいねを物理削除
		if err := s.ReviewCommandRepository.DeleteReviewCommentFavoriteByID(c, user.ID, reviewCommentID); err != nil {
			return errors.Wrap(err, "failed to delete user_favorite_review_comment")
		}

		// レビューコメントへのいいね数を1デクリメント
		if err := s.ReviewCommandRepository.DecrementReviewCommentFavoriteCount(c, reviewCommentID); err != nil {
			return errors.Wrap(err, "failed to decrement review_comment.favorite_count")
		}

		return nil
	})
}

func (s *ReviewCommandServiceImpl) CreateReviewComment(user *entity.User, reviewID int, body string) error {
	reviewComment := entity.NewReviewComment(user.ID, reviewID, body)

	return s.TransactionService.Do(func(c context.Context) error {
		// レビューコメントを追加
		if err := s.ReviewCommandRepository.CreateReviewComment(c, reviewComment); err != nil {
			return err
		}

		// レビューにひもづくコメント数をインクリメント
		if err := s.ReviewCommandRepository.IncrementReviewCommentCount(c, reviewID); err != nil {
			return err
		}

		review, err := s.ReviewQueryRepository.FindByID(reviewID)
		if err != nil {
			return err
		}

		return s.NoticeDomainService.ReviewComment(c, reviewComment, review)
	})
}

func (s *ReviewCommandServiceImpl) persistReviewMedia(medias []*entity.ReviewMedia) error {
	for _, media := range medias {
		if err := s.ReviewCommandRepository.PersistReviewMedia(media); err != nil {
			return errors.Wrapf(err, "failed to persist media(id=%s)", media.ID)
		}
	}

	return nil
}

func (s *ReviewCommandServiceImpl) DeleteReviewComment(user *entity.User, commentID int) error {
	return s.TransactionService.Do(func(c context.Context) error {
		comment, err := s.ReviewCommandRepository.ShowReviewComment(c, commentID)
		if err != nil {
			return errors.Wrap(err, "failed to find review comment")
		}
		// ユーザに削除権限がなければForbidden
		if !comment.IsOwner(user.ID) {
			return serror.New(nil, serror.CodeForbidden, "forbidden")
		}

		// レビューにひもづくコメント数をデクリメント
		if err := s.ReviewCommandRepository.DecrementReviewCommentCount(c, comment.ReviewID); err != nil {
			return err
		}

		return s.ReviewCommandRepository.DeleteReviewCommentByID(c, comment.ID)
	})
}

func (s *ReviewCommandServiceImpl) convertCreateReviewCommentReplyToEntity(user *entity.User, cmd *command.CreateReviewCommentReply) *entity.ReviewCommentReply {
	return &entity.ReviewCommentReply{
		UserID:          user.ID,
		ReviewCommentID: cmd.ReviewCommentID,
		Body:            cmd.Body,
	}
}
