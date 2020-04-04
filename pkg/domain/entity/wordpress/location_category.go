package wordpress

type (
	LocationCategory struct {
		ID     int
		Name   string
		Slug   string       `json:"slug"`
		Parent int          `json:"parent"`
		Type   CategoryType `json:"-"`
	}
)
