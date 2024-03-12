package constant

// portfolio status
const (
	PORTFOLIO_EDIT = 0
	PORTFOLIO_OPEN = 1
)

func GetPortfolioStatusLabel(status int) string {
	switch status {
	case ARTICLE_OPEN:
		return "公開中"
	case ARTICLE_EDIT:
		return "下書き"
	default:
		return "下書き"
	}
}
