package input

import "net/url"

type (
	IDParam struct {
		ID int `param:"id" validate:"required"`
	}

	PathString string

	PaginationQuery struct {
		Page    int `query:"page"`
		PerPage int `query:"perPage"`
	}
)

func (ps *PathString) UnmarshalParam(s string) error {
	var err error
	s, err = url.QueryUnescape(s)
	if err != nil {
		return err
	}
	*ps = PathString(s)
	return nil
}
