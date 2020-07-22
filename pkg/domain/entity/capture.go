package entity

type (
	CaptureResult struct {
		CfProjectID    int `json:"cfProjectId"`
		SupporterCount int `json:"supporterCount"`
		AchievedPrice  int `json:"achievedPrice"`
	}
)
