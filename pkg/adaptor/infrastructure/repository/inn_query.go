package repository

import (
	"fmt"
	"path"
	"strconv"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"

	"github.com/stayway-corp/stayway-media-api/pkg/config"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/client"

	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/dto"

	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	// Inn参照系レポジトリ実装
	InnQueryRepositoryImpl struct {
		MetasearchConfig config.StaywayMetasearch
		Client           client.Client
	}
)

var InnQueryRepositorySet = wire.NewSet(
	wire.Struct(new(InnQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.InnQueryRepository), new(*InnQueryRepositoryImpl)),
)

const staywayInnAPIPath = "/api/inns"
const staywayInnNaiveAPIPath = "/api/inns/naive"

// AreaID, SubAreaID, SubSubAreaIDから参照したInnのIDのスライスを返す
func (r *InnQueryRepositoryImpl) FindIDsByAreaID(areaID, subAreaID, subSubAreaID int) ([]int, error) {

	if areaID == 0 && subAreaID == 0 && subSubAreaID == 0 {
		return nil, nil
	}

	opts := buildFindIDsByAreaIDQuery(areaID, subAreaID, subSubAreaID)

	var res dto.Inns

	u := r.MetasearchConfig.BaseURL
	u.Path = path.Join(u.Path, staywayInnNaiveAPIPath)
	if err := r.Client.GetJSON(u.String(), opts, &res); err != nil {
		return nil, errors.Wrapf(err, "failed to get inns from stayway api by areaID: %d subAreaID: %d subSubAreaID: %d", areaID, subAreaID, subSubAreaID)
	}
	return res.InnsToIDs(), nil
}

func (r *InnQueryRepositoryImpl) FindAreaIDsByID(id int) (*entity.InnAreaTypeIDs, error) {
	var res dto.InnArea

	u := r.MetasearchConfig.BaseURL
	u.Path = path.Join(u.Path, staywayInnAPIPath, fmt.Sprintf("/%d/area", id))

	if err := r.Client.GetJSON(u.String(), nil, &res); err != nil {
		return nil, errors.Wrapf(err, "failed to get inn area details from stayway api by innID: %d", id)
	}

	return res.ToInnAreaTypeIDs(), nil
}

func (r *InnQueryRepositoryImpl) FindByParams(query *query.FindInn) (*entity.Inns, error) {
	opts := buildFindByParamQuery(query)

	var res dto.Inns
	u := r.MetasearchConfig.BaseURL
	u.Path = path.Join(u.Path, staywayInnNaiveAPIPath)
	if err := r.Client.GetJSON(u.String(), opts, &res); err != nil {
		return nil, errors.Wrap(err, "failed to get inns from stayway api")
	}

	return res.ConvertToEntity(), nil
}

func buildFindIDsByAreaIDQuery(areaID, subAreaID, subSubAreaID int) *client.Option {
	opts := &client.Option{
		QueryParams: map[string][]string{},
	}

	if areaID != 0 {
		opts.QueryParams.Add("area_id", strconv.Itoa(areaID))
	}

	if subAreaID != 0 {
		opts.QueryParams.Add("sub_area_id", strconv.Itoa(subAreaID))
	}

	if subSubAreaID != 0 {
		opts.QueryParams.Add("sub_sub_are_id", strconv.Itoa(subSubAreaID))
	}

	opts.QueryParams.Add("per_page", strconv.Itoa(defaultAcquisitionNumber))

	return opts
}

func buildFindByParamQuery(query *query.FindInn) *client.Option {
	opts := &client.Option{
		QueryParams: map[string][]string{},
	}

	if query.MetasearchAreaID != 0 {
		opts.QueryParams.Add("area_id", strconv.Itoa(query.MetasearchAreaID))
	}

	if query.MetasearchSubAreaID != 0 {
		opts.QueryParams.Add("sub_area_id", strconv.Itoa(query.MetasearchSubAreaID))
	}

	if query.MetasearchSubSubAreaID != 0 {
		opts.QueryParams.Add("sub_sub_area_id", strconv.Itoa(query.MetasearchSubSubAreaID))
	}
	if query.Latitude != 0 && query.Longitude != 0 {
		opts.QueryParams.Add("geocode", query.GetGeoCode())
	}

	opts.QueryParams.Add("per_page", strconv.Itoa(defaultAcquisitionNumber))

	return opts
}
