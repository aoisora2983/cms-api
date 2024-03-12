package article

import (
	"cms/constant"
	"cms/db/models"
	code "cms/package/error"
	"cms/package/request"
	"cms/package/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArticle(c *gin.Context) {
	var req request.GetArticleRequest
	if err := c.Bind(&req); err != nil {
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

	content, err := models.GetBlogContent(req.Id, req.IdBranch, false)
	if err != nil {
		response.CustomErrorResponse(
			c,
			http.StatusInternalServerError,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
	}
	// statusを読める形に変換
	status := int(content["status"].(int16))
	statusLabel := constant.GetArticleStatusLabel(status)
	content["status"] = map[string]interface{}{
		"id":    status,
		"label": statusLabel,
	}

	articleMap := make(map[string]interface{})

	// タグ取得
	blogTags, err := models.GetBlogTags(content["id"].(int32), content["id_branch"].(int32))
	if err != nil {
		response.CustomErrorResponse(
			c,
			http.StatusInternalServerError,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}
	var tags []map[string]interface{}
	for _, blogTag := range blogTags {
		tags = append(tags, map[string]interface{}{
			"id":        blogTag["id_tag"].(int32),
			"icon_path": blogTag["icon_path"],
			"label":     blogTag["name"],
		})
	}

	articleMap["content"] = content
	articleMap["tags"] = tags

	c.JSON(http.StatusCreated, articleMap)
}
