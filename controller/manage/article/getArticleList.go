package article

import (
	"cms/constant"
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArticleList(c *gin.Context) {
	var req request.GetArticleListRequest
	if !helper.BindQuery(c, &req) {
		return
	}

	param := models.BlogContentParam{
		Keyword: req.Keyword,
		Tags:    req.Tags,
	}

	articleList, err := models.GetBlogContentList(param)
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	var index int16 = 0
	var total int64
	// id => id_branch => contentの形に整形しつつカテゴリ、タグを付与
	articleMap := make(map[int32]interface{})
	// 順番維持して index => contentの形に最終的にまとめる
	sortMap := make(map[int16]interface{})
	for _, article := range articleList {
		id := article["id"].(int32)
		idBranch := article["id_branch"].(int32)
		content := make(map[string]interface{})
		idBranchMap := make(map[int32]interface{})
		if articleMap[id] != nil {
			idBranchMap = articleMap[id].(map[int32]interface{})
		}
		content["content"] = article
		idBranchMap[idBranch] = article
		total = article["total"].(int64)

		// statusを読める形に変換
		status := int(article["status"].(int16))
		statusLabel := constant.GetArticleStatusLabel(status)
		article["status"] = map[string]interface{}{
			"id":    status,
			"label": statusLabel,
		}

		// タグ取得
		blogTags, err := models.GetBlogTags(article["id"].(int32), article["id_branch"].(int32))
		if err != nil {
			helper.HandleError(c, err, http.StatusInternalServerError)
			return
		}

		tags := make([]models.Tag, 0)
		if blogTags != nil {
			var searchTags []int32
			for _, blogTag := range blogTags {
				searchTags = append(searchTags, blogTag["id_tag"].(int32))
			}

			tags, err := models.GetTagListByIds(searchTags)
			if err != nil {
				helper.HandleError(c, err, http.StatusInternalServerError)
				return
			}
			content["tags"] = tags
		} else {
			content["tags"] = tags
		}

		idBranchMap[idBranch] = content
		articleMap[id] = idBranchMap
		sortMap[index] = articleMap[id]
		index++
	}

	helper.CreatedResponse(c, gin.H{
		"articles": sortMap,
		"total":    total,
	})
}
