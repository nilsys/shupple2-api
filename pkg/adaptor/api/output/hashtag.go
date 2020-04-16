package output

type (
	Hashtag struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

func NewHashtag(id int, name string) *Hashtag {
	return &Hashtag{
		ID:   id,
		Name: name,
	}
}
