package slack

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	// request
	Report struct {
		HeadText HeadText
		Fields   Fields
		Elements Elements
	}

	HeadText struct {
		Type string `json:"type"`
		Text Text   `json:"text"`
	}
	Text struct {
		Type string `json:"type"`
		Text string `json:"text"`
	}

	Fields struct {
		Type   string  `json:"type"`
		Fields []Field `json:"fields"`
	}
	Field struct {
		Type string `json:"type"`
		Text string `json:"text"`
	}

	Elements struct {
		Type string    `json:"type"`
		List []Element `json:"elements"`
	}

	Element struct {
		Type  string      `json:"type"`
		Text  ElementText `json:"text"`
		Style string      `json:"style"`
		Value string      `json:"value"`
	}

	ElementText struct {
		Type  string `json:"type"`
		Emoji bool   `json:"emoji"`
		Text  string `json:"text"`
	}

	// TODO: 暫定対応
	Response struct {
		Ok      bool   `json:"ok"`
		Channel string `json:"channel"`
		Ts      string `json:"ts"`
		Message struct {
			BotID      string `json:"bot_id"`
			Type       string `json:"type"`
			Text       string `json:"text"`
			User       string `json:"user"`
			Ts         string `json:"ts"`
			Team       string `json:"team"`
			BotProfile struct {
				ID      string `json:"id"`
				Deleted bool   `json:"deleted"`
				Name    string `json:"name"`
				Updated int    `json:"updated"`
				AppID   string `json:"app_id"`
				Icons   struct {
					Image36 string `json:"image_36"`
					Image48 string `json:"image_48"`
					Image72 string `json:"image_72"`
				} `json:"icons"`
				TeamID string `json:"team_id"`
			} `json:"bot_profile"`
			Blocks []struct {
				Type    string `json:"type"`
				BlockID string `json:"block_id"`
				Text    struct {
					Type     string `json:"type"`
					Text     string `json:"text"`
					Verbatim bool   `json:"verbatim"`
				} `json:"text,omitempty"`
				Fields []struct {
					Type     string `json:"type"`
					Text     string `json:"text"`
					Verbatim bool   `json:"verbatim"`
				} `json:"fields,omitempty"`
				Elements []struct {
					Type     string `json:"type"`
					ActionID string `json:"action_id"`
					Text     struct {
						Type  string `json:"type"`
						Text  string `json:"text"`
						Emoji bool   `json:"emoji"`
					} `json:"text"`
					Style string `json:"style"`
					Value string `json:"value"`
				} `json:"elements,omitempty"`
			} `json:"blocks"`
		} `json:"message"`
	}
)

func NewSlackReport(targetType model.ReportTargetType, targetURL string, targetID int, targetBody, body string, reason model.ReportReasonType, reportUserID, reportedUserID int) *Report {
	text := Text{
		Type: "mrkdwn",
		Text: fmt.Sprintf("通報がありました:\n*<%s|対応して下さい>*", targetURL),
	}
	fieldList := []Field{
		{
			Type: "mrkdwn",
			Text: fmt.Sprintf("*ID:*\n\t%d", targetID),
		},
		{
			Type: "mrkdwn",
			Text: fmt.Sprintf("*Type:\n\t%s", targetType),
		},
		{
			Type: "mrkdwn",
			Text: fmt.Sprintf("*Body:*\n\t%s", targetBody),
		},
		{
			Type: "mrkdwn",
			Text: fmt.Sprintf("*When:*\n\t%s", time.Now().Format(time.RFC3339)),
		},
		{
			Type: "mrkdwn",
			Text: fmt.Sprintf("*Reason:*\n\t%s:\n\t%s", reason, body),
		},
		{
			Type: "mrkdwn",
			Text: fmt.Sprintf("*ReportUser:*\n\t%d", reportUserID),
		},
		{
			Type: "mrkdwn",
			Text: fmt.Sprintf("*ReportedUser:*\n\t%d", reportedUserID),
		},
	}
	approveEle := Element{
		Type: "button",
		Text: ElementText{
			Type:  "plain_text",
			Emoji: true,
			Text:  "Approve",
		},
		Style: "primary",
		Value: fmt.Sprintf("%s-%d-%d", targetType, targetID, reportUserID),
	}
	denyEle := Element{
		Type: "button",
		Text: ElementText{
			Type:  "plain_text",
			Emoji: true,
			Text:  "Deny",
		},
		Style: "danger",
		Value: fmt.Sprintf("%s-%d-%d", targetType, targetID, reportUserID),
	}

	headText := HeadText{
		Type: "section",
		Text: text,
	}
	fields := Fields{
		Type:   "section",
		Fields: fieldList,
	}
	elements := Elements{
		Type: "actions",
		List: []Element{approveEle, denyEle},
	}

	return &Report{
		HeadText: headText,
		Fields:   fields,
		Elements: elements,
	}
}

func (s *Report) ToSlackFmt() string {
	headTextJSON, _ := json.Marshal(s.HeadText)
	fieldsJSON, _ := json.Marshal(s.Fields)
	elementsJSON, _ := json.Marshal(s.Elements)
	logger.Info("[" + string(headTextJSON) + "," + string(fieldsJSON) + "," + string(elementsJSON) + "]")
	return "[" + string(headTextJSON) + "," + string(fieldsJSON) + "," + string(elementsJSON) + "]"
}
