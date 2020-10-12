package entity

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	SearchSuggestion struct {
		ID   int                  `json:"id"`
		Type model.SuggestionType `json:"type"`
		Name string               `json:"name"`
	}

	SearchSuggestions struct {
		Area        []*AreaCategory    `json:"area"`
		TouristSpot []*TouristSpotTiny `json:"touristSpot"`
		Hashtag     []*Hashtag         `json:"hashtag"`
		User        []*UserTiny        `json:"user"`
	}
)

func NewSearchSuggestions(area []*AreaCategory, touristSpot []*TouristSpotTiny, hashtag []*Hashtag, user []*UserTiny) *SearchSuggestions {
	return &SearchSuggestions{
		Area:        area,
		TouristSpot: touristSpot,
		Hashtag:     hashtag,
		User:        user,
	}
}
