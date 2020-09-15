package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

func (c Converters) ConvertSearchSuggestionsToOutput(suggestions *entity.SearchSuggestions) *output.SearchSuggestions {
	areaCategories := make([]*output.AreaCategory, len(suggestions.Area))
	touristSpots := make([]*output.TouristSpotTiny, len(suggestions.TouristSpot))
	hashtags := make([]*output.Hashtag, len(suggestions.Hashtag))
	users := make([]*output.UserSummary, len(suggestions.User))

	for i, areaCategorie := range suggestions.Area {
		areaCategories[i] = c.ConvertAreaCategoryToOutput(areaCategorie)
	}

	for i, touristSpot := range suggestions.TouristSpot {
		touristSpots[i] = c.ConvertTouristSpotTinyToOutput(touristSpot)
	}

	for i, hashtag := range suggestions.Hashtag {
		// MEMO: 現状isFollowを必要としない為第３引数にfalseを入れている
		hashtags[i] = output.NewHashtag(hashtag.ID, hashtag.Name, false)
	}

	for i, user := range suggestions.User {
		users[i] = c.convertUserTinyToOutput(user)
	}

	return &output.SearchSuggestions{
		Area:        areaCategories,
		TouristSpot: touristSpots,
		Hashtag:     hashtags,
		User:        users,
	}
}
