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
)

// TODO: サムネイル生成ロジック
func (user *User) GenerateThumbnailURL() string {
	return "foo"
}
