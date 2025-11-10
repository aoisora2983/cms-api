package article

import (
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteTag(c *gin.Context) {
	var req request.DeleteTagRequest
	if !helper.BindRequest(c, &req) {
		return
	}

	err := models.DeleteTag(req.Id)
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.OKResponse(c)
}
