package entity

type (
	Hashtag struct {
		ID          int `gorm:"primary_key"`
		Name        string
		PostCount   int
		ReviewCount int
		Score       int
	}

	Hashtags []*Hashtag
)

func (h Hashtags) IDs() []int {
	resolve := make([]int, len([]*Hashtag(h)))

	for i, tiny := range []*Hashtag(h) {
		resolve[i] = tiny.ID
	}

	return resolve
}
