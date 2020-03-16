package param

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
)

// 動画一覧取得パラメータ
type (
	ShowVlog struct {
		ID int `param:"id" validat:"required"`
	}

	ListVlogParam struct {
		AreaID        int               `query:"areaId"`
		SubAreaID     int               `query:"subAreaId"`
		SubSubAreaID  int               `query:"subSubAreaId"`
		TouristSpotID int               `query:"touristSpotId"`
		SortBy        model.MediaSortBy `query:"sortBy"`
		Keyward       string            `query:"q"`
		Page          int               `query:"page"`
		PerPage       int               `query:"perPage"`
	}
)

const listVlogDefaultPerPage = 10

// いずれのクエリも飛んでこない場合 or sortの値が期待値以外の場合エラーを返す
func (param ListVlogParam) Validate() error {
	if param.AreaID == 0 && param.SubAreaID == 0 && param.SubSubAreaID == 0 && param.TouristSpotID == 0 && param.Keyward == "" {
		return serror.New(nil, serror.CodeInvalidParam, "Invalid show vlog list param")
	}

	return nil
}

// PerPageがクエリで飛んで来なかった場合、デフォルト値である10を返す
func (param ListVlogParam) GetLimit() int {
	if param.PerPage == 0 {
		return listVlogDefaultPerPage
	}
	return param.PerPage
}

// offSetを返す(sqlで使う想定)
func (param ListVlogParam) GetOffSet() int {
	if param.Page == 1 || param.Page == 0 {
		return 0
	}
	return param.GetLimit()*(param.Page-1) + 1
}
