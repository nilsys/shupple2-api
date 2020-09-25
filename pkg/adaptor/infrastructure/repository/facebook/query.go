package facebook

import (
	"encoding/json"

	"github.com/google/wire"
	"github.com/huandu/facebook/v2"
	"github.com/pkg/errors"
	facebookRepo "github.com/stayway-corp/stayway-media-api/pkg/domain/repository/facebook"
)

type QueryRepositoryImpl struct {
	FacebookSession *facebook.Session
}

var QueryRepositorySet = wire.NewSet(
	wire.Struct(new(QueryRepositoryImpl), "*"),
	wire.Bind(new(facebookRepo.QueryRepository), new(*QueryRepositoryImpl)),
)

func (r *QueryRepositoryImpl) GetShareCountByURL(url string) (int, error) {
	res, err := r.FacebookSession.Get("", map[string]interface{}{
		"id":     url,
		"fields": "engagement",
	})
	if err != nil {
		return 0, errors.Wrap(err, "failed facebook graph api")
	}

	engagement := res["engagement"].(map[string]interface{})

	shareCnt, ok := engagement["share_count"].(json.Number)
	if !ok {
		return 0, errors.New("can't assert engagement.share_count -> json.Number")
	}

	shareCntInt64, err := shareCnt.Int64()
	if err != nil {
		return 0, errors.Wrap(err, "can't assert engagement.share_count.(json.Number) -> int64")
	}

	return int(shareCntInt64), nil
}
