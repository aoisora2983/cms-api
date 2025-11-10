package article

import (
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterReplaceWord(c *gin.Context) {
	var req request.RegisterReplaceWordRequest
	if !helper.BindRequest(c, &req) {
		return
	}

	// 置換文字列更新
	err := models.SaveCorrectWord(models.CorrectWord{
		Id:              req.Id,
		IdAccessibility: req.IdAccessibility,
		WordFrom:        req.WordFrom,
		WordTo:          req.WordTo,
		Level:           req.Level,
	})
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.OKResponse(c)
}
