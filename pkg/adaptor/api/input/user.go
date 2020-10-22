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
	}
)
