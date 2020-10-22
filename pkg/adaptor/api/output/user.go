package output

import "github.com/uma-co82/shupple2-api/pkg/domain/model"

type (
	User struct {
		ID             int                  `json:"id"`
		Name           string               `json:"name"`
		Birthdate      model.Date           `json:"birthdate"`
		Profile        string               `json:"profile"`
		Gender         model.Gender         `json:"gender"`
		Prefecture     model.Prefecture     `json:"prefecture"`
		MatchingReason model.MatchingReason `json:"matchingReason"`
		IsMatching     bool                 `json:"isMatching"`
		Images         []UserImage          `json:"images"`
	}

	UserImage struct {
		Priority int    `json:"priority"`
		URL      string `json:"url"`
	}
)
