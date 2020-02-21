package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	SearchQueryService interface {
		ShowSearchSuggestionListByKeyward(keywards string) ([]*entity.SearchSuggetion, error)
	}

	SearchQueryServiceImpl struct {
		repository.CategoryQueryRepository
		repository.TouristSpotQueryRepository
		repository.HashtagQueryRepository
		repository.UserQueryRepository
	}
)

var SearchQueryServiceSet = wire.NewSet(
	wire.Struct(new(SearchQueryServiceImpl), "*"),
	wire.Bind(new(SearchQueryService), new(*SearchQueryServiceImpl)),
)

func (s *SearchQueryServiceImpl) ShowSearchSuggestionListByKeyward(keyward string) ([]*entity.SearchSuggetion, error) {
	categories, err := s.CategoryQueryRepository.SearchByName(keyward)
	if err != nil {
		return nil, errors.Wrap(err, "failed search category list by keyward")
	}
	touristSpots, err := s.TouristSpotQueryRepository.SearchByName(keyward)
	if err != nil {
		return nil, errors.Wrap(err, "failed search tourist_spot list by keyward")
	}
	hashtags, err := s.HashtagQueryRepository.SearchByName(keyward)
	if err != nil {
		return nil, errors.Wrap(err, "failed search hashtag list by keyward")
	}
	users, err := s.UserQueryRepository.SearchByName(keyward)
	if err != nil {
		return nil, errors.Wrap(err, "failed search user list by keyward")
	}

	suggestions := s.convertTargetToSuggetion(categories, touristSpots, hashtags, users)

	return suggestions, nil
}

// TODO: リファクタ
// ただでさえ処理重めなのに。。
func (s *SearchQueryServiceImpl) convertTargetToSuggetion(categories []*entity.Category, touristSpots []*entity.TouristSpot, hashtags []*entity.Hashtag, users []*entity.User) []*entity.SearchSuggetion {
	suggetions := make([]*entity.SearchSuggetion, len(categories)+len(touristSpots)+len(hashtags)+len(users))

	for i, category := range categories {
		if category.Type == model.CategoryTypeArea {
			suggetions[i] = &entity.SearchSuggetion{
				ID:   category.ID,
				Type: model.SuggestionTypeArea,
				Name: category.Name,
			}
		}
		if category.Type == model.CategoryTypeSubArea {
			suggetions[i] = &entity.SearchSuggetion{
				ID:   category.ID,
				Type: model.SuggestionTypeSubArea,
				Name: category.Name,
			}
		}
		if category.Type == model.CategoryTypeSubSubArea {
			suggetions[i] = &entity.SearchSuggetion{
				ID:   category.ID,
				Type: model.SuggestionTypeSubSubArea,
				Name: category.Name,
			}
		}
	}

	for i, touristSpot := range touristSpots {
		suggetions[len(categories)+i] = &entity.SearchSuggetion{
			ID:   touristSpot.ID,
			Type: model.SuggestionTypeTouristSpot,
			Name: touristSpot.Name,
		}
	}

	for i, hashTag := range hashtags {
		suggetions[len(categories)+len(touristSpots)+i] = &entity.SearchSuggetion{
			ID:   hashTag.ID,
			Type: model.SuggestionTypeHashTag,
			Name: hashTag.Name,
		}
	}

	for i, user := range users {
		suggetions[len(categories)+len(touristSpots)+len(hashtags)+i] = &entity.SearchSuggetion{
			ID:   user.ID,
			Type: model.SuggestionTypeUser,
			Name: user.Name,
		}
	}

	return suggetions
}
