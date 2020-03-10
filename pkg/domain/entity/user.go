package entity

import (
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	User struct {
		ID          int `gorm:"primary_key"`
		CognitoID   string
		WordpressID int
		Name        string
		Email       string
		Birthdate   time.Time
		Gender      model.Gender
		Profile     string
		AvatarUUID  string
		CreatedAt   time.Time `gorm:"-;default:current_timestamp"`
		UpdatedAt   time.Time `gorm:"-;default:current_timestamp"`
		DeletedAt   *time.Time
	}

	// MEMO: 他でも使う様になったら名前変更
	QueryRankingUser struct {
		User
		Interests []*Interest `gorm:"many2many:user_interest;jointable_foreignkey:user_id;"`
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
