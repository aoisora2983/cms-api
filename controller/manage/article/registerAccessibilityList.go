package article

import (
	"cms/db/models"
	code "cms/package/error"
	"cms/package/request"
	"cms/package/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAccessibilityList(c *gin.Context) {
	var req request.RegisterAccessibilityListRequest
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

	// アクセシビリティ設定更新
	for _, accessibility := range req.AccessibilityList {
		models.SaveAccessibility(models.Accessibility{
			Id:      accessibility.Id,
			Message: accessibility.Message,
			Level:   accessibility.Level,
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
