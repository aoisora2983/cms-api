package correct

import (
	"cms/constant"
	pgCorrect "cms/package/correct"
	code "cms/package/error"
	"cms/package/request"
	"cms/package/response"
	"net/http"
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

	status := pgCorrect.NewCorrectStatus()

	// 1.1.1 非テキストコンテンツ
	pgCorrect.CorrectNonText(doc, status, skipIds)
	if status.GetLevel() != constant.CORRECT_OK {
		c.JSON(http.StatusCreated, status.Response())
		return
	}

	// 文字列置換(not WCAG)
	pgCorrect.ReplaceWord(sentence, status, skipIds)
	if status.GetLevel() != constant.CORRECT_OK {
		c.JSON(http.StatusCreated, status.Response())
		return
	}

	c.JSON(http.StatusCreated, status.Response())
}
