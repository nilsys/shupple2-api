package service

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
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
		repository.CategoryQueryRepository
		repository.InnQueryRepository
		repository.TouristSpotCommandRepository
		TransactionService
	}
)

var ReviewCommandServiceSet = wire.NewSet(
	wire.Struct(new(ReviewCommandServiceImpl), "*"),
	wire.Bind(new(ReviewCommandService), new(*ReviewCommandServiceImpl)),
)

// touristSpotと紐付くレビューの場合
func (s *ReviewCommandServiceImpl) StoreTouristSpotReview(review *entity.Review) error {

	// レビューに紐付くtouristSpotに紐付くカテゴリーを取得
	categories, err := s.CategoryQueryRepository.FindByTouristSpotID(review.TouristSpotID)
	if err != nil {
		return errors.Wrap(err, "failed find category list by tourist_spot")
	}
	// レビューに紐付くtouristSpotに紐付くカテゴリーとハッシュタグを紐付けるHashtagCategoryをgen
	hashtagCategoryList := s.convertCategoryAndHashtagIDsToHashtagCategory(categories, review.HashtagIDs)

	// TODO: lock時間長くなるのが気になる
	return s.TransactionService.Do(func(c context.Context) error {
		if err := s.ReviewCommandRepository.StoreReview(c, review); err != nil {
			return errors.Wrap(err, "failed store review")
		}

		if err := s.persistReviewMedia(review.Medias); err != nil {
			return errors.Wrap(err, "failed to persist media")
		}

		if !review.HashHashtagIDs() {
			return nil
		}

		// TODO: forで回したく無い...
		for _, hashtagCategory := range hashtagCategoryList {
			// HashtagCategoryを永続化
			if err := s.HashtagCommandRepository.StoreHashtagCategory(c, hashtagCategory); err != nil {
				return errors.Wrap(err, "failed store hashtag_category")
			}
		}
		for _, hashtagCategory := range hashtagCategoryList {
			// 紐付けられたハッシュタグのscoreをインクリメント
			if err := s.HashtagCommandRepository.IncrementScoreByID(c, hashtagCategory.HashtagID); err != nil {
				return errors.Wrap(err, "failed increment hashtag.score")
			}
		}

		// 紐づくtourist_spotの平均値を更新
		if err := s.TouristSpotCommandRepository.UpdateScoreByID(c, review.TouristSpotID); err != nil {
			return errors.Wrap(err, "failed increment hashtag.score")
		}

		return nil
	})
}

// innと紐付くレビューの場合
func (s *ReviewCommandServiceImpl) StoreInnReview(review *entity.Review) error {

	// レビューに紐付くinnに紐付くカテゴリーIDをタイプ別に取得
	innAreaTypeIDs, err := s.InnQueryRepository.FindAreaIDsByID(review.InnID)
	if err != nil {
		return errors.Wrap(err, "failed get inn area details from stayway api")
	}

	// stayway APIから取得したカテゴリーIDはcategoryテーブルのmetasearchIDと紐づいているので
	categories, err := s.CategoryQueryRepository.FindByMetaSearchID(innAreaTypeIDs)
	if err != nil {
		return errors.Wrap(err, "failed get category list by metasearch id")
	}
	hashtagCategoryList := s.convertCategoryAndHashtagIDsToHashtagCategory(categories, review.HashtagIDs)

	return s.TransactionService.Do(func(c context.Context) error {
		if err := s.ReviewCommandRepository.StoreReview(c, review); err != nil {
			return errors.Wrap(err, "failed store review")
		}

		if err := s.persistReviewMedia(review.Medias); err != nil {
			return errors.Wrap(err, "failed to persist media")
		}

		if !review.HashHashtagIDs() {
			return nil
		}
		for _, hashtagCategory := range hashtagCategoryList {
			// HashtagCategoryを永続化
			if err := s.HashtagCommandRepository.StoreHashtagCategory(c, hashtagCategory); err != nil {
				return errors.Wrap(err, "failed store hashtag_category")
			}
		}
		for _, hashtagCategory := range hashtagCategoryList {
			// 紐付けられたハッシュタグのscoreをインクリメント
			if err := s.HashtagCommandRepository.IncrementScoreByID(c, hashtagCategory.HashtagID); err != nil {
				return errors.Wrap(err, "failed increment hashtag.score")
			}
		}
		return nil
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

		return nil
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

		return nil
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

func (s *ReviewCommandServiceImpl) convertCategoryAndHashtagIDsToHashtagCategory(categories []*entity.Category, hashtagIDs []*entity.ReviewHashtag) []*entity.HashtagCategory {
	var hashtagCategories []*entity.HashtagCategory

	for _, id := range hashtagIDs {
		for _, category := range categories {
			hashtagCategories = append(hashtagCategories, entity.NewHashtagCategory(id.HashtagID, category.ID))
		}
	}

	return hashtagCategories
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

		return nil
	})
}

func (s *ReviewCommandServiceImpl) persistReviewMedia(medias []*entity.ReviewMedia) error {
	for _, media := range medias {
		if err := s.ReviewCommandRepository.PersistReviewMedia(media); err != nil {
			return errors.Wrapf(err, "failed to presist media(id=%s)", media.ID)
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

		return s.ReviewCommandRepository.DeleteReviewCommentByID(c, comment)
	})
}

func (s *ReviewCommandServiceImpl) convertCreateReviewCommentReplyToEntity(user *entity.User, cmd *command.CreateReviewCommentReply) *entity.ReviewCommentReply {
	return &entity.ReviewCommentReply{
		UserID:          user.ID,
		ReviewCommentID: cmd.ReviewCommentID,
		Body:            cmd.Body,
	}
}
