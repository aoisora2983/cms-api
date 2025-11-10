package article

import (
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTag(c *gin.Context) {
	var req request.GetTagRequest
	if !helper.BindQuery(c, &req) {
		return
	}

	tag, err := models.GetTag(req.Id)
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.CreatedResponse(c, tag)
}
