package opinion

import (
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * 問い合わせ登録
 */
func PostOpinion(c *gin.Context) {
	var req request.PostOpinionRequest
	if !helper.BindQuery(c, &req) {
		return
	}

	model := models.Opinion{
		Name:     req.Name,
		Email:    req.Email,
		Content:  req.Content,
		IP:       c.ClientIP(),
		SendTime: time.Now().Format(time.RFC3339),
	}

	err := models.SaveOpinion(model)
	if err != nil {
		helper.HandleError(c, err, http.StatusBadRequest)
		return
	}

	helper.OKResponse(c)
}
