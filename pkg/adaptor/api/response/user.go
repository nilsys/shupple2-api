package response

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

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
	}

	// ユーザーランキングで返すレスポンス型
	RankinUser struct {
		ID        int      `jso:"id"`
		UID       string   `json:"uid"`
		Name      string   `json:"name"`
		Profile   string   `json:"profile"`
		Thumbnail string   `json:"iconUrl"`
		Interest  []string `json:"interest"`
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
)

func NewCreatorFromUser(user *entity.User) Creator {
	return NewCreator(user.ID, user.UID, user.GenerateThumbnailURL(), user.Name, user.Profile, user.FacebookURL, user.InstagramURL, user.TwitterURL)
}

func NewCreator(id int, uid, thumbnail, name, profile, facebookURL, instagramURL, twitterURL string) Creator {
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
