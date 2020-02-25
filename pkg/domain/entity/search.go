package entity

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	SearchSuggetion struct {
		ID   int                  `json:"id"`
		Type model.SuggestionType `json:"type"`
		Name string               `json:"name"`
	}

	SearchSuggetions struct {
		Area        []*SearchSuggetion `json:"area"`
		TouristSpot []*SearchSuggetion `json:"touristSpot"`
		Hashtag     []*SearchSuggetion `json:"hashtag"`
		User        []*SearchSuggetion `json:"user"`
	}
)

func NewSearchSuggestions(area, touristSpot, hashtag, user []*SearchSuggetion) *SearchSuggetions {
	return &SearchSuggetions{
		Area:        area,
		TouristSpot: touristSpot,
		Hashtag:     hashtag,
		User:        user,
	}
}
