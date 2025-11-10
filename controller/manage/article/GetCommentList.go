package article

import (
	"cms/db/models"
	"cms/package/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCommentList(c *gin.Context) {
	commentList, err := models.GetCommentList()
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.CreatedResponse(c, commentList)
}
