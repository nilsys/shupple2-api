package param

type (
	// Area, SubArea, SubSubAreaで共通使用
	GetArea struct {
		ID int `param:"id" validate:"required"`
	}

	// AreaのList取得の際のParams
	ListAreaParams struct {
		AreaGroupID int   `query:"areaGroupId"`
		PerPage     int   `query:"perPage"`
		ExcludeID   []int `query:"excludeId"`
		Page        int   `query:"page"`
	}

	// SubAreaのList取得の際のParams
	ListSubAreaParams struct {
		AreaID    int   `query:"areaId"`
		PerPage   int   `query:"perPage"`
		ExcludeID []int `query:"excludeId"`
		Page      int   `query:"page"`
	}

	// SubSubAreaのList取得の際のParams
	ListSubSubAreaParams struct {
		SubAreaID int   `query:"subAreaId"`
		PerPage   int   `query:"perPage"`
		ExcludeID []int `query:"excludeId"`
		Page      int   `query:"page"`
	}
)
