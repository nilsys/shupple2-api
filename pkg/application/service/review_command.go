package service

import (
	"context"
	"strings"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/service"
)

type (
	// Reviewコマンド系サービス
	ReviewCommandService interface {
		StoreTouristSpotReview(user *entity.User, review *entity.Review) error
		StoreInnReview(user *entity.User, review *entity.Review) error
		UpdateReview(review *entity.Review, cmd *command.UpdateReview) error
		DeleteReview(review *entity.Review) error
		CreateReviewComment(user *entity.User, reviewID int, body string) (*entity.ReviewCommentTiny, error)
		CreateReviewCommentReply(user *entity.User, cmd *command.CreateReviewCommentReply) (*entity.ReviewCommentReplyTiny, error)
		DeleteReviewComment(user *entity.User, commentID int) error
		DeleteReviewCommentReply(user *entity.User, replyID int) error
		FavoriteReviewComment(user *entity.User, reviewCommentID int) error
		UnfavoriteReviewComment(user *entity.User, reviewCommentID int) error
	}

	// Reviewコマンド系サービス実装
	ReviewCommandServiceImpl struct {
		repository.ReviewQueryRepository
		repository.ReviewCommandRepository
		repository.HashtagCommandRepository
		repository.MediaCommandRepository
		repository.InnQueryRepository
		repository.TouristSpotCommandRepository
		service.NoticeDomainService
		TransactionService
		MediaCommandService
	}
)

var ReviewCommandServiceSet = wire.NewSet(
	wire.Struct(new(ReviewCommandServiceImpl), "*"),
	wire.Bind(new(ReviewCommandService), new(*ReviewCommandServiceImpl)),
)

// touristSpotと紐付くレビューの場合
func (s *ReviewCommandServiceImpl) StoreTouristSpotReview(user *entity.User, review *entity.Review) error {
	// TODO: lock時間長くなるのが気になる
	return s.TransactionService.Do(func(c context.Context) error {
		if err := s.ReviewCommandRepository.StoreReview(c, review); err != nil {
			return errors.Wrap(err, "failed store review")
		}

		if err := s.persistReviewMedia(review.Medias); err != nil {
			return errors.Wrap(err, "failed to persist media")
		}

		// 紐づくtourist_spotの平均値を更新
		if err := s.TouristSpotCommandRepository.UpdateScoreByID(c, review.TouristSpotID.Int64); err != nil {
			return errors.Wrap(err, "failed tourist_spot.rate")
		}

		return s.NoticeDomainService.Review(c, review, user)
	})
}

// innと紐付くレビューの場合
func (s *ReviewCommandServiceImpl) StoreInnReview(user *entity.User, review *entity.Review) error {
	return s.TransactionService.Do(func(c context.Context) error {
		if err := s.ReviewCommandRepository.StoreReview(c, review); err != nil {
			return errors.Wrap(err, "failed store review")
		}

		if err := s.persistReviewMedia(review.Medias); err != nil {
			return errors.Wrap(err, "failed to persist media")
		}

		return s.NoticeDomainService.Review(c, review, user)
	})
}

func (s *ReviewCommandServiceImpl) UpdateReview(review *entity.Review, cmd *command.UpdateReview) error {
	return s.TransactionService.Do(func(c context.Context) error {

		if err := s.ReviewCommandRepository.StoreReview(c, review); err != nil {
			return errors.Wrap(err, "failed store review")
		}

		if review.TouristSpotID.Valid {
			if err := s.TouristSpotCommandRepository.UpdateScoreByID(c, review.TouristSpotID.Int64); err != nil {
				return errors.Wrap(err, "failed update tourist_spot.rate")
			}
		}

		// media(写真)に変更があった場合のみ
		if cmd.HasMedia() {
			if err := s.persistReviewMedia(review.Medias); err != nil {
				return errors.Wrap(err, "failed to persist media")
			}
		}

		return nil
	})
}

