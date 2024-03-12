package article

import (
	"cms/db/models"
	code "cms/package/error"
	"cms/package/request"
	"cms/package/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteComment(c *gin.Context) {
	var req request.DeleteCommentRequest
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

	// TODO：紐付くコメントがある場合、それも消えるので警告を出してから消す

	// タグ追加
	models.DeleteComment(req.Ids)

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
