package comment

import (
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * コメントのいいねをカウントアップする
 */
func CountUpCommentGood(c *gin.Context) {
	var req request.CountUpCommentGoodRequest
	if !helper.BindQuery(c, &req) {
		return
	}

	err := models.CountUpComment(req.Id)
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.OKResponse(c)
}
