package article

import (
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterTag(c *gin.Context) {
	var req request.RegisterTagRequest
	if !helper.BindRequest(c, &req) {
		return
	}

	tag := map[string]interface{}{
		"id":       req.Id,
		"name":     req.Name,
		"filename": req.Filename,
	}

	err := models.SaveTag(tag)
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.OKResponse(c)
}
