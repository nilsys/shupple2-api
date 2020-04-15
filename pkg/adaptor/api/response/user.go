package response

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	Creator struct {
		ID           int    `json:"id"`
		UID          string `json:"uid"`
		Thumbnail    string `json:"iconUrl"`
		Name         string `json:"name"`
		Profile      string `json:"profile"`
		FacebookURL  string `json:"facebookUrl"`
		InstagramURL string `json:"instagramUrl"`
		TwitterURL   string `json:"twitterUrl"`
		YoutubeURL   string `json:"youtubeUrl"`
	}

	// ユーザーランキングで返すレスポンス型
	RankinUser struct {
		ID        int      `jso:"id"`
		UID       string   `json:"uid"`
		Name      string   `json:"name"`
		Profile   string   `json:"profile"`
		Thumbnail string   `json:"iconUrl"`
		Interests []string `json:"interests"`
	}

	// MEMO: ユースケースが増えれば命名返る
	FollowUser struct {
		ID        int    `json:"id"`
		UID       string `json:"uid"`
		Name      string `json:"name"`
		Thumbnail string `json:"iconUrl"`
	}

	UserSummary struct {
		ID   int    `json:"id"`
		UID  string `json:"uid"`
		Name string `json:"name"`
		Icon string `json:"iconUrl"`
	}

	MyPageUser struct {
		ID             int          `json:"id"`
		UID            string       `json:"uid"`
		Name           string       `json:"name"`
		Profile        string       `json:"profile"`
		Birthdate      model.Date   `json:"birthdate"`
		Email          string       `json:"email"`
		Gender         model.Gender `json:"gender"`
		Icon           string       `json:"iconUrl"`
		Header         string       `json:"headerUrl"`
		Facebook       string       `json:"facebook"`
		Instagram      string       `json:"instagram"`
		Twitter        string       `json:"twitter"`
		LivingArea     string       `json:"livingArea"`
		PostCount      int          `json:"postCount"`
		FollowingCount int          `json:"followingCount"`
		FollowedCount  int          `json:"followedCount"`
		Interests      []string     `json:"interests"`
	}
)

func NewCreatorFromUser(user *entity.User) Creator {
	return NewCreator(
		user.ID, user.UID, user.IconURL(), user.Name, user.Profile,
		user.FacebookURL, user.InstagramURL, user.TwitterURL, user.YoutubeURL,
	)
}

func NewCreator(id int, uid, thumbnail, name, profile, facebookURL, instagramURL, twitterURL, youtubeURL string) Creator {
	return Creator{
		ID:           id,
		UID:          uid,
		Thumbnail:    thumbnail,
		Name:         name,
		Profile:      profile,
		FacebookURL:  facebookURL,
		InstagramURL: instagramURL,
		TwitterURL:   twitterURL,
	}
}

func NewUserSummary(id int, uid, name, icon string) *UserSummary {
	return &UserSummary{
		ID:   id,
		UID:  uid,
		Name: name,
		Icon: icon,
	}
}
