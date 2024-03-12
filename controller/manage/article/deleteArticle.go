package article

import (
	"cms/db/models"
	code "cms/package/error"
	"cms/package/request"
	"cms/package/response"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func DeleteArticle(c *gin.Context) {
	var req request.DeleteArticleRequest
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

	for _, idBranch := range req.Targets {
		var result = strings.Split(idBranch, "-")
		var id, _ = strconv.Atoi(result[0])
		var idBranch, _ = strconv.Atoi(result[1])

		err := models.DeleteContent(id, idBranch)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":         http.StatusInternalServerError,
				"errorMessage": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
