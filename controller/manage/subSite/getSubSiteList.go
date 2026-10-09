package subSite

import (
	"cms/db/models"
	code "cms/package/error"
	"cms/package/helper"
	"cms/package/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSubSiteList(c *gin.Context) {
	subSites, err := models.GetSubSiteList(nil)
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
