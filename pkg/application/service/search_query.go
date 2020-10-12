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
		ListSuggestionsByKeyward(keyword string) (*entity.SearchSuggestions, error)
		ListSuggestionsByType(keyword string, suggestionType model.SuggestionType) (*entity.SearchSuggestions, error)
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

func (s *SearchQueryServiceImpl) ListSuggestionsByKeyward(keyword string) (*entity.SearchSuggestions, error) {
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
	users, err := s.UserQueryRepository.SearchByNameOrUID(keyword)
	if err != nil {
		return nil, errors.Wrap(err, "failed search user list by keyword")
	}

	areaCategoryWithThemeCategory := make([]*entity.AreaCategoryWithThemeCategory, len(*areaCategories))
	for i, areaCategory := range *areaCategories {
		themes, err := s.ThemeCategoryQueryRepository.FindThemesByAreaCategoryID(areaCategory.AreaID, int(areaCategory.SubAreaID.Int64), int(areaCategory.SubSubAreaID.Int64), nil)
		if err != nil {
			return nil, errors.Wrap(err, "failed find theme_category")
		}

		areaCategoryWithThemeCategory[i] = &entity.AreaCategoryWithThemeCategory{
			AreaCategory:    areaCategory,
			ThemeCategories: themes,
		}
	}

	return entity.NewSearchSuggestions(areaCategories, areaCategoryWithThemeCategory, touristSpots, hashtags, users), nil
}

// 選択された候補のみ返す
// TODO: リファクタ
func (s *SearchQueryServiceImpl) ListSuggestionsByType(keyword string, suggestionType model.SuggestionType) (*entity.SearchSuggestions, error) {
	if suggestionType == model.SuggestionTypeArea {
		areaCategories, err := s.AreaCategoryQueryRepository.SearchAreaByName(keyword)
		if err != nil {
			return nil, errors.Wrap(err, "failed search area list by keyword")
		}
		return entity.NewSearchSuggestions(areaCategories, nil, nil, nil, nil), nil
	}
	if suggestionType == model.SuggestionTypeSubArea {
		areaCategories, err := s.AreaCategoryQueryRepository.SearchSubAreaByName(keyword)
		if err != nil {
			return nil, errors.Wrap(err, "failed search sub_area list by keyword")
		}
		return entity.NewSearchSuggestions(areaCategories, nil, nil, nil, nil), nil
	}
	if suggestionType == model.SuggestionTypeSubSubArea {
		areaCategories, err := s.AreaCategoryQueryRepository.SearchSubSubAreaByName(keyword)
		if err != nil {
			return nil, errors.Wrap(err, "failed search sub_sub_area list by keyword")
		}
		return entity.NewSearchSuggestions(areaCategories, nil, nil, nil, nil), nil
	}
	if suggestionType == model.SuggestionTypeTouristSpot {
		touristSpots, err := s.TouristSpotQueryRepository.SearchByName(keyword)
		if err != nil {
			return nil, errors.Wrap(err, "failed search tourist_spot list by keyword")
		}
		return entity.NewSearchSuggestions(nil, nil, touristSpots, nil, nil), nil
	}
	if suggestionType == model.SuggestionTypeHashTag {
		hashtags, err := s.HashtagQueryRepository.SearchByName(keyword)
		if err != nil {
			return nil, errors.Wrap(err, "failed search hashtag list by keyword")
		}
		return entity.NewSearchSuggestions(nil, nil, nil, hashtags, nil), nil
	}
	if suggestionType == model.SuggestionTypeUser {
		users, err := s.UserQueryRepository.SearchByNameOrUID(keyword)
		if err != nil {
			return nil, errors.Wrap(err, "failed search user list by keyword")
		}
		return entity.NewSearchSuggestions(nil, nil, nil, nil, users), nil
	}
	return nil, nil
}
