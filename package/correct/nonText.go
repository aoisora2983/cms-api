package correct

import (
	"cms/constant"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// 1.1.1 非テキストコンテンツ
// https://waic.jp/translations/UNDERSTANDING-WCAG20/text-equiv-all.html
func CorrectNonText(doc *goquery.Document, status *CorrectStatus, skipIds []int, errorMessage string) {
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		alt, hasAttribute := s.Attr("alt")
		p := s.Parent()

		if !hasAttribute {
			// 属性が無いのはNG
			status.SetLevel(constant.CORRECT_NG)
			status.SetMessage(errorMessage)
			p.AddClass(fmt.Sprintf("correct-tooltip correct-tooltip-%d", constant.CORRECT_NG))
			p.SetAttr("data-tooltip", errorMessage)
		} else {
			// 属性があるうえで文字列が空なら警告だけ出す
			if alt == "" {
				// 既にNGがあるなら表示はエラーを優先
				if status.GetLevel() != constant.CORRECT_NG {
					status.SetLevel(constant.CORRECT_WARNING)
					status.SetMessage(errorMessage)
				}
				p.AddClass(fmt.Sprintf("correct-tooltip correct-tooltip-%d", constant.CORRECT_WARNING))
				p.SetAttr("data-tooltip", errorMessage)
			}
		}
	})

	html, _ := doc.Find("body").Html()
	status.SetSentence(html)
}
