package article

import (
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOpenTagList(c *gin.Context) {
	var req request.GetOpenTagListRequest
	if !helper.BindQuery(c, &req) {
		return
	}

	tagList, err := models.GetOpenTagList(req.IdSubSite)
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.CreatedResponse(c, tagList)
}
