package constant

import (
	"time"
)

// article status
const (
	ARTICLE_EDIT    = 0
	ARTICLE_OPEN    = 1
	ARTICLE_WAITING = 2
	ARTICLE_EXPIRED = 3
)

func GetArticleStatusId(status int, startTime string, endTime *string) int {
	_startTime, _ := time.Parse(time.RFC3339, startTime)
	_endTime := time.Now().Add(999)
	if endTime != nil {
		__endTime, _ := time.Parse(time.RFC3339, *endTime)
		_endTime = __endTime
	}
	now := time.Now()

	switch status {
	case ARTICLE_OPEN:
		// 掲載開始日が現在時刻より前なら公開待ち
		if _startTime.After(now) {
			return ARTICLE_WAITING
		}

		// 掲載終了日が現在時刻より前なら期限切れ
		if _endTime.Before(now) {
			return ARTICLE_EXPIRED
		}

		return ARTICLE_OPEN
	case ARTICLE_EDIT:
		return ARTICLE_EDIT
	default:
		return ARTICLE_EDIT
	}
}

func GetArticleStatusLabel(status int) string {
	switch status {
	case ARTICLE_OPEN:
		return "公開中"
	case ARTICLE_WAITING:
		return "公開待"
	case ARTICLE_EXPIRED:
		return "期限切"
	case ARTICLE_EDIT:
		return "下書き"
	default:
		return "下書き"
	}
}
