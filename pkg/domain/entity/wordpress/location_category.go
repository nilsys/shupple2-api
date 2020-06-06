package wordpress

type (
	LocationCategory struct {
		ID     int
		Name   string
		Slug   URLEscapedString `json:"slug"`
		Parent int              `json:"parent"`
		Type   CategoryType     `json:"-"`
	}
)
