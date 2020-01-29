package repository

import (
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

func newUser(id int) *entity.User {
	user := &entity.User{
		ID:        id,
		Birthdate: time.Date(2020, 7, 7, 0, 0, 0, 0, time.Local),
		Gender:    entity.GenderMale,
	}
	util.FillDymmyString(user, id)
	return user
}
