package command

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	CreateReview struct {
		TravelDate    model.YearMonth
		Accompanying  model.AccompanyingType
		TouristSpotID int
		InnID         int
		Score         int
		Body          string
		MediaUUIDs    []*CreateReviewMedia
	}

	CreateReviewMedia struct {
		UUID     string
		MimeType string
	}
)
