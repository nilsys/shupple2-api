package response

type (
	Category struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

func NewCategory(id int, name string) Category {
	return Category{
		ID:   id,
		Name: name,
	}
}
