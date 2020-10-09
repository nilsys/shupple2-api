package facebook

import (
	"encoding/json"

	facebookEntity "github.com/stayway-corp/stayway-media-api/pkg/domain/entity/facebook"

	"github.com/google/wire"
	"github.com/huandu/facebook/v2"
	"github.com/pkg/errors"
	facebookRepo "github.com/stayway-corp/stayway-media-api/pkg/domain/repository/facebook"
)

type (
	QueryRepositoryImpl struct {
		FacebookSession *facebook.Session
	}
)

var QueryRepositorySet = wire.NewSet(
	wire.Struct(new(QueryRepositoryImpl), "*"),
	wire.Bind(new(facebookRepo.QueryRepository), new(*QueryRepositoryImpl)),
)

func (r *QueryRepositoryImpl) GetShareCountByURLBatchRequest(query []facebook.Params) (facebookEntity.EngagementAndIDList, error) {
	results, err := r.FacebookSession.Batch(facebook.Params{"include_headers": false}, query...)
	if err != nil {
		return nil, errors.Wrap(err, "failed facebook graph api")
	}

	var res []*facebookEntity.EngagementAndID

	for _, result := range results {
		var resolve facebookEntity.EngagementAndID
		body, ok := result.Get("body").(string)
		if !ok {
			return nil, errors.New("can't cast facebook response body")
		}
		if err := json.Unmarshal([]byte(body), &resolve); err != nil {
			return nil, errors.Wrap(err, "failed cast")
		}
		res = append(res, &resolve)
	}

	return res, nil
}
