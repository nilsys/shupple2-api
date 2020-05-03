package input

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
)

type (
	// Area, SubArea, SubSubAreaで共通使用
	GetArea struct {
		ID int `param:"id" validate:"required"`
	}

	// AreaのList取得の際のParams
	ListAreaParams struct {
		AreaGroup model.AreaGroup `query:"areaGroup" validate:"required"`
		PerPage   int             `query:"perPage"`
		ExcludeID []int           `query:"excludeId"`
		Page      int             `query:"page"`
	}

	// SubAreaのList取得の際のParams
	ListSubAreaParams struct {
		AreaID    int   `query:"areaId" validate:"required"`
		PerPage   int   `query:"perPage"`
		ExcludeID []int `query:"excludeId"`
		Page      int   `query:"page"`
	}

	// SubSubAreaのList取得の際のParams
	ListSubSubAreaParams struct {
		SubAreaID int   `query:"subAreaId" validate:"required"`
		PerPage   int   `query:"perPage"`
		ExcludeID []int `query:"excludeId"`
		Page      int   `query:"page"`
	}
)

func (param *ListAreaParams) Validate() error {
	if param.AreaGroup == 0 {
		return serror.New(nil, serror.CodeInvalidParam, "unknown area group")
	}
	return nil
}
