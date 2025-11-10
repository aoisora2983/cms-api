package article

import (
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * 記事のいいねをカウントアップする
 */
func CountUpArticleGood(c *gin.Context) {
	var req request.CountUpArticleGoodRequest
	if !helper.BindQuery(c, &req) {
		return
	}

	err := models.CountUpArticleGood(req.Id)
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.OKResponse(c)
}
