package comment

import (
	"cms/constant"
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * コメント登録
 */
func PostComment(c *gin.Context) {
	var req request.PostCommentRequest
	if !helper.BindQuery(c, &req) {
		return
	}

	model := models.BlogComment{
		IdBlogContent: req.ArticleId,
		IdReplay:      req.ReplayId,
		UserName:      req.UserName,
		Comment:       req.Comment,
		Ip:            c.ClientIP(),
		Status:        constant.COMMENT_WAITING_APPROVAL,
		CommentTime:   time.Now().Format(time.RFC3339),
	}

	err := models.InsertComment(model)
	if err != nil {
		helper.HandleError(c, err, http.StatusBadRequest)
		return
	}

	helper.OKResponse(c)
}
