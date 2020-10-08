package model

func (x MediaType) IsUserImage() bool {
	return x == MediaTypeUserIcon || x == MediaTypeUserHeader
}
