package output

type (
	Hashtag struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		IsFollow bool   `json:"isFollow"`
	}
)

func NewHashtag(id int, name string, isFollow bool) *Hashtag {
	return &Hashtag{
		ID:       id,
		Name:     name,
		IsFollow: isFollow,
	}
}
