package article

import (
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * 特殊記事取得
 */
func GetUniqueArticle(c *gin.Context) {
	var req request.GetUniqueArticleRequest
	if !helper.BindQuery(c, &req) {
		return
	}

	content, err := models.GetBlogContentByPageType(req.PageType, true)
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.CreatedResponse(c, content)
}
