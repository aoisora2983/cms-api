package article

import (
	"cms/constant"
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArticle(c *gin.Context) {
	var req request.GetArticleRequest
	if !helper.BindQuery(c, &req) {
		return
	}

	content, err := models.GetBlogContent(req.Id, req.IdBranch, false)
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
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
		helper.HandleError(c, err, http.StatusInternalServerError)
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
