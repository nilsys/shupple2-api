package wordpress

type (
	Category struct {
		ID     int          `json:"id"`
		Name   string       `json:"name"`
		Parent int          `json:"parent"`
		Type   CategoryType `json:"-"`
	}
)
