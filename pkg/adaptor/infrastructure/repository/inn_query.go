package repository

import (
	"strconv"

	"github.com/stayway-corp/stayway-media-api/pkg/config"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/client"

	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/dto"

	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	// Inn参照系レポジトリ実装
	InnQueryRepositoryImpl struct {
		config.StaywayConfig
		Client client.Client
	}
)

var InnQueryRepositorySet = wire.NewSet(
	wire.Struct(new(InnQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.InnQueryRepository), new(*InnQueryRepositoryImpl)),
)

const staywayInnAPIPath = "/api/inns"

// AreaID, SubAreaID, SubSubAreaIDから参照したInnのIDのスライスを返す
func (r *InnQueryRepositoryImpl) FindIDsByAreaID(areaId, subAreaId, subSubAreaId int) ([]int, error) {

	if areaId == 0 && subAreaId == 0 && subSubAreaId == 0 {
		return nil, nil
	}

	opts := buildFindIDsByAreaIDAQuery(areaId, subAreaId, subSubAreaId)

	var res dto.Inns
	u := r.StaywayConfig.BaseURL + staywayInnAPIPath
	if err := r.Client.GetJSON(u, opts, &res); err != nil {
		return nil, errors.Wrapf(err, "failed to get inns from stayway api by areaID: %d subAreaID: %d subSubAreaID: %d", areaId, subAreaId, subSubAreaId)
	}

	return res.InnsToIDs(), nil
}

func buildFindIDsByAreaIDAQuery(areaId, subAreaId, subSubAreaId int) *client.Option {
	opts := &client.Option{
		QueryParams: map[string][]string{},
	}

	if areaId != 0 {
		opts.QueryParams.Add("area_id", strconv.Itoa(areaId))
	}

	if subAreaId != 0 {
		opts.QueryParams.Add("sub_area_id", strconv.Itoa(subAreaId))
	}

	if subSubAreaId != 0 {
		opts.QueryParams.Add("sub_sub_are_id", strconv.Itoa(subSubAreaId))
	}

	// MEMO: 仮置き　どの値入れよう。。
	opts.QueryParams.Add("per_page", strconv.Itoa(100))

	return opts
}
