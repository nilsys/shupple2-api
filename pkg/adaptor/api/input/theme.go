package input

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"

type (
	ListThemeParams struct {
		ExcludeID    []int `query:"excludeId"`
		AreaID       int   `query:"areaId"`
		SubAreaID    int   `query:"subAreaId"`
		SubSubAreaID int   `query:"subSubAreaId"`
	}

	ListSubThemeParams struct {
		ThemeID      int   `validate:"required" query:"themeId"`
		ExcludeID    []int `query:"excludeId"`
		AreaID       int   `query:"areaId"`
		SubAreaID    int   `query:"subAreaId"`
		SubSubAreaID int   `query:"subSubAreaId"`
	}
)

func (p *ListThemeParams) GetAreaCategoryID() int {
	if p.AreaID != 0 {
		return p.AreaID
	}
	if p.SubAreaID != 0 {
		return p.SubAreaID
	}
	if p.SubSubAreaID != 0 {
		return p.SubSubAreaID
	}
	return 0
}

func (p *ListSubThemeParams) GetAreaCategoryID() int {
	if p.AreaID != 0 {
		return p.AreaID
	}
	if p.SubAreaID != 0 {
		return p.SubAreaID
	}
	if p.SubSubAreaID != 0 {
		return p.SubSubAreaID
	}
	return 0
}

func (p *ListThemeParams) Validate() error {
	if (p.AreaID != 0 && p.SubAreaID != 0) || (p.SubAreaID != 0 && p.SubSubAreaID != 0) || (p.AreaID != 0 && p.SubSubAreaID != 0) {
		return serror.New(nil, serror.CodeInvalidParam, "invalid param")
	}
	return nil
}

func (p *ListSubThemeParams) Validate() error {
	if (p.AreaID != 0 && p.SubAreaID != 0) || (p.SubAreaID != 0 && p.SubSubAreaID != 0) || (p.AreaID != 0 && p.SubSubAreaID != 0) {
		return serror.New(nil, serror.CodeInvalidParam, "invalid param")
	}
	return nil
}
