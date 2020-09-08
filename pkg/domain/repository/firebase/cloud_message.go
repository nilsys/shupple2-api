package firebase

type (
	CloudMessageCommandRepository interface {
		Send(token, body string, data map[string]string, badge int) error
	}
)
