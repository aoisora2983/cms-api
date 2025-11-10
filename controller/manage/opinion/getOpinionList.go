package opinion

import (
	"cms/db/models"
	"cms/package/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOpinionList(c *gin.Context) {
	opinions, err := models.GetOpinionList()
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.CreatedResponse(c, opinions)
}
