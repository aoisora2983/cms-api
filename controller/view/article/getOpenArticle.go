package article

import (
	"cms/constant"
	"cms/db/models"
	code "cms/package/error"
	"cms/package/request"
	"cms/package/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * 一般向け：公開中記事を取得する
 */
func GetOpenArticle(c *gin.Context) {
	var req request.GetOpenArticleRequest
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

	content, err := models.GetBlogContent(req.Id, -1, true)
	if err != nil {
		response.CustomErrorResponse(
			c,
			http.StatusInternalServerError,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}

	articleMap := make(map[string]interface{})

	// タグ取得
	tags, err := getTagList(content)
	if err != nil {
		response.CustomErrorResponse(
			c,
			http.StatusInternalServerError,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}
	articleMap["tags"] = tags

	// metaタグ取得
	metaList, err := getMetaList(req.Id)
	if err != nil {
		response.CustomErrorResponse(
			c,
			http.StatusInternalServerError,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}
	articleMap["meta"] = metaList

	// コメント取得
	commentList, err := getCommentList(req.Id)
	if err != nil {
		response.CustomErrorResponse(
			c,
			http.StatusInternalServerError,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}
	articleMap["comments"] = commentList

	articleMap["content"] = content

	c.JSON(http.StatusCreated, articleMap)
}

func getTagList(content map[string]interface{}) ([]models.Tag, error) {
	blogTags, err := models.GetBlogTags(content["id"].(int32), content["id_branch"].(int32))
	if err != nil {
		return nil, err
	}

	tags := make([]models.Tag, 0)
	if blogTags != nil {
		var searchTags []int32
		for _, blogTag := range blogTags {
			searchTags = append(searchTags, blogTag["id_tag"].(int32))
		}

		tags, err := models.GetTagListByIds(searchTags)
		if err != nil {
			return nil, err
		} else {
			return tags, nil
		}
	}

	return tags, nil
}

func getMetaList(id int) (map[string]interface{}, error) {
	metaList, err := models.GetBlogContentMetaList(id)
	if err != nil {
		return nil, err
	}

	// "meta" => "meta_key" => val... でアクセスできるようにする
	_metaList := make(map[string]interface{})
	for _, meta := range metaList {
		_metaList[meta.MetaKey] = meta.MetaValue
	}

	return _metaList, nil
}

type Comment struct {
	Id       int       `json:"id"`
	IdReplay int       `json:"id_replay"`
	UserName string    `json:"user_name"`
	Comment  string    `json:"comment"`
	Good     int       `json:"good"`
	Time     string    `json:"time"`
	Status   int       `json:"status"`
	Children []Comment `json:"children"`
}

func getCommentList(id int) ([]Comment, error) {
	commentList, err := models.GetCommentListById(id)
	if err != nil {
		return nil, err
	}

	var _commentList []Comment

	// commentを親子関係でまとめる
	for _, comment := range commentList {
		time, _ := time.Parse(time.RFC3339, comment.CommentTime)

		// status見て内容をマスク
		if comment.Status == constant.COMMENT_WAITING_APPROVAL {
			comment.UserName = "承認待ちユーザー"
			comment.Comment = "このコメントは承認待ちです。"
		}

		_comment := Comment{
			Id:       int(comment.Id),
			IdReplay: comment.IdReplay,
			UserName: comment.UserName,
			Comment:  comment.Comment,
			Good:     comment.Good,
			Time:     time.Format(constant.DISPLAY_TIME_LAYOUT),
			Status:   comment.Status,
			Children: make([]Comment, 0),
		}

		// 返信先が0なら親データ
		if comment.IdReplay == 0 {
			_commentList = append(_commentList, _comment)
		} else {
			// それ以外なら既存データを再帰で探索して紐付ける
			searchComment(_commentList, _comment)
		}
	}
	return _commentList, nil
}

func searchComment(commentList []Comment, targetData Comment) {
	for index := range commentList {
		// 親IDを見つけたら子に自分を格納して返却
		if targetData.IdReplay == commentList[index].Id {
			commentList[index].Children = append(commentList[index].Children, targetData)
			return
		} else {
			// 見つけられないかつ子があれば再帰
			if len(commentList[index].Children) > 0 {
				searchComment(commentList[index].Children, targetData)
			}
			// 見つけられないかつ子が空なら次へ
		}
	}
}
