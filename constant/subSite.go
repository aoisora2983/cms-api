package constant

// subsite status
const (
	SUB_SITE_EDIT = 0
	SUB_SITE_OPEN = 1
)

func GetSubSiteStatusLabel(status int) string {
	switch status {
	case SUB_SITE_OPEN:
		return "公開中"
	case SUB_SITE_EDIT:
		return "非公開"
	default:
		return "非公開"
	}
}
