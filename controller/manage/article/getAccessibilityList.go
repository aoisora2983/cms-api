package article

import (
	"cms/db/models"
	code "cms/package/error"
	"cms/package/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAccessibilityList(c *gin.Context) {
	accessibilityList, err := models.GetAccessibilityList()
	if err != nil {
		response.CustomErrorResponse(
			c,
			http.StatusInternalServerError,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}

	c.JSON(http.StatusCreated, accessibilityList)
}
