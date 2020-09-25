package widgetoonjsoon

import (
	"fmt"

	widgetoonJsoonDto "github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/dto/widgetoonjsoon"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/client"
	widgetoonJsoon "github.com/stayway-corp/stayway-media-api/pkg/domain/repository/widgetoonjsoon"
)

type QueryRepositoryImpl struct {
	Client client.Client
}

const widgetoonJsoonBaseURL = "https://jsoon.digitiminimi.com/twitter/count.json"

var QueryRepositorySet = wire.NewSet(
	wire.Struct(new(QueryRepositoryImpl), "*"),
	wire.Bind(new(widgetoonJsoon.QueryRepository), new(*QueryRepositoryImpl)),
)

func (r *QueryRepositoryImpl) GetTwitterCountByURL(url string) (*widgetoonJsoonDto.TwitterCount, error) {
	var res widgetoonJsoonDto.TwitterCount
	opts := &client.Option{
		QueryParams: map[string][]string{},
	}
	opts.QueryParams.Add("url", fmt.Sprintf("\"%s\"", url))

	// MEMO: https://jsoon.digitiminimi.com/usr-f7501b98e2d7d55f36737605af6d963e
	if err := r.Client.GetJSON(widgetoonJsoonBaseURL, opts, &res); err != nil {
		return nil, errors.Wrap(err, "failed get count from widgetoon.jp&count.jsoon")
	}

	return &res, nil
}
