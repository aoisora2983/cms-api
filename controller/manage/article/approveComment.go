package article

import (
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApproveComment(c *gin.Context) {
	var req request.ApproveCommentRequest
	if !helper.BindRequest(c, &req) {
		return
	}

	err := models.ApproveComment(req.Ids)
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.OKResponse(c)
}
