package subSite

import (
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteSubSite(c *gin.Context) {
	var req request.DeleteSubSiteRequest
	if !helper.BindRequest(c, &req) {
		return
	}

	err := models.DeleteSubSite(req.Id)
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.OKResponse(c)
}
