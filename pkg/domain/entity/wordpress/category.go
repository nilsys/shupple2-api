package wordpress

type (
	Category struct {
		ID     int              `json:"id"`
		Name   string           `json:"name"`
		Slug   URLEscapedString `json:"slug"`
		Parent int              `json:"parent"`
		Type   CategoryType     `json:"-"`
	}
)
