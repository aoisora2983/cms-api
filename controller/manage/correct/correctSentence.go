package correct

import (
	"cms/constant"
	"cms/db/models"
	pgCorrect "cms/package/correct"
	code "cms/package/error"
	"cms/package/request"
	"cms/package/response"
	"net/http"
	"slices"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

// 文章の校正
func CorrectSentence(c *gin.Context) {
	var req request.CorrectRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		if validErr, ok := err.(response.ValidationError); ok {
			c.JSON(validErr.GetStatus(), validErr.GetResponse())
			return
		}

		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}

	skipIds := req.SkipIds
	sentence := req.Sentence
	stringReader := strings.NewReader(string(sentence))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}
	// tooltip類は前回チェックが残っているだけなので削除しておく
	pgCorrect.ResetTooltip(doc)
	html, _ := doc.Find("body").Html()
	sentence = html

	accessibilityList, err := models.GetAccessibilityList()
	if err != nil {
		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
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

			c.JSON(http.StatusCreated, status.Response())
			return
		}
	}

	c.JSON(http.StatusCreated, status.Response())
}
