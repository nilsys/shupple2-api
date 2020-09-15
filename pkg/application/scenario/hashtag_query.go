package scenario

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	HashtagQueryScenario interface {
		ListRecommend(areID, subAreaID, subSubAreaID, limit int, ouser *entity.OptionalUser) (*entity.Hashtags, map[int]bool, error)
		Show(name string, ouser *entity.OptionalUser) (*entity.Hashtag, map[int]bool, error)
	}

	HashtagQueryScenarioImpl struct {
		service.HashtagQueryService
		repository.HashtagQueryRepository
	}
)

var HashtagQueryScenarioSet = wire.NewSet(
	wire.Struct(new(HashtagQueryScenarioImpl), "*"),
	wire.Bind(new(HashtagQueryScenario), new(*HashtagQueryScenarioImpl)),
)

func (s *HashtagQueryScenarioImpl) ListRecommend(areID, subAreaID, subSubAreaID, limit int, ouser *entity.OptionalUser) (*entity.Hashtags, map[int]bool, error) {
	var (
		isFollowing map[int]bool
	)

	hashtags, err := s.HashtagQueryService.ShowRecommendList(areID, subAreaID, subSubAreaID, limit)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed find hashtag")
	}

	if ouser.IsAuthorized() {
		isFollowing, err = s.IsFollowing(ouser.ID, hashtags.IDs())
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed find user_follow_hashtag")
		}
	}

	return hashtags, isFollowing, nil
}

func (s *HashtagQueryScenarioImpl) Show(name string, ouser *entity.OptionalUser) (*entity.Hashtag, map[int]bool, error) {
	var (
		isFollowing map[int]bool
	)

	hashtag, err := s.HashtagQueryService.Show(name)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed find hashtag")
	}

	if ouser.IsAuthorized() {
		isFollowing, err = s.IsFollowing(ouser.ID, []int{hashtag.ID})
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed find user_follow_hashtag")
		}
	}

	return hashtag, isFollowing, nil
}
