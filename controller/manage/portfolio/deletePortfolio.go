package portfolio

import (
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeletePortfolio(c *gin.Context) {
	var req request.DeletePortfolioRequest
	if !helper.BindRequest(c, &req) {
		return
	}

	err := models.DeletePortfolio(req.Id)
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.OKResponse(c)
}
