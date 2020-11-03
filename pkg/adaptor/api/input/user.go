package input

import "github.com/uma-co82/shupple2-api/pkg/domain/model"

type (
	RegisterUser struct {
		FirebaseToken  string               `json:"firebaseToken" validate:"required"`
		Name           string               `json:"name" validate:"required"`
		Email          string               `json:"email" validate:"required"`
		Birthdate      model.Date           `json:"birthdate" validate:"required"`
		Profile        string               `json:"profile" validate:"required"`
		Gender         model.Gender         `json:"gender" validate:"required"`
		Prefecture     model.Prefecture     `json:"prefecture" validate:"required"`
		MatchingReason model.MatchingReason `json:"matchingReason" validate:"required"`
	}

	RegisterUserImage struct {
		Priority    int    `json:"priority"`
		MimeType    string `json:"mimeType"`
		ImageBase64 string `json:"imageBase64"`
	}

	ReviewMainMatching struct {
		MatchingUserID IDParam
	}

	DeleteUserImage struct {
		ID string `param:"id"`
	}
)
