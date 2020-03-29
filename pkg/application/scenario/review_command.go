package scenario

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
)

type (
	// reviewコマンド系シナリオ
	ReviewCommandScenario interface {
		Create(user *entity.User, params *command.CreateReview) error
		UpdateReview(user *entity.User, prams *command.UpdateReview) error
		DeleteReview(id int, user *entity.User) error
	}

	// reviewコマンド系シナリオ実装
	ReviewCommandScenarioImpl struct {
		service.ReviewQueryService
		service.ReviewCommandService
		service.HashtagCommandService
	}
)

var ReviewCommandScenarioSet = wire.NewSet(
	wire.Struct(new(ReviewCommandScenarioImpl), "*"),
	wire.Bind(new(ReviewCommandScenario), new(*ReviewCommandScenarioImpl)),
)

func (s *ReviewCommandScenarioImpl) Create(user *entity.User, param *command.CreateReview) error {
	review := s.convertStoreReviewPramToEntity(param, user)

	hashtags, err := s.HashtagCommandService.FindOrCreateHashtags(model.FindHashtags(review.Body))
	if err != nil {
		return errors.Wrap(err, "failed store and show hashtag")
	}
	review.HashtagIDs = s.convertReviewAndHashtagToReviewHashtag(hashtags, review)

	// touristSpotと紐付くレビューの場合
	if review.TouristSpotID != 0 {
		return s.ReviewCommandService.StoreTouristSpotReview(review)
	}

	// innと紐付くレビューの場合
	return s.ReviewCommandService.StoreInnReview(review)
}

func (s *ReviewCommandScenarioImpl) UpdateReview(user *entity.User, params *command.UpdateReview) error {
	review, err := s.ReviewQueryService.ShowReview(params.ID)
	if err != nil {
		return errors.Wrapf(err, "failed to find review id=%d", params.ID)
	}
	if !review.IsOwner(user.ID) {
		return serror.New(nil, serror.CodeForbidden, "forbidden error")
	}

	hashtags, err := s.HashtagCommandService.FindOrCreateHashtags(model.FindHashtags(params.Body))
	if err != nil {
		return errors.Wrap(err, "failed find or create hashtags")
	}

	review.HashtagIDs = s.convertReviewAndHashtagToReviewHashtag(hashtags, review)

	s.updateReview(params, review)

	return s.ReviewCommandService.UpdateReview(review)
}

func (s *ReviewCommandScenarioImpl) DeleteReview(id int, user *entity.User) error {
	review, err := s.ReviewQueryService.ShowReview(id)
	if err != nil {
		return errors.Wrapf(err, "failed to find review id=%d", id)
	}
	if !review.IsOwner(user.ID) {
		return serror.New(nil, serror.CodeForbidden, "forbidden error")
	}

	return s.ReviewCommandService.DeleteReview(review)
}

func (s *ReviewCommandScenarioImpl) convertReviewAndHashtagToReviewHashtag(hashtags []*entity.Hashtag, review *entity.Review) []*entity.ReviewHashtag {
	reviewHashtagList := make([]*entity.ReviewHashtag, len(hashtags))
	for i, hashtag := range hashtags {
		reviewHashtagList[i] = &entity.ReviewHashtag{
			ReviewID:  review.ID,
			HashtagID: hashtag.ID,
		}
	}
	return reviewHashtagList
}

// MEMO: HashtagIDsはセットされない
func (s *ReviewCommandScenarioImpl) convertStoreReviewPramToEntity(param *command.CreateReview, user *entity.User) *entity.Review {
	reviewMedias := make([]*entity.ReviewMedia, len(param.MediaUUIDs))
	for i, media := range param.MediaUUIDs {
		reviewMedias[i] = entity.NewReviewMedia(media.UUID, media.MimeType, i+1)
	}

	if param.TouristSpotID != 0 {
		return &entity.Review{
			UserID:        user.ID,
			TouristSpotID: param.TouristSpotID,
			Score:         param.Score,
			MediaCount:    len(param.MediaUUIDs),
			Body:          param.Body,
			TravelDate:    param.TravelDate.Time,
			Accompanying:  param.Accompanying,
			Medias:        reviewMedias,
		}
	}

	return &entity.Review{
		UserID:       user.ID,
		InnID:        param.InnID,
		Score:        param.Score,
		MediaCount:   len(param.MediaUUIDs),
		Body:         param.Body,
		TravelDate:   param.TravelDate.Time,
		Accompanying: param.Accompanying,
		Medias:       reviewMedias,
	}
}

// TODO: https://github.com/stayway-corp/stayway-media-api/pull/133#discussion_r394333171
func (s *ReviewCommandScenarioImpl) updateReview(param *command.UpdateReview, review *entity.Review) {
	review.Body = param.Body
	review.TravelDate = param.TravelDate.Time
	review.Accompanying = param.Accompanying
	review.Score = param.Score

	reviewMedias := make([]*entity.ReviewMedia, len(param.MediaUUIDs))
	for i, media := range param.MediaUUIDs {
		reviewMedias[i] = entity.NewReviewMedia(media.UUID, media.MimeType, i+1)
	}
	if len(reviewMedias) > 0 {
		review.Medias = reviewMedias
		review.MediaCount = len(reviewMedias)
	}
}
