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
		ListSuggestionsByKeyward(keyword string) (*entity.SearchSuggetions, error)
		ListSuggestionsByType(keyword string, suggestionType model.SuggestionType) (*entity.SearchSuggetions, error)
	}

	SearchQueryServiceImpl struct {
		repository.AreaCategoryQueryRepository
		repository.ThemeCategoryQueryRepository
		repository.TouristSpotQueryRepository
		repository.HashtagQueryRepository
		repository.UserQueryRepository
	}
)

var SearchQueryServiceSet = wire.NewSet(
	wire.Struct(new(SearchQueryServiceImpl), "*"),
	wire.Bind(new(SearchQueryService), new(*SearchQueryServiceImpl)),
)

func (s *SearchQueryServiceImpl) ListSuggestionsByKeyward(keyword string) (*entity.SearchSuggetions, error) {
	areaCategories, err := s.AreaCategoryQueryRepository.SearchByName(keyword)
	if err != nil {
		return nil, errors.Wrap(err, "failed search areaCategory list by keyword")
	}
	touristSpots, err := s.TouristSpotQueryRepository.SearchByName(keyword)
	if err != nil {
		return nil, errors.Wrap(err, "failed search tourist_spot list by keyword")
	}
	hashtags, err := s.HashtagQueryRepository.SearchByName(keyword)
	if err != nil {
		return nil, errors.Wrap(err, "failed search hashtag list by keyword")
	}
	users, err := s.UserQueryRepository.SearchByName(keyword)
	if err != nil {
		return nil, errors.Wrap(err, "failed search user list by keyword")
	}

	return s.convertTargetToSuggestion(areaCategories, touristSpots, hashtags, users), nil
}

// 選択された候補のみ返す
// TODO: リファクタ
func (s *SearchQueryServiceImpl) ListSuggestionsByType(keyword string, suggestionType model.SuggestionType) (*entity.SearchSuggetions, error) {
	if suggestionType == model.SuggestionTypeArea {
		areaCategories, err := s.AreaCategoryQueryRepository.SearchByName(keyword)
		if err != nil {
			return nil, errors.Wrap(err, "failed search area list by keyword")
		}
		return s.convertTargetToSuggestion(areaCategories, nil, nil, nil), nil
	}
	if suggestionType == model.SuggestionTypeSubArea {
		areaCategories, err := s.AreaCategoryQueryRepository.SearchSubAreaByName(keyword)
		if err != nil {
			return nil, errors.Wrap(err, "failed search sub_area list by keyword")
		}
		return s.convertTargetToSuggestion(areaCategories, nil, nil, nil), nil
	}
	if suggestionType == model.SuggestionTypeSubSubArea {
		areaCategories, err := s.AreaCategoryQueryRepository.SearchSubSubAreaByName(keyword)
		if err != nil {
			return nil, errors.Wrap(err, "failed search sub_sub_area list by keyword")
		}
		return s.convertTargetToSuggestion(areaCategories, nil, nil, nil), nil
	}
	if suggestionType == model.SuggestionTypeTouristSpot {
		touristSpots, err := s.TouristSpotQueryRepository.SearchByName(keyword)
		if err != nil {
			return nil, errors.Wrap(err, "failed search tourist_spot list by keyword")
		}
		return s.convertTargetToSuggestion(nil, touristSpots, nil, nil), nil
	}
	if suggestionType == model.SuggestionTypeHashTag {
		hashtags, err := s.HashtagQueryRepository.SearchByName(keyword)
		if err != nil {
			return nil, errors.Wrap(err, "failed search hashtag list by keyword")
		}
		return s.convertTargetToSuggestion(nil, nil, hashtags, nil), nil
	}
	if suggestionType == model.SuggestionTypeUser {
		users, err := s.UserQueryRepository.SearchByName(keyword)
		if err != nil {
			return nil, errors.Wrap(err, "failed search user list by keyword")
		}
		return s.convertTargetToSuggestion(nil, nil, nil, users), nil
	}
	return nil, nil
}

// TODO: リファクタ
func (s *SearchQueryServiceImpl) convertTargetToSuggestion(areaCategories []*entity.AreaCategory, touristSpots []*entity.TouristSpot, hashtags []*entity.Hashtag, users []*entity.User) *entity.SearchSuggetions {
	areaSuggestions := make([]*entity.SearchSuggetion, len(areaCategories))
	touristSpotSuggestions := make([]*entity.SearchSuggetion, len(touristSpots))
	hashtagSuggestions := make([]*entity.SearchSuggetion, len(hashtags))
	userSuggestions := make([]*entity.SearchSuggetion, len(users))

	for i, areaCategory := range areaCategories {
		if areaCategory.Type == model.AreaCategoryTypeArea {
			areaSuggestions[i] = &entity.SearchSuggetion{
				ID:   areaCategory.ID,
				Type: model.SuggestionTypeArea,
				Name: areaCategory.Name,
			}
		}
		if areaCategory.Type == model.AreaCategoryTypeSubArea {
			areaSuggestions[i] = &entity.SearchSuggetion{
				ID:   areaCategory.ID,
				Type: model.SuggestionTypeSubArea,
				Name: areaCategory.Name,
			}
		}
		if areaCategory.Type == model.AreaCategoryTypeSubSubArea {
			areaSuggestions[i] = &entity.SearchSuggetion{
				ID:   areaCategory.ID,
				Type: model.SuggestionTypeSubSubArea,
				Name: areaCategory.Name,
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
