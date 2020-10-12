package entity

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	SearchSuggestion struct {
		ID   int                  `json:"id"`
		Type model.SuggestionType `json:"type"`
		Name string               `json:"name"`
	}

	SearchSuggestions struct {
		Area                          *AreaCategories
		AreaCategoryWithThemeCategory []*AreaCategoryWithThemeCategory
		TouristSpot                   []*TouristSpotTiny
		Hashtag                       []*Hashtag
		User                          []*UserTiny
	}

	// value object
	AreaCategoryWithThemeCategory struct {
		AreaCategory    *AreaCategory
		ThemeCategories []*ThemeCategoryWithPostCount
	}
)

func NewSearchSuggestions(area *AreaCategories, areaCategoryWithThemeCategory []*AreaCategoryWithThemeCategory, touristSpot []*TouristSpotTiny, hashtag []*Hashtag, user []*UserTiny) *SearchSuggestions {
	return &SearchSuggestions{
		Area:                          area,
		AreaCategoryWithThemeCategory: areaCategoryWithThemeCategory,
		TouristSpot:                   touristSpot,
		Hashtag:                       hashtag,
		User:                          user,
	}
}
