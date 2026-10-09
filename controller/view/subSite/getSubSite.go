package subSite

import (
	"cms/db/models"
	code "cms/package/error"
	"cms/package/helper"
	"cms/package/request"
	"cms/package/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSubSite(c *gin.Context) {
	var req request.GetSubSiteRequest
	if !helper.BindQuery(c, &req) {
		return
	}

	var subSite models.SubSite

	fmt.Println(req.SubDirName)

	if req.Id != nil {
		_subSite, err := models.GetSubSiteById(*req.Id)
		if err != nil {
			response.CustomErrorResponse(
				c,
				http.StatusInternalServerError,
				map[string]string{code.SERVER_ERROR: err.Error()},
			)
			return
		} else {
			subSite = _subSite
		}
	} else if req.SubDirName != nil {
		_subSite, err := models.GetSubSiteByDirname(*req.SubDirName)
		if err != nil {
			response.CustomErrorResponse(
				c,
				http.StatusInternalServerError,
				map[string]string{code.SERVER_ERROR: err.Error()},
			)
			return
		} else {
			subSite = _subSite
		}
	}

	helper.CreatedResponse(c, subSite)
}
