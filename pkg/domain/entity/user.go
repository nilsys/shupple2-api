package entity

import (
	"fmt"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/util"

	"github.com/pkg/errors"

	uuid "github.com/satori/go.uuid"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"gopkg.in/guregu/null.v3"
)

type (
	User struct {
		UserTiny
		UserInterests []*UserInterest `gorm:"foreignkey:UserID"`
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
		DeviceToken     null.String
		AssociateID     null.String
		IsNonLogin      bool
		Times
		// 如何なる場面でも必要な為Userの最小単位に置いておく
		UserAttributes []*UserAttribute `gorm:"foreignkey:UserID"`
	}

	UserTinyList struct {
		List []*UserTiny
	}

	UserTinyWithIsFollow struct {
		UserTiny
		IsFollow bool
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
		IsBlocking     bool
	}

	UserDetail struct {
		User
		Interests []*Interest `gorm:"many2many:user_interest;jointable_foreignkey:user_id;"`
	}

	UserInterest struct {
		UserID     int `gorm:"primary_key"`
		InterestID int `gorm:"primary_key"`
	}

	UserBlockUser struct {
		UserID        int `gorm:"primary_key"`
		BlockedUserID int `gorm:"primary_key"`
	}

	UserFollowHashtag struct {
		UserID    int `gorm:"primary_key"`
		HashtagID int `gorm:"primary_key"`
	}

	UserFollowHashtags []*UserFollowHashtag

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

	UserFollowings []*UserFollowing

	UserBlocking []*UserBlockUser

	UserAttribute struct {
		UserID    int                 `gorm:"primary_key"`
		Attribute model.UserAttribute `gorm:"primary_key"`
	}

	// それぞれのmapを持った値オブジェクト
	UserRelationFlgMap struct {
		IDIsFollowMap   map[int]bool
		IDIsBlockingMap map[int]bool
	}
	UserInn struct {
		UserID int `gorm:"primary_key"`
		InnID  int `gorm:"primary_key"`
	}

	UserTouristSpot struct {
		UserID        int `gorm:"primary_key"`
		TouristSpotID int `gorm:"primary_key"`
	}
)

const (
	nonLoginUserUIDLength = 12
)

func NewUserByWordpressUser(wpUser *wordpress.User) *User {
	return &User{
		UserTiny: UserTiny{
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
		},
	}
}

func NewUserBlock(userID, blockedUserID int) *UserBlockUser {
	return &UserBlockUser{
		UserID:        userID,
		BlockedUserID: blockedUserID,
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
	if u.AvatarUUID == "" {
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

func (u *UserTiny) PayjpCustomerID() string {
	return fmt.Sprintf("sw_%d", u.ID)
}

// TODO:
func (u *UserTiny) LoginWithMigrationCodeURL() string {
	return fmt.Sprintf("https://stayway.jp/users/signup?migration_code=%s", u.MigrationCode.String)
}

func (u *UserTiny) AddAttribute(attr model.UserAttribute) {
	u.UserAttributes = append(u.UserAttributes, &UserAttribute{
		Attribute: attr,
	})
}

// id:IsExist のMapを返す
func (u UserFollowings) ToIDExistMap(userIDs []int) map[int]bool {
	resolve := make(map[int]bool, len(userIDs))

	for _, id := range userIDs {
		for _, tiny := range u {
			if id == tiny.TargetID {
				resolve[id] = true
			}
		}
	}

	return resolve
}

// id:IsExist のMapを返す
func (u UserBlocking) ToIDExistMap(userIDs []int) map[int]bool {
	resolve := make(map[int]bool, len(userIDs))

	for _, id := range userIDs {
		for _, tiny := range u {
			if id == tiny.BlockedUserID {
				resolve[id] = true
			}
		}
	}

	return resolve
}

func (u UserFollowHashtags) ToIDExistMap(userIDs []int) map[int]bool {
	resolve := make(map[int]bool, len(userIDs))

	for _, id := range userIDs {
		for _, tiny := range []*UserFollowHashtag(u) {
			if id == tiny.HashtagID {
				resolve[id] = true
			} else {
				resolve[id] = false
			}
		}
	}

	return resolve
}

// 非ログインユーザー
func NewIsNonLoginUserTiny(name string) (*UserTiny, error) {
	now := time.Now()

	migrationCode, err := model.NewRandUUID()
	if err != nil {
		return nil, errors.Wrap(err, "failed gen uuid")
	}

	uid, err := model.RandomStr(nonLoginUserUIDLength)
	if err != nil {
		return nil, errors.Wrap(err, "failed gen rand str")
	}

	return &UserTiny{
		Name:          name,
		UID:           uid,
		MigrationCode: null.StringFrom(migrationCode),
		Birthdate:     time.Date(1000, 1, 1, 0, 0, 0, 0, util.JSTLoc),
		IsNonLogin:    true,
		Times: Times{
			DeletedAt: &now,
		},
	}, nil
}

func (u *UserTiny) MediaWebURL(baseURL config.URL) *config.URL {
	baseURL.Path = fmt.Sprintf("/users/%s", u.UID)
	return &baseURL
}

func (u *UserRelationFlgMap) IsFollowByUserID(targetUserID int) bool {
	return u.IDIsFollowMap[targetUserID]
}

func (u *UserRelationFlgMap) IsBlockingByUserID(targetUserID int) bool {
	return u.IDIsBlockingMap[targetUserID]
}
