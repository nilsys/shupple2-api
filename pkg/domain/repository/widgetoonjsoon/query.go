package widgetoonjsoon

import widgetoonJsoonDto "github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/dto/widgetoonjsoon"

type (
	QueryRepository interface {
		TwitterCountByURL(url string) (*widgetoonJsoonDto.TwitterCount, error)
	}
)