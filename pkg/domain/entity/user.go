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
		ID            int `gorm:"primary_key"`
		UID           string
		CognitoID     null.String
		WordpressID   null.Int
		MigrationCode null.String
		Name          string
		Email         string
		Birthdate     time.Time
		Gender        model.Gender
		Profile       string
		AvatarUUID    string
		HeaderUUID    string
		URL           string
		FacebookURL   string
		InstagramURL  string
		TwitterURL    string
		YoutubeURL    string
		LivingArea    string
		Interests     []*UserInterest `gorm:"foreignkey:UserID"`
		Times
	}

	UserTable struct {
		ID            int `gorm:"primary_key"`
		UID           string
		CognitoID     null.String
		WordpressID   null.Int
		MigrationCode null.String
		Name          string
		Email         string
		Birthdate     time.Time
		Gender        model.Gender
		Profile       string
		AvatarUUID    string
		HeaderUUID    string
		URL           string
		FacebookURL   string
		InstagramURL  string
		TwitterURL    string
		YoutubeURL    string
		LivingArea    string
		Times
	}

	OptionalUser struct {
		User
		Authenticated bool
	}

	UserDetailWithMediaCount struct {
		UserDetail
		FollowingCount int
		FollowerCount  int
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
	filesURL.Path = u.S3AvatarPath()
	return filesURL.String()
}

func (u *User) HeaderURL(filesURL config.URL) string {
	if u.HeaderUUID == "" {
		return ""
	}
	filesURL.Path = u.S3HeaderPath()
	return filesURL.String()
}

func (u *UserTable) AvatarURL(filesURL config.URL) string {
	if u.HeaderUUID == "" {
		return ""
	}
	filesURL.Path = u.S3AvatarPath()
	return filesURL.String()
}

func (u *UserTable) S3AvatarPath() string {
	return fmt.Sprintf("user/%s", u.AvatarUUID)
}

func (u *User) S3AvatarPath() string {
	return fmt.Sprintf("user/%s", u.AvatarUUID)
}

func (u *User) S3HeaderPath() string {
	return fmt.Sprintf("user/%s", u.HeaderUUID)
}

func (u *User) IsSelfID(id int) bool {
	return u.ID == id
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

func (u *UserTable) TableName() string {
	return "user"
}

func (q *UserDetail) TableName() string {
	return "user"
}

func (u *OptionalUser) IsAuthorized() bool {
	return u.Authenticated
}
