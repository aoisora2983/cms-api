package article

import (
	"cms/db/models"
	code "cms/package/error"
	"cms/package/request"
	"cms/package/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterReplaceWord(c *gin.Context) {
	var req request.RegisterReplaceWordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
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

	// 置換文字列更新
	models.SaveCorrectWord(models.CorrectWord{
		Id:              req.Id,
		IdAccessibility: req.IdAccessibility,
		WordFrom:        req.WordFrom,
		WordTo:          req.WordTo,
		Level:           req.Level,
	})

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
