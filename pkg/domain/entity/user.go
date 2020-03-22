package entity

import (
	"database/sql"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	User struct {
		ID            int `gorm:"primary_key"`
		UID           string
		CognitoID     string
		WordpressID   int
		MigrationCode sql.NullString
		Name          string
		Email         string
		Birthdate     time.Time
		Gender        model.Gender
		Profile       string
		AvatarUUID    string
		Interests     []*UserInterest `gorm:"foreignkey:UserID"`
		CreatedAt     time.Time       `gorm:"-;default:current_timestamp"`
		UpdatedAt     time.Time       `gorm:"-;default:current_timestamp"`
		DeletedAt     *time.Time
	}

	// MEMO: 他でも使う様になったら名前変更
	QueryRankingUser struct {
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
func (user *User) GenerateThumbnailURL() string {
	return "https://files.stayway.jp/avatar/" + user.AvatarUUID
}

func (user *User) IsSelfID(id int) bool {
	return user.ID == id
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
