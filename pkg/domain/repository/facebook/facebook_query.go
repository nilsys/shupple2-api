package facebook

type (
	QueryRepository interface {
		GetShareCountByURL(url string) (int, error)
	}
)
