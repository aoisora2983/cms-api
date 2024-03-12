package portfolio

import (
	"cms/db/models"
	code "cms/package/error"
	"cms/package/request"
	"cms/package/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterPortfolio(c *gin.Context) {
	var req request.RegisterPortfolioRequest
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

	// タグ追加
	models.SavePortfolio(portfolio)

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
