package portfolio

import (
	"cms/db/models"
	code "cms/package/error"
	"cms/package/request"
	"cms/package/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeletePortfolio(c *gin.Context) {
	var req request.DeletePortfolioRequest
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

	models.DeletePortfolio(req.Id)

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
