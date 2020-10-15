package entity

import (
	"time"

	"github.com/uma-co82/shupple2-api/pkg/domain/model"
)

type (
	UserTiny struct {
		ID         int `gorm:"primary_key"`
		FirebaseID string
		Name       string
		Email      string
		Birthdate  time.Time
		Profile    string
		Gender     model.Gender
	}
)
