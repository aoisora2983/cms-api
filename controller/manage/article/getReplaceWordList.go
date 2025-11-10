package article

import (
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetReplaceWordList(c *gin.Context) {
	var req request.GetReplaceWordListRequest
	if !helper.BindQuery(c, &req) {
		return
	}

	wordList, err := models.GetCorrectWordListById(req.Id)
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.CreatedResponse(c, wordList)
}
