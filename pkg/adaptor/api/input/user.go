package input

import "github.com/uma-co82/shupple2-api/pkg/domain/model"

type (
	RegisterUser struct {
		FirebaseToken  string               `json:"firebaseToken"`
		Name           string               `json:"name"`
		Email          string               `json:"email"`
		Birthdate      model.Date           `json:"birthdate"`
		Profile        string               `json:"profile"`
		Gender         model.Gender         `json:"gender"`
		Prefecture     model.Prefecture     `json:"prefecture"`
		MatchingReason model.MatchingReason `json:"matchingReason"`
		Images         []RegisterUserImage  `json:"images"`
	}

	RegisterUserImage struct {
		Priority    int    `json:"priority"`
		MimeType    string `json:"mimeType"`
		ImageBase64 string `json:"imageBase64"`
	}

	ConfirmMatching struct {
		MatchingUserID IDParam
	}
)
