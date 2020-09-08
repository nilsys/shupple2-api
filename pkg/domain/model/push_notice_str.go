package model

import "fmt"

// Push通知body
func PushNoticeBody(userName string, target NoticeActionTargetType, action NoticeActionType) string {
	// フォロー・タグ付の場合は文言がactionだけで決定する
	switch action {
	case NoticeActionTypeFOLLOW:
		return fmt.Sprintf("%sさんがあなたをフォローしました", userName)
	case NoticeActionTypeTAGGED:
		return fmt.Sprintf("%sさんがあなたをタグ付けしました", userName)
	case NoticeActionTypeFAVORITE:
		return fmt.Sprintf("%sさんがあなたの%sにいいねしました", userName, target.Label())
	case NoticeActionTypeCOMMENT:
		return fmt.Sprintf("%sさんがあなたの%sにコメントしました", userName, target.Label())
	case NoticeActionTypeREPLY:
		return fmt.Sprintf("%sさんがあなたの%sにリプライしました", userName, target.Label())
	default:
		return ""
	}
}
