package correct

import (
	"cms/constant"
	"fmt"
	"slices"

	"github.com/PuerkitoBio/goquery"
)

// 1.1.1 非テキストコンテンツ
// https://waic.jp/translations/UNDERSTANDING-WCAG20/text-equiv-all.html
func CorrectNonText(doc *goquery.Document, status *CorrectStatus, skipIds []int) {
	status.SetKind(constant.NON_TEXT)

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		alt, hasAttribute := s.Attr("alt")
		p := s.Parent()

		if !hasAttribute {
			// 属性が無いのはNG
			status.SetLevel(constant.CORRECT_NG)
			status.SetMessage(constant.NO_ALT_ATTRIBUTE)
			p.AddClass(fmt.Sprintf("correct-tooltip correct-tooltip-%d", constant.CORRECT_NG))
			p.SetAttr("data-tooltip", constant.NO_ALT_ATTRIBUTE)
		} else {
			// 属性があるうえで文字列が空なら警告だけ出す(警告OKならスキップ)
			if !slices.Contains(skipIds, constant.NON_TEXT) && alt == "" {
				// 既にNGがあるなら表示はエラーを優先
				if status.GetLevel() != constant.CORRECT_NG {
					status.SetLevel(constant.CORRECT_WARNING)
					status.SetMessage(constant.NO_ALT_TEXT)
				}
				p.AddClass(fmt.Sprintf("correct-tooltip correct-tooltip-%d", constant.CORRECT_WARNING))
				p.SetAttr("data-tooltip", constant.NO_ALT_TEXT)
			}
		}
	})

	html, _ := doc.Find("body").Html()
	status.SetSentence(html)
}
