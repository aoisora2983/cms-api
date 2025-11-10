package article

import (
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func DeleteArticle(c *gin.Context) {
	var req request.DeleteArticleRequest
	if !helper.BindRequest(c, &req) {
		return
	}

	for _, idBranch := range req.Targets {
		var result = strings.Split(idBranch, "-")
		var id, _ = strconv.Atoi(result[0])
		var idBranch, _ = strconv.Atoi(result[1])

		err := models.DeleteContent(id, idBranch)
		if err != nil {
			helper.HandleError(c, err, http.StatusInternalServerError)
			return
		}
	}

	helper.OKResponse(c)
}
