package response

type (
	Creator struct {
		ID        int    `json:"id"`
		Thumbnail string `json:"iconUrl"`
		Name      string `json:"name"`
		Profile   string `json:"profile,omitempty"`
	}

	// ユーザーランキングで返すレスポンス型
	RankinUser struct {
		ID        int      `jso:"id"`
		Name      string   `json:"name"`
		Profile   string   `json:"profile"`
		Thumbnail string   `json:"thumbnail"`
		Interest  []string `json:"interest"`
	}

	// MEMO: ユースケースが増えれば命名返る
	FollowUser struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Thumbnail string `json:"thumbnail"`
	}

	UserSummary struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Icon string `json:"iconUrl"`
	}
)

func NewCreator(id int, thumbnail, name, profile string) Creator {
	return Creator{
		ID:        id,
		Thumbnail: thumbnail,
		Name:      name,
		Profile:   profile,
	}
}

func NewUserSummary(id int, name string, Icon string) *UserSummary {
	return &UserSummary{
		ID:   id,
		Name: name,
		Icon: Icon,
	}
}
