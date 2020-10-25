package entity

import (
	"fmt"
	"time"

	"gopkg.in/guregu/null.v3"

	"github.com/uma-co82/shupple2-api/pkg/config"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/uma-co82/shupple2-api/pkg/domain/model/command"

	"github.com/uma-co82/shupple2-api/pkg/domain/model"
)

type (
	User struct {
		UserTiny
		Images []*UserImage `gorm:"foreignkey:UserID"`
	}

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
		Times
	}

	UserImage struct {
		UUID     string `gorm:"primary_key"`
		UserID   int
		Priority int
		MimeType string
		TimesWithoutDeletedAt
	}

	UserMatchingHistory struct {
		UserID                int `gorm:"primary_key"`
		MatchingUserID        int `gorm:"primary_key"`
		UserConfirmed         null.Bool
		MatchingUserConfirmed null.Bool
		MatchingExpiredAt     time.Time
		TimesWithoutDeletedAt
	}
)

func (u *User) InsertUserID2Images() {
	for _, image := range u.Images {
		image.UserID = u.ID
	}
}

func (u *UserImage) S3Path() string {
	return fmt.Sprintf("user/%d/%s", u.UserID, u.UUID)
}

func (u *UserImage) URL(filesURL config.URL) string {
	filesURL.Path = u.S3Path()
	return filesURL.String()
}

func (u *UserMatchingHistory) IsExpired() bool {
	return u.MatchingExpiredAt.Before(time.Now())
}

/*
	constructor
*/
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
	}
}

func NewUserMatchingHistory(userID, matchingUserID int, now time.Time) *UserMatchingHistory {
	return &UserMatchingHistory{
		UserID:            userID,
		MatchingUserID:    matchingUserID,
		MatchingExpiredAt: now.Add(24 * time.Hour),
	}
}

func NewUser(cmd command.StoreUser, firebaseID string) (*User, error) {
	images := make([]*UserImage, len(cmd.Images))
	for i, image := range images {
		uuid, err := uuid.NewRandom()
		if err != nil {
			return nil, errors.Wrap(err, "failed gen uuid")
		}

		images[i] = &UserImage{
			UUID:     uuid.String(),
			Priority: image.Priority,
			MimeType: image.MimeType,
		}
	}
	return &User{
		UserTiny: *NewUserTinyFromCmd(cmd, firebaseID),
		Images:   images,
	}, nil
}
