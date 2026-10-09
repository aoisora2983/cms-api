package subSite

import (
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterSubSite(c *gin.Context) {
	var req request.RegisterSubSiteRequest
	if !helper.BindRequest(c, &req) {
		return
	}

	err := models.SaveSubSite(req)
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.OKResponse(c)
}
