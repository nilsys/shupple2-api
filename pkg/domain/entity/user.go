package entity

import (
	"fmt"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"gopkg.in/guregu/null.v3"
)

type (
	User struct {
		ID            int `gorm:"primary_key"`
		UID           string
		CognitoID     string
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
		CreatedAt     time.Time       `gorm:"-;default:current_timestamp"`
		UpdatedAt     time.Time       `gorm:"-;default:current_timestamp"`
		DeletedAt     *time.Time
	}

	OptionalUser struct {
		User
		Authenticated bool
	}

	UserDetailWithCount struct {
		UserDetail
		FollowingCount int
		FollowerCount  int
		PostCount      int
		ReviewCount    int
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

	Interest struct {
		ID   int    `gorm:"primary_key" json:"id"`
		Name string `json:"name"`
	}
)

// MEMO: サムネイルロジック仮置き
func (u *User) IconURL() string {
	return "https://files.stayway.jp/avatar/" + u.AvatarUUID
}

func (u *User) HeaderURL() string {
	return "https://files.stayway.jp/avatar/" + u.HeaderUUID
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

func (q *UserDetail) TableName() string {
	return "user"
}

func (u *User) S3AvatarPath() string {
	return fmt.Sprintf("user/%s", u.AvatarUUID)
}

func (u *User) S3HeaderPath() string {
	return fmt.Sprintf("user/%s", u.HeaderUUID)
}

func (u *OptionalUser) IsAuthorized() bool {
	return u.Authenticated
}
