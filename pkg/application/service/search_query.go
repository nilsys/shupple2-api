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
		ListSuggestionsByKeyward(keyward string) (*entity.SearchSuggetions, error)
		ListSuggestionsByType(keyward string, suggestionType model.SuggestionType) (*entity.SearchSuggetions, error)
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

func (s *SearchQueryServiceImpl) ListSuggestionsByKeyward(keyward string) (*entity.SearchSuggetions, error) {
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

	return s.convertTargetToSuggestion(categories, touristSpots, hashtags, users), nil
}

// 選択された候補のみ返す
// TODO: リファクタ
func (s *SearchQueryServiceImpl) ListSuggestionsByType(keyward string, suggestionType model.SuggestionType) (*entity.SearchSuggetions, error) {
	if suggestionType == model.SuggestionTypeArea {
		categories, err := s.CategoryQueryRepository.SearchAreaByName(keyward)
		if err != nil {
			return nil, errors.Wrap(err, "failed search area list by keyward")
		}
		return s.convertTargetToSuggestion(categories, nil, nil, nil), nil
	}
	if suggestionType == model.SuggestionTypeSubArea {
		categories, err := s.CategoryQueryRepository.SearchSubAreaByName(keyward)
		if err != nil {
			return nil, errors.Wrap(err, "failed search sub_area list by keyward")
		}
		return s.convertTargetToSuggestion(categories, nil, nil, nil), nil
	}
	if suggestionType == model.SuggestionTypeSubSubArea {
		categories, err := s.CategoryQueryRepository.SearchSubSubAreaByName(keyward)
		if err != nil {
			return nil, errors.Wrap(err, "failed search sub_sub_area list by keyward")
		}
		return s.convertTargetToSuggestion(categories, nil, nil, nil), nil
	}
	if suggestionType == model.SuggestionTypeTouristSpot {
		touristSpots, err := s.TouristSpotQueryRepository.SearchByName(keyward)
		if err != nil {
			return nil, errors.Wrap(err, "failed search tourist_spot list by keyward")
		}
		return s.convertTargetToSuggestion(nil, touristSpots, nil, nil), nil
	}
	if suggestionType == model.SuggestionTypeHashTag {
		hashtags, err := s.HashtagQueryRepository.SearchByName(keyward)
		if err != nil {
			return nil, errors.Wrap(err, "failed search hashtag list by keyward")
		}
		return s.convertTargetToSuggestion(nil, nil, hashtags, nil), nil
	}
	if suggestionType == model.SuggestionTypeUser {
		users, err := s.UserQueryRepository.SearchByName(keyward)
		if err != nil {
			return nil, errors.Wrap(err, "failed search user list by keyward")
		}
		return s.convertTargetToSuggestion(nil, nil, nil, users), nil
	}
	return nil, nil
}

// TODO: リファクタ
func (s *SearchQueryServiceImpl) convertTargetToSuggestion(categories []*entity.Category, touristSpots []*entity.TouristSpot, hashtags []*entity.Hashtag, users []*entity.User) *entity.SearchSuggetions {
	areaSuggestions := make([]*entity.SearchSuggetion, len(categories))
	touristSpotSuggestions := make([]*entity.SearchSuggetion, len(touristSpots))
	hashtagSuggestions := make([]*entity.SearchSuggetion, len(hashtags))
	userSuggestions := make([]*entity.SearchSuggetion, len(users))

	for i, category := range categories {
		if category.Type == model.CategoryTypeArea {
			areaSuggestions[i] = &entity.SearchSuggetion{
				ID:   category.ID,
				Type: model.SuggestionTypeArea,
				Name: category.Name,
			}
		}
		if category.Type == model.CategoryTypeSubArea {
			areaSuggestions[i] = &entity.SearchSuggetion{
				ID:   category.ID,
				Type: model.SuggestionTypeSubArea,
				Name: category.Name,
			}
		}
		if category.Type == model.CategoryTypeSubSubArea {
			areaSuggestions[i] = &entity.SearchSuggetion{
				ID:   category.ID,
				Type: model.SuggestionTypeSubSubArea,
				Name: category.Name,
			}
		}
	}

	for i, touristSpot := range touristSpots {
		touristSpotSuggestions[i] = &entity.SearchSuggetion{
			ID:   touristSpot.ID,
			Type: model.SuggestionTypeTouristSpot,
			Name: touristSpot.Name,
		}
	}

	for i, hashTag := range hashtags {
		hashtagSuggestions[i] = &entity.SearchSuggetion{
			ID:   hashTag.ID,
			Type: model.SuggestionTypeHashTag,
			Name: hashTag.Name,
		}
	}

	for i, user := range users {
		userSuggestions[i] = &entity.SearchSuggetion{
			ID:   user.ID,
			Type: model.SuggestionTypeUser,
			Name: user.Name,
		}
	}

	return entity.NewSearchSuggestions(areaSuggestions, touristSpotSuggestions, hashtagSuggestions, userSuggestions)
}
