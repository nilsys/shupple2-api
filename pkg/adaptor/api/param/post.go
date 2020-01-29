package param

type (
	GetPost struct {
		ID int `param:"id" validate:"required"`
	}

	StorePostParam struct {
		Title string
		Body string
	}
)
