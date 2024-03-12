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

/**
 * 一般向け：公開中記事一覧を取得する
 */
func GetOpenArticleList(c *gin.Context) {
	var req request.GetOpenArticleListRequest
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

	param := models.BlogContentParam{
		IsOpen:  true,
		Keyword: req.Keyword,
		Tags:    req.Tags,
		Limit:   req.Limit,
		Offset:  req.Limit * req.Page,
	}
	articleList, err := models.GetBlogContentList(param)
	if err != nil {
		response.CustomErrorResponse(
			c,
			http.StatusInternalServerError,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}

	var index int16 = 0
	var total int64
	// 順番を維持しつつカテゴリ、タグを付与
	articleMap := make(map[int16]interface{})
	for _, article := range articleList {
		content := make(map[string]interface{})
		content["content"] = article
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
			response.CustomErrorResponse(
				c,
				http.StatusInternalServerError,
				map[string]string{code.SERVER_ERROR: err.Error()},
			)
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
				response.CustomErrorResponse(
					c,
					http.StatusInternalServerError,
					map[string]string{code.SERVER_ERROR: err.Error()},
				)
				return
			}
			content["tags"] = tags
		} else {
			content["tags"] = tags
		}

		metaList, err := models.GetBlogContentMetaList(int(article["id"].(int32)))
		if err != nil {
			response.CustomErrorResponse(
				c,
				http.StatusInternalServerError,
				map[string]string{code.SERVER_ERROR: err.Error()},
			)
			return
		}

		// "meta" => "meta_key" => val... でアクセスできるようにする
		_metaList := make(map[string]interface{})
		for _, meta := range metaList {
			_metaList[meta.MetaKey] = meta.MetaValue
		}
		content["meta"] = _metaList

		articleMap[index] = content
		index++
	}

	c.JSON(http.StatusCreated, gin.H{
		"articles": articleMap,
		"total":    total,
	})
}
