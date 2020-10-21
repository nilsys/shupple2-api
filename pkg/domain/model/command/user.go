package command

import (
	"time"

	"github.com/uma-co82/shupple2-api/pkg/domain/model"
)

type (
	StoreUser struct {
		Name       string
		Email      string
		Birthdate  time.Time
		Profile    string
		Gender     model.Gender
		Prefecture model.Prefecture
	}
)
