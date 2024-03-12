package correct

import (
	"github.com/PuerkitoBio/goquery"
)

// 校正用のtooltip削除
func ResetTooltip(doc *goquery.Document) {
	doc.Find("span.correct-tooltip").Each(func(i int, s *goquery.Selection) {
		s.Contents().Unwrap()
	})

	// 画像の場合spanで囲えないので親要素にクラスと属性を付与しているため、クラス類だけ削除
	doc.Find(".correct-tooltip").Each(func(i int, s *goquery.Selection) {
		s.RemoveClass("correct-tooltip")
		s.RemoveClass("correct-tooltip-1")
		s.RemoveClass("correct-tooltip-2")
		s.RemoveAttr("data-tooltip")
	})
}
