package article

import (
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteComment(c *gin.Context) {
	var req request.DeleteCommentRequest
	if !helper.BindRequest(c, &req) {
		return
	}

	// TODO：紐付くコメントがある場合、それも消えるので警告を出してから消す

	err := models.DeleteComment(req.Ids)
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.OKResponse(c)
}
