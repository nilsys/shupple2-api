package response

type (
	Creator struct {
		Thumbnail string `json:"thumbnail"`
		Name      string `json:"name"`
		Profile   string `json:"profile,omitempty"`
	}
)

func NewCreator(thumbnail string, name string, profile string) Creator {
	return Creator{
		Thumbnail: thumbnail,
		Name:      name,
		Profile:   profile,
	}
}
