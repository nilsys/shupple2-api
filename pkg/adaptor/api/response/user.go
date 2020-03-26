package response

type (
	Creator struct {
		ID        int    `json:"id"`
		UID       string `json:"uid"`
		Thumbnail string `json:"iconUrl"`
		Name      string `json:"name"`
		Profile   string `json:"profile,omitempty"`
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

func NewCreator(id int, uid, thumbnail, name, profile string) Creator {
	return Creator{
		ID:        id,
		UID:       uid,
		Thumbnail: thumbnail,
		Name:      name,
		Profile:   profile,
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
