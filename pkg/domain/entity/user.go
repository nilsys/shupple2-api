package entity

import (
	"time"

	"github.com/uma-co82/shupple2-api/pkg/domain/model/command"

	"github.com/uma-co82/shupple2-api/pkg/domain/model"
)

type (
	UserTiny struct {
		ID             int `gorm:"primary_key"`
		FirebaseID     string
		Name           string
		Email          string
		Birthdate      time.Time
		Profile        string
		Gender         model.Gender
		Prefecture     model.Prefecture
		MatchingReason model.MatchingReason
		IsMatching     bool
	}

	UserMatchingHistory struct {
		UserID         int `gorm:"primary_key"`
		MatchingUserID int `gorm:"primary_key"`
	}
)

func NewUserTinyFromCmd(cmd command.StoreUser, firebaseID string) *UserTiny {
	return &UserTiny{
		FirebaseID:     firebaseID,
		Name:           cmd.Name,
		Email:          cmd.Email,
		Birthdate:      cmd.Birthdate,
		Profile:        cmd.Profile,
		Gender:         cmd.Gender,
		Prefecture:     cmd.Prefecture,
		MatchingReason: cmd.MatchingReason,
		IsMatching:     false,
	}
}

func NewUserMatchingHistory(userID, matchingUserID int) *UserMatchingHistory {
	return &UserMatchingHistory{
		UserID:         userID,
		MatchingUserID: matchingUserID,
	}
}
