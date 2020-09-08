package firebase

type (
	CloudMessageData struct {
		Token string // 送信先デバイストークン
		Body  string
		Data  map[string]string
		Badge int // アプリバッジ数
	}
)

func NewCloudMessageData(token, body string, data map[string]string, badge int) *CloudMessageData {
	return &CloudMessageData{
		Token: token,
		Body:  body,
		Data:  data,
		Badge: badge,
	}
}

func (c *CloudMessageData) AddData(key, val string) {
	c.Data[key] = val
}
