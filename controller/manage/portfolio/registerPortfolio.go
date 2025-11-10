package portfolio

import (
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterPortfolio(c *gin.Context) {
	var req request.RegisterPortfolioRequest
	if !helper.BindRequest(c, &req) {
		return
	}

	portfolio := map[string]interface{}{
		"id":           req.Id,
		"title":        req.Title,
		"description":  req.Description,
		"thumbnail":    req.Thumbnail,
		"detail_url":   req.DetailUrl,
		"release_time": req.ReleaseTime,
		"status":       req.Status,
		"sort_order":   req.SortOrder,
	}

	err := models.SavePortfolio(portfolio)
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.OKResponse(c)
}
