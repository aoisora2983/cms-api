package subSite

import (
	"cms/constant"
	"cms/db/models"
	code "cms/package/error"
	"cms/package/helper"
	"cms/package/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOpenSubSiteList(c *gin.Context) {
	status := constant.SUB_SITE_OPEN
	subSites, err := models.GetSubSiteList(&status)
	if err != nil {
		response.CustomErrorResponse(
			c,
			http.StatusInternalServerError,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}

	helper.CreatedResponse(c, subSites)
}
