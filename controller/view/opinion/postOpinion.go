package opinion

import (
	"cms/db/models"
	code "cms/package/error"
	"cms/package/request"
	"cms/package/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * 問い合わせ登録
 */
func PostOpinion(c *gin.Context) {
	var req request.PostOpinionRequest
	if err := c.Bind(&req); err != nil {
		if validErr, ok := err.(response.ValidationError); ok {
			c.JSON(validErr.GetStatus(), validErr.GetResponse())
			return
		}

		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
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
		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
