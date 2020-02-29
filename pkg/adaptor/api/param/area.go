package param

type (
	AreaID struct {
		ID int `param:"id" validate:"required"`
	}

	GetArea struct {
		ParentCategoryID int   `query:"parentCategoryId"`
		PerPage          int   `query:"perPage"`
		ExcludeID        []int `query:"excludeId"`
		Page             int   `query:"page"`
	}
)
