package portfolio

import (
	"cms/db/models"
	"cms/package/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPortfolioList(c *gin.Context) {
	portFolioList, err := models.GetPortfolioList()
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.CreatedResponse(c, portFolioList)
}
