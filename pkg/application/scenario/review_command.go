package scenario

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
	"gopkg.in/guregu/null.v3"
)

type (
	// reviewコマンド系シナリオ
	ReviewCommandScenario interface {
		Create(user *entity.User, params *command.CreateReview) (*entity.ReviewDetailWithIsFavorite, error)
		UpdateReview(user *entity.User, prams *command.UpdateReview) error
		DeleteReview(id int, user *entity.User) error
	}

	// reviewコマンド系シナリオ実装
	ReviewCommandScenarioImpl struct {
		service.ReviewQueryService
		service.ReviewCommandService
		service.HashtagCommandService
		repository.UserQueryRepository
	}
)

var ReviewCommandScenarioSet = wire.NewSet(
	wire.Struct(new(ReviewCommandScenarioImpl), "*"),
	wire.Bind(new(ReviewCommandScenario), new(*ReviewCommandScenarioImpl)),
)

func (s *ReviewCommandScenarioImpl) Create(user *entity.User, param *command.CreateReview) (*entity.ReviewDetailWithIsFavorite, error) {
	review := s.convertStoreReviewPramToEntity(param, user)

	hashtags, err := s.HashtagCommandService.FindOrCreateHashtags(model.FindHashtags(review.Body))
	if err != nil {
		return nil, errors.Wrap(err, "failed store and show hashtag")
	}
	review.HashtagIDs = s.convertReviewAndHashtagToReviewHashtag(hashtags, review)

	ouser := user.ConvertToOptionalUser()

	// touristSpotと紐付くレビューの場合
	if review.TouristSpotID.Valid {
		if err := s.ReviewCommandService.StoreTouristSpotReview(user, review); err != nil {
			return nil, errors.Wrap(err, "failed to store touristSpotReview")
		}
	} else {
		if err := s.ReviewCommandService.StoreInnReview(user, review); err != nil {
			return nil, errors.Wrap(err, "failed to store innReview")
		}
	}

	// 保存したreviewを取得
	resolve, err := s.ReviewQueryService.ShowQueryReview(review.ID, *ouser)
	if err != nil {
		return nil, errors.Wrap(err, "failed show review")
	}

	// 自身をフォローする事はできないので、review.UserのIsFollowチェックはしない
	return resolve, nil
}

func (s *ReviewCommandScenarioImpl) UpdateReview(user *entity.User, cmd *command.UpdateReview) error {
	review, err := s.ReviewQueryService.ShowReview(cmd.ID)
	if err != nil {
		return errors.Wrapf(err, "failed to find review id=%d", cmd.ID)
	}
	if !user.IsSelfID(review.UserID) {
		return serror.New(nil, serror.CodeForbidden, "forbidden error")
	}

	hashtags, err := s.HashtagCommandService.FindOrCreateHashtags(model.FindHashtags(cmd.Body))
	if err != nil {
		return errors.Wrap(err, "failed find or create hashtags")
	}

	review.HashtagIDs = s.convertReviewAndHashtagToReviewHashtag(hashtags, review)

	s.updateReview(cmd, review)

	return s.ReviewCommandService.UpdateReview(review, cmd)
}

func (s *ReviewCommandScenarioImpl) DeleteReview(id int, user *entity.User) error {
	review, err := s.ReviewQueryService.ShowReview(id)
	if err != nil {
		return errors.Wrapf(err, "failed to find review id=%d", id)
	}
	if !user.IsSelfID(review.UserID) {
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
			TouristSpotID: null.IntFrom(int64(param.TouristSpotID)),
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
		InnID:        null.IntFrom(int64(param.InnID)),
		Score:        param.Score,
		MediaCount:   len(param.MediaUUIDs),
		Body:         param.Body,
		TravelDate:   param.TravelDate.Time,
		Accompanying: param.Accompanying,
		Medias:       reviewMedias,
	}
}

// TODO: https://github.com/stayway-corp/stayway-media-api/pull/133#discussion_r394333171
func (s *ReviewCommandScenarioImpl) updateReview(cmd *command.UpdateReview, review *entity.Review) {
	review.Body = cmd.Body
	review.TravelDate = cmd.TravelDate.Time
	review.Accompanying = cmd.Accompanying
	review.Score = cmd.Score

	reviewMedias := make([]*entity.ReviewMedia, len(cmd.MediaUUIDs))
	for i, media := range cmd.MediaUUIDs {
		reviewMedias[i] = entity.NewReviewMedia(media.UUID, media.MimeType, i+1)
	}
	if len(reviewMedias) > 0 {
		review.Medias = reviewMedias
		review.MediaCount = len(reviewMedias)
	}
}
