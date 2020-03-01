package response

type (
	Creator struct {
		Thumbnail string `json:"thumbnail"`
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
)

func NewCreator(thumbnail string, name string, profile string) Creator {
	return Creator{
		Thumbnail: thumbnail,
		Name:      name,
		Profile:   profile,
	}
}
