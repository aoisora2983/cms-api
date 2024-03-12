package constant

// comment status
const (
	COMMENT_WAITING_APPROVAL = 0
	COMMENT_OPEN             = 1
)

func GetCommentStatusLabel(status int) string {
	switch status {
	case COMMENT_OPEN:
		return "公開中"
	default:
		return "承認待"
	}
}
