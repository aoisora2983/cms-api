package article

import (
	"cms/db/models"
	code "cms/package/error"
	"cms/package/request"
	"cms/package/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetReplaceWordList(c *gin.Context) {
	var req request.GetReplaceWordListRequest
	wordList, err := models.GetCorrectWordListById(req.Id)
	if err != nil {
		response.CustomErrorResponse(
			c,
			http.StatusInternalServerError,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}

	c.JSON(http.StatusCreated, wordList)
}
