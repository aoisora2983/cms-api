package comment

import (
	"cms/constant"
	"cms/db/models"
	code "cms/package/error"
	"cms/package/request"
	"cms/package/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * コメント登録
 */
func PostComment(c *gin.Context) {
	var req request.PostCommentRequest
	if err := c.Bind(&req); err != nil {
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
		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
