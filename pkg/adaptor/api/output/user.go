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
		Images         []UserImage          `json:"images"`
	}

	UserImage struct {
		ID       string `json:"id"`
		Priority int    `json:"priority"`
		URL      string `json:"url"`
	}
)
