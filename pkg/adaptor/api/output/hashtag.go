package output

type (
	Hashtag struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		IsFollow    bool   `json:"isFollow"`
		PostCount   int    `json:"postCount"`
		ReviewCount int    `json:"reviewCount"`
		Score       int    `json:"score"`
	}
)

func NewHashtag(id int, name string, isFollow bool, postCount, reviewCount, score int) *Hashtag {
	return &Hashtag{
		ID:          id,
		Name:        name,
		IsFollow:    isFollow,
		PostCount:   postCount,
		ReviewCount: reviewCount,
		Score:       score,
	}
}
