package article

import (
	"cms/db/models"
	"cms/package/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAccessibilityList(c *gin.Context) {
	accessibilityList, err := models.GetAccessibilityList()
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.CreatedResponse(c, accessibilityList)
}
