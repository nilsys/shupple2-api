package output

type (
	SearchSuggestions struct {
		Area        []*AreaCategory    `json:"area"`
		TouristSpot []*TouristSpotTiny `json:"touristSpot"`
		Hashtag     []*Hashtag         `json:"hashtag"`
		User        []*UserSummary     `json:"user"`
	}
)
