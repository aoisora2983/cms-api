package article

import (
	"cms/constant"
	"cms/db/models"
	"cms/package/correct"
	code "cms/package/error"
	"cms/package/request"
	"cms/package/response"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func RegisterArticle(c *gin.Context) {
	var req request.RegisterArticleRequest
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

	cUserId, exists := c.Get("userId")
	if !exists {
		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.NOT_EXISTS: "存在しないユーザーです。"},
		)
		return
	}
	userIdStr := cUserId.(string)
	userId, _ := strconv.Atoi(userIdStr)

	var id int = 1
	var idBranch int = 1
	isCreate := true
	// id, branch指定があれば、
	if req.Id != nil && req.IdBranch != nil {
		id = *req.Id
		idBranch = *req.IdBranch

		//前の記事情報を取得
		content, err := models.GetBlogContent(id, idBranch, false)
		if err != nil {
			response.CustomErrorResponse(
				c,
				http.StatusInternalServerError,
				map[string]string{code.SERVER_ERROR: err.Error()},
			)
		}

		// 枝番繰り上げ or 上書き判定
		// 一時保存 -> 一時保存, 公開 -> 公開 ただのupdate
		prevStatus := int(content["status"].(int16))
		nextStatus := int(req.Status)
		if prevStatus == req.Status {
			isCreate = false
		} else if prevStatus == constant.ARTICLE_EDIT && nextStatus == constant.ARTICLE_OPEN {
			// 一時保存 -> 公開 updateでいいが、前の枝番は全て削除する
			isCreate = false
			models.DeleteContentsBeforeBranch(id, idBranch)

		} else if prevStatus == constant.ARTICLE_OPEN && nextStatus == constant.ARTICLE_EDIT {
			// 公開 -> 一時保存 insertし、前の枝番は残しておく
			idBranch++
		}
	} else { // 無ければ新規登録
		// 最新のID取得
		latestId, err := models.GetLatestId()
		if err != nil {
			response.CustomErrorResponse(
				c,
				http.StatusInternalServerError,
				map[string]string{code.SERVER_ERROR: err.Error()},
			)
			return
		}

		id = latestId + 1
	}

	var publishedStartTime string = req.PublishedStartDate + req.PublishedStartTime
	var publishedEndTime *string = nil
	if req.PublishedEndDate != "" && req.PublishedEndTime != "" {
		_publishedEndTime := req.PublishedEndDate + req.PublishedEndTime
		publishedEndTime = &_publishedEndTime
	}

	// 本文の警告類を削除
	stringReader := strings.NewReader(string(req.Content))
	doc, _ := goquery.NewDocumentFromReader(stringReader)

	// tooltip類は前回チェックが残っているだけなので削除しておく
	correct.ResetTooltip(doc)
	html, _ := doc.Find("body").Html()

	// insert blog_content
	content := map[string]interface{}{
		"id":                   id,
		"id_branch":            idBranch,
		"id_user":              userId,
		"title":                req.Title,
		"content":              html,
		"status":               req.Status,
		"thumbnail":            req.Thumbnail,
		"published_start_time": publishedStartTime,
		"published_end_time":   publishedEndTime,
		"description":          req.Description,
	}

	var err error
	if isCreate {
		// 新規登録
		err = models.InsertContent(content)
	} else {
		//更新
		err = models.UpdateContent(content)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":         http.StatusInternalServerError,
			"errorMessage": err.Error(),
		})
		return
	}

	// insert tag
	if !isCreate {
		// 過去のデータをすべて消しておく
		models.DeleteBlogTag(id, idBranch)
	}

	for _, tagId := range req.Tags {
		tag := map[string]interface{}{
			"id_content":        id,
			"id_branch_content": idBranch,
			"id_tag":            tagId,
		}

		models.InsertBlogTag(tag)

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
