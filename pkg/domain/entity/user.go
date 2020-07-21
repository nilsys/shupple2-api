package entity

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"gopkg.in/guregu/null.v3"
)

type (
	User struct {
		ID              int `gorm:"primary_key"`
		UID             string
		CognitoID       null.String
		WordpressID     null.Int
		CognitoUserName string
		MigrationCode   null.String
		Name            string
		Email           string
		Birthdate       time.Time
		Gender          model.Gender
		Profile         string
		AvatarUUID      string
		HeaderUUID      string
		URL             string
		FacebookURL     string
		InstagramURL    string
		TwitterURL      string
		YoutubeURL      string
		LivingArea      string
		UserInterests   []*UserInterest `gorm:"foreignkey:UserID"`
		Times
	}

	UserTiny struct {
		ID              int `gorm:"primary_key"`
		UID             string
		CognitoID       null.String
		WordpressID     null.Int
		CognitoUserName string
		MigrationCode   null.String
		Name            string
		Email           string
		Birthdate       time.Time
		Gender          model.Gender
		Profile         string
		AvatarUUID      string
		HeaderUUID      string
		URL             string
		FacebookURL     string
		InstagramURL    string
		TwitterURL      string
		YoutubeURL      string
		LivingArea      string
		Times
	}

	UserTinyList struct {
		List []*UserTiny
	}

	OptionalUser struct {
		User
		Authenticated bool
	}

	UserDetailWithMediaCount struct {
		UserDetail
		FollowingCount int
		FollowedCount  int
		PostCount      int
		ReviewCount    int
		VlogCount      int
		IsFollow       bool
	}

	UserDetail struct {
		User
		Interests []*Interest `gorm:"many2many:user_interest;jointable_foreignkey:user_id;"`
	}

	UserInterest struct {
		UserID     int `gorm:"primary_key"`
		InterestID int `gorm:"primary_key"`
	}

	UserFollowHashtag struct {
		UserID    int `gorm:"primary_key"`
		HashtagID int `gorm:"primary_key"`
	}

	UserFollowed struct {
		// フォローされた人
		UserID int `gorm:"primary_key"`
		// フォローした人
		TargetID int `gorm:"primary_key"`
	}

	UserFollowing struct {
		// フォローした人
		UserID int `gorm:"primary_key"`
		// フォローされた人
		TargetID int `gorm:"primary_key"`
	}
)

func NewUserByWordpressUser(wpUser *wordpress.User) *User {
	return &User{
		UID:           string(wpUser.Slug),
		Name:          wpUser.Name,
		MigrationCode: null.StringFrom(uuid.NewV4().String()),
		WordpressID:   null.IntFrom(int64(wpUser.ID)),
		Profile:       wpUser.Description,
		Birthdate:     time.Date(1900, 1, 1, 0, 0, 0, 0, time.Local),
		FacebookURL:   wpUser.Meta.Facebook,
		TwitterURL:    wpUser.Meta.Twitter,
		InstagramURL:  wpUser.Meta.Instagram,
		YoutubeURL:    wpUser.Meta.Youtube,
	}
}

func (u *User) ConvertToOptionalUser() *OptionalUser {
	return &OptionalUser{
		User:          *u,
		Authenticated: true,
	}
}

func (u *User) PatchByWordpressUser(wpUser *wordpress.User) {
	u.Name = wpUser.Name
	u.Profile = wpUser.Description
	u.FacebookURL = wpUser.Meta.Facebook
	u.TwitterURL = wpUser.Meta.Twitter
	u.InstagramURL = wpUser.Meta.Instagram
	u.YoutubeURL = wpUser.Meta.Youtube
}

// MEMO: サムネイルロジック仮置き
func (u *User) AvatarURL(filesURL config.URL) string {
	if u.AvatarUUID == "" {
		return ""
	}
	filesURL.Path = model.UserS3Path(u.AvatarUUID)
	return filesURL.String()
}

func (u *User) HeaderURL(filesURL config.URL) string {
	if u.HeaderUUID == "" {
		return ""
	}
	filesURL.Path = model.UserS3Path(u.HeaderUUID)
	return filesURL.String()
}

func (u *UserTiny) AvatarURL(filesURL config.URL) string {
	if u.HeaderUUID == "" {
		return ""
	}
	filesURL.Path = model.UserS3Path(u.AvatarUUID)
	return filesURL.String()
}

func (u *User) IsSelfID(id int) bool {
	return u.ID == id
}

func (u *UserTinyList) Emails() []string {
	resolve := make([]string, len(u.List))
	for i, user := range u.List {
		resolve[i] = user.Email
	}
	return resolve
}

func NewUserFollowing(userID, targetID int) *UserFollowing {
	return &UserFollowing{
		UserID:   userID,
		TargetID: targetID,
	}
}

func NewUserFollowed(userID, targetID int) *UserFollowed {
	return &UserFollowed{
		UserID:   userID,
		TargetID: targetID,
	}
}

func NewUserFollowHashtag(userID, hashtagID int) *UserFollowHashtag {
	return &UserFollowHashtag{
		UserID:    userID,
		HashtagID: hashtagID,
	}
}

func (u *UserTiny) TableName() string {
	return "user"
}

func (q *UserDetail) TableName() string {
	return "user"
}

func (u *OptionalUser) IsAuthorized() bool {
	return u.Authenticated
}

func (u *User) PayjpCustomerID() string {
	return fmt.Sprintf("sw_%s", u.UID)
}
