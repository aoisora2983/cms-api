package correct

import (
	"cms/constant"
	"cms/db/models"
	pgCorrect "cms/package/correct"
	"cms/package/helper"
	"cms/package/request"
	"net/http"
	"slices"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

// 文章の校正
func CorrectSentence(c *gin.Context) {
	var req request.CorrectRequest

	if !helper.BindRequest(c, &req) {
		return
	}

	skipIds := req.SkipIds
	sentence := req.Sentence
	stringReader := strings.NewReader(string(sentence))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		helper.HandleError(c, err, http.StatusBadRequest)
		return
	}
	// tooltip類は前回チェックが残っているだけなので削除しておく
	pgCorrect.ResetTooltip(doc)
	html, _ := doc.Find("body").Html()
	sentence = html

	accessibilityList, err := models.GetAccessibilityList()
	if err != nil {
		helper.HandleError(c, err, http.StatusBadRequest)
		return
	}

	status := pgCorrect.NewCorrectStatus()

	for _, accessibility := range accessibilityList {
		status.SetKind(accessibility.Id)

		switch accessibility.Id {
		case constant.NON_TEXT:
			// 1.1.1 非テキストコンテンツ
			pgCorrect.CorrectNonText(doc, status, skipIds, accessibility.Message)
		case constant.REPLACE_WORD:
			// 文字列置換(not WCAG)
			pgCorrect.ReplaceWord(sentence, status, skipIds, accessibility)
		}

		if status.GetLevel() != constant.CORRECT_NO_CHECK {
			if accessibility.Level == constant.CORRECT_WARNING { // 警告設定ならNG項目でも警告に上書き
				status.SetLevel(constant.CORRECT_WARNING)
			}
			// WARNINGがあってもスキップIDならスキップ
			if status.GetLevel() == constant.CORRECT_WARNING && !slices.Contains(skipIds, accessibility.Id) {
				continue
			}

			helper.CreatedResponse(c, status.Response())
			return
		}
	}

	helper.CreatedResponse(c, status.Response())
}
