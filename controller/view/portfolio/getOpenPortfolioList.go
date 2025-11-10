package portfolio

import (
	"cms/db/models"
	"cms/package/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOpenPortfolioList(c *gin.Context) {
	portfolioList, err := models.GetOpenPortfolioList()
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.CreatedResponse(c, portfolioList)
}
