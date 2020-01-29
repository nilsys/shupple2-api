package entity

import "time"

type (
	User struct {
		ID          int `gorm:"primary_key"`
		CognitoID   string
		WordpressID int
		Name        string
		Email       string
		Birthdate   time.Time
		Gender      Gender
		Profile     string
		AvatarUUID  string
		CreatedAt   time.Time `gorm:"default:current_timestamp"`
		UpdatedAt   time.Time `gorm:"default:current_timestamp"`
		DeletedAt   *time.Time
	}
)
