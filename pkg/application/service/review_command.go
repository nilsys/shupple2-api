package service

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	// Reviewコマンド系サービス
	ReviewCommandService interface {
		StoreTouristSpotReview(review *entity.Review) error
		StoreInnReview(review *entity.Review) error
		CreateReviewComment(user *entity.User, reviewID int, body string) error
	}

	// Reviewコマンド系サービス実装
	ReviewCommandServiceImpl struct {
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

		if len(review.HashtagIDs) <= 0 {
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

		if len(review.HashtagIDs) <= 0 {
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

		// レビューにひもづくコメント数
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
