package scenario

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	// reviewコマンド系シナリオ
	ReviewCommandScenario interface {
		Create(param *param.StoreReviewParam, user *entity.User) error
	}

	// reviewコマンド系シナリオ実装
	ReviewCommandScenarioImpl struct {
		service.ReviewCommandService
		service.HashtagCommandService
	}
)

var ReviewCommandScenarioSet = wire.NewSet(
	wire.Struct(new(ReviewCommandScenarioImpl), "*"),
	wire.Bind(new(ReviewCommandScenario), new(*ReviewCommandScenarioImpl)),
)

func (s *ReviewCommandScenarioImpl) Create(param *param.StoreReviewParam, user *entity.User) error {
	review := s.convertStoreReviewPramToEntity(param, user)

	hashtagFromStr := s.convertStringSliceToHashtag(model.FindHashtags(review.Body))

	hashtags := make([]*entity.Hashtag, len(hashtagFromStr))

	for i, hashtag := range hashtagFromStr {
		hashtagEntity, err := s.HashtagCommandService.FindOrCreateHashtag(hashtag)
		if err != nil {
			return errors.Wrap(err, "failed store and show hashtag")
		}
		hashtags[i] = hashtagEntity
	}

	review.HashtagIDs = s.convertReviewAndHashtagToReviewHashtag(hashtags, review)

	// touristSpotと紐付くレビューの場合
	if review.TouristSpotID != 0 {
		return s.ReviewCommandService.StoreTouristSpotReview(review)
	}

	return s.ReviewCommandService.StoreInnReview(review)
}

func (s *ReviewCommandScenarioImpl) convertStringSliceToHashtag(stringList []string) []*entity.Hashtag {
	hashtags := make([]*entity.Hashtag, len(stringList))
	for i, str := range stringList {
		hashtags[i] = &entity.Hashtag{
			Name: str,
		}
	}
	return hashtags
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
func (s *ReviewCommandScenarioImpl) convertStoreReviewPramToEntity(param *param.StoreReviewParam, user *entity.User) *entity.Review {
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
		InnID:        param.InnId,
		Score:        param.Score,
		MediaCount:   len(param.MediaUUIDs),
		Body:         param.Body,
		TravelDate:   param.TravelDate.Time,
		Accompanying: param.Accompanying,
		Medias:       reviewMedias,
	}
}
