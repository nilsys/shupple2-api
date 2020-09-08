package model

func (x NoticeActionTargetType) Label() string {
	switch x {
	case NoticeActionTargetTypePOST:
		return "記事"
	case NoticeActionTargetTypeVLOG:
		return "Vlog"
	case NoticeActionTargetTypeREVIEW:
		return "レビュー"
	case NoticeActionTargetTypeCOMMENT:
		return "コメント"
	case NoticeActionTargetTypeREPLY:
		return "リプライ"
	case NoticeActionTargetTypeUSER:
		// USERの時はFOLLOWしか存在しないので実質使われない
		return "ユーザー"
	case NoticeActionTargetTypeCOMIC:
		return "漫画"
	default:
		return ""
	}
}