func (s *ReviewCommandServiceImpl) CreateReviewCommentReply(user *entity.User, cmd *command.CreateReviewCommentReply) (*entity.ReviewCommentReplyTiny, error) {
	reviewCommentReply := s.convertCreateReviewCommentReplyToEntity(user, cmd)

	err := s.TransactionService.Do(func(c context.Context) error {
		if err := s.ReviewCommandRepository.StoreReviewCommentReply(c, reviewCommentReply); err != nil {
			return errors.Wrap(err, "failed to store review_comment_reply")
		}

		if err := s.ReviewCommandRepository.IncrementReviewCommentReplyCount(c, reviewCommentReply.ReviewCommentID); err != nil {
			return errors.Wrap(err, "failed to increment review_comment.favorite_count")
		}

		comment, err := s.ReviewQueryRepository.FindReviewCommentByID(reviewCommentReply.ReviewCommentID)
		if err != nil {
			return errors.Wrap(err, "failed to find review_comment by id")
		}

		review, err := s.ReviewQueryRepository.FindByID(comment.ReviewID)
		if err != nil {
			return errors.Wrap(err, "failed find review")
		}

		return s.NoticeDomainService.ReviewCommentReply(c, reviewCommentReply, comment, review, user)
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create review_comment_reply transaction")
	}

	reply, err := s.ReviewQueryRepository.FindReviewCommentReplyByID(reviewCommentReply.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find review_comment_reply.idj")
	}

	return reply, nil
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

		review, err := s.ReviewQueryRepository.FindByID(comment.ReviewID)
		if err != nil {
			return errors.Wrap(err, "failed find review")
		}

		return s.NoticeDomainService.FavoriteReviewComment(c, favorite, comment, review, user)
	})
}

func (s *ReviewCommandServiceImpl) DeleteReview(review *entity.Review) error {
	return s.TransactionService.Do(func(ctx context.Context) error {
		if err := s.ReviewCommandRepository.DeleteReview(ctx, review); err != nil {
			return errors.Wrap(err, "failed to delete delete")
		}

		if review.TouristSpotID.Valid {
			if err := s.TouristSpotCommandRepository.UpdateScoreByID(ctx, review.TouristSpotID.Int64); err != nil {
				return errors.Wrap(err, "filed update tourist_spot.score")
			}
		}
		return nil
	})
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

func (s *ReviewCommandServiceImpl) CreateReviewComment(user *entity.User, reviewID int, body string) (*entity.ReviewCommentTiny, error) {
	reviewComment := entity.NewReviewComment(user.ID, reviewID, body)

	err := s.TransactionService.Do(func(c context.Context) error {
		// レビューコメントを追加
		if err := s.ReviewCommandRepository.StoreReviewComment(c, reviewComment); err != nil {
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

		return s.NoticeDomainService.ReviewComment(c, reviewComment, review, user)
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create review_comment")

	}

	comment, err := s.ReviewQueryRepository.FindReviewCommentByID(reviewComment.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find review_comment.id")
	}

	return comment, nil
}

func (s *ReviewCommandServiceImpl) persistReviewMedia(mediaList []*entity.ReviewMedia) error {
	// 先に全部のMimeTypeをチェックする
	mediaTypes := make([]model.MediaType, len(mediaList))
	for i, media := range mediaList {
		switch {
		case strings.HasPrefix(media.MimeType, "video/"):
			mediaTypes[i] = model.MediaTypeReviewVideo
		case strings.HasPrefix(media.MimeType, "image/"):
			mediaTypes[i] = model.MediaTypeReviewImage
		default:
			return serror.New(nil, serror.CodeUnsupportedMedia, "unsupported media: %s", media.MimeType)
		}
	}

	for i, media := range mediaList {
		if err := s.MediaCommandService.PreparePersist(media.ID, media.S3Path(), mediaTypes[i]); err != nil {
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
		if !user.IsSelfID(comment.UserID) {
			return serror.New(nil, serror.CodeForbidden, "forbidden")
		}

		// レビューにひもづくコメント数をデクリメント
		if err := s.ReviewCommandRepository.DecrementReviewCommentCount(c, comment.ReviewID); err != nil {
			return err
		}

		return s.ReviewCommandRepository.DeleteReviewCommentByID(c, comment.ID)
	})
}

func (s *ReviewCommandServiceImpl) DeleteReviewCommentReply(user *entity.User, replyID int) error {
	reply, err := s.ReviewQueryRepository.FindReviewCommentReplyByID(replyID)
	if err != nil {
		return errors.Wrap(err, "failed to find review comment reply")
	}
	// ユーザに削除権限がなければForbidden
	if !user.IsSelfID(reply.UserID) {
		return serror.New(nil, serror.CodeForbidden, "forbidden")
	}

	return s.TransactionService.Do(func(c context.Context) error {

		// コメントにひもづくリプライ数をデクリメント
		if err := s.ReviewCommandRepository.DecrementReviewCommentReplyCount(c, reply.ReviewCommentID); err != nil {
			return err
		}

		return s.ReviewCommandRepository.DeleteReviewCommentReplyByID(c, reply.ID)
	})
}

func (s *ReviewCommandServiceImpl) convertCreateReviewCommentReplyToEntity(user *entity.User, cmd *command.CreateReviewCommentReply) *entity.ReviewCommentReplyTiny {
	return &entity.ReviewCommentReplyTiny{
		UserID:          user.ID,
		ReviewCommentID: cmd.ReviewCommentID,
		Body:            cmd.Body,
	}
}
