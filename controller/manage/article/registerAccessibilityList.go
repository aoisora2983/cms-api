package article

import (
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAccessibilityList(c *gin.Context) {
	var req request.RegisterAccessibilityListRequest
	if !helper.BindRequest(c, &req) {
		return
	}

	// アクセシビリティ設定更新
	for _, accessibility := range req.AccessibilityList {
		err := models.SaveAccessibility(models.Accessibility{
			Id:      accessibility.Id,
			Message: accessibility.Message,
			Level:   accessibility.Level,
		})
		if err != nil {
			helper.HandleError(c, err, http.StatusInternalServerError)
			return
		}
	}

	helper.OKResponse(c)
}
