package article

import (
	"cms/db/models"
	code "cms/package/error"
	"cms/package/request"
	"cms/package/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterTag(c *gin.Context) {
	var req request.RegisterTagRequest
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

	tag := map[string]interface{}{
		"id":       req.Id,
		"name":     req.Name,
		"filename": req.Filename,
	}

	// タグ追加
	models.SaveTag(tag)

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
