package wordpress

//go:generate go-enum -f=$GOFILE --marshal

/*
ENUM(undefined, japan, world, theme)
*/
type CategoryType int

type (
	Category struct {
		ID     int          `json:"id"`
		Name   string       `json:"name"`
		Parent int          `json:"parent"`
		Type   CategoryType `json:"-"`
	}
)
