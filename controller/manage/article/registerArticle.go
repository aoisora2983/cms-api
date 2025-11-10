package article

import (
	"cms/constant"
	"cms/db/models"
	"cms/package/correct"
	"cms/package/helper"
	"cms/package/request"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func RegisterArticle(c *gin.Context) {
	var req request.RegisterArticleRequest
	if !helper.BindRequest(c, &req) {
		return
	}

	userId, ok := helper.GetUserIDFromContext(c)
	if !ok {
		return
	}

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
			helper.HandleError(c, err, http.StatusInternalServerError)
			return
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
			helper.HandleError(c, err, http.StatusInternalServerError)
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
		helper.HandleError(c, err, http.StatusInternalServerError)
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

		err := models.InsertBlogTag(tag)
		if err != nil {
			helper.HandleError(c, err, http.StatusInternalServerError)
			return
		}
	}

	helper.OKResponse(c)
}
