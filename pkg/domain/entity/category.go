package entity

type (
	CategoryBase struct {
		ID   int `gorm:"primary_key"`
		Name string
		Slug string
	}

	Category interface {
		CategoryID() int
		CategoryName() string
		CategorySlug() string
		CategoryType() string
	}
)

func (c CategoryBase) CategoryID() int {
	return c.ID
}

func (c CategoryBase) CategoryName() string {
	return c.Name
}

func (c CategoryBase) CategorySlug() string {
	return c.Slug
}
