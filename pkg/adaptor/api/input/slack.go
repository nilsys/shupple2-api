package input

import (
	"strconv"
	"strings"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	// TODO: slackのレスポンスが受け取りづらい
	SlackCallbackPayload struct {
		Payload string `json:"payload"`
	}

	SlackCallback struct {
		Type        string         `json:"type"`
		Team        SlackTeam      `json:"team"`
		User        SlackUser      `json:"user"`
		APIAppID    string         `json:"api_app_id"`
		Token       string         `json:"token"`
		Container   SlackContainer `json:"container"`
		TriggerID   string         `json:"trigger_id"`
		Channel     SlackChannel   `json:"channel"`
		Message     SlackMessage   `json:"message"`
		ResponseURL string         `json:"response_url"`
		Actions     []SlackAction  `json:"actions"`
	}

	SlackTeam struct {
		ID     string `json:"id"`
		Domain string `json:"domain"`
	}

	SlackUser struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Name     string `json:"name"`
		TeamID   string `json:"team_id"`
	}

	SlackContainer struct {
		Type        string `json:"type"`
		MessageTs   string `json:"message_ts"`
		ChannelID   string `json:"channel_id"`
		IsEphemeral bool   `json:"is_ephemeral"`
	}

	SlackChannel struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	SlackMessage struct {
		Type    string       `json:"type"`
		Subtype string       `json:"subtype"`
		Text    string       `json:"text"`
		Ts      string       `json:"ts"`
		BotID   string       `json:"bot_id"`
		Blocks  []SlackBlock `json:"blocks"`
	}

	SlackBlock struct {
		Type     string         `json:"type"`
		BlockID  string         `json:"block_id"`
		Text     SlackBlockText `json:"text,omitempty"`
		Fields   []SlackField   `json:"fields,omitempty"`
		Elements []SlackElement `json:"elements,omitempty"`
	}

	SlackBlockText struct {
		Type     string `json:"type"`
		Text     string `json:"text"`
		Verbatim bool   `json:"verbatim"`
	}

	SlackField struct {
		Type     string `json:"type"`
		Text     string `json:"text"`
		Verbatim bool   `json:"verbatim"`
	}

	SlackElement struct {
		Type     string           `json:"type"`
		ActionID string           `json:"action_id"`
		Text     SlackElementText `json:"text"`
		Style    string           `json:"style"`
		Value    string           `json:"value"`
	}

	SlackElementText struct {
		Type  string `json:"type"`
		Text  string `json:"text"`
		Emoji bool   `json:"emoji"`
	}

	SlackAction struct {
		ActionID string          `json:"action_id"`
		BlockID  string          `json:"block_id"`
		Text     SlackActionText `json:"text"`
		Value    string          `json:"value"`
		Style    string          `json:"style"`
		Type     string          `json:"type"`
		ActionTs string          `json:"action_ts"`
	}

	SlackActionText struct {
		Type  string `json:"type"`
		Text  string `json:"text"`
		Emoji bool   `json:"emoji"`
	}
)

// TODO: マジックナンバーかつ、起こらないとは思うが、rangeでコケる可能性あり
// 運用上は問題ないがプログラムとしては×
func (s *SlackCallback) ReportUserID() int {
	val := strings.Split(s.Actions[0].Value, "-")
	userID, _ := strconv.Atoi(val[2])
	return userID
}

func (s *SlackCallback) TargetID() int {
	val := strings.Split(s.Actions[0].Value, "-")
	targetID, _ := strconv.Atoi(val[1])
	return targetID
}

func (s *SlackCallback) TargetType() model.ReportTargetType {
	val := strings.Split(s.Actions[0].Value, "-")
	targetType, _ := model.ParseReportTargetType(val[0])
	return targetType
}

func (s *SlackCallback) IsApproved() bool {
	return s.Actions[0].Text.Text == "Approve"
}
