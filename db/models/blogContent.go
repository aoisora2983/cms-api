package models

import (
	"cms/constant"
	"cms/db"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BlogContent struct {
	gorm.Model
	Id                   int     `gorm:"autoIncrement:"false";column:id" json:"id"`
	IdBranch             int     `gorm:"column:id_branch" json:"id_branch"`
	IdUser               int     `gorm:"column:id_user" json:"id_user"`
	Title                string  `gorm:"column:title" json:"title"`
	Content              string  `gorm:"column:content" json:"content"`
	Status               int     `gorm:"column:status" json:"status"`
	Thumbnail            string  `gorm:"column:thumbnail" json:"thumbnail"`
	PublishedStartTime   string  `gorm:"column:published_start_time" json:"published_start_time"`
	PublishedEndTime     *string `gorm:"column:published_end_time" json:"published_end_time"`
	PublishedUpdatedTime *string `gorm:"column:published_updated_time" json:"published_updated_time"`
	Description          string  `gorm:"column:description" json:"description"`
}

type BlogContentParam struct {
	Id       *int
	IdBranch *int
	Keyword  string
	Tags     []int
	IsOpen   bool
	Limit    int
	Offset   int
}

func GetLatestId() (int, error) {
	var model BlogContent
	var result struct {
		Id int
	}

	database := db.GetDB()

	err := database.Unscoped().Model(&model).Select("MAX(id) as id").Find(&result).Error

	if err != nil {
		return result.Id, err
	}

	if result.Id == 0 {
		return 0, nil
	}
	return result.Id, nil
}

func InsertContent(data map[string]interface{}) error {
	database := db.GetDB()

	if err := database.Create(&BlogContent{
		Id:                 data["id"].(int),
		IdBranch:           data["id_branch"].(int),
		IdUser:             data["id_user"].(int),
		Title:              data["title"].(string),
		Content:            data["content"].(string),
		Status:             data["status"].(int),
		Thumbnail:          data["thumbnail"].(string),
		PublishedStartTime: data["published_start_time"].(string),
		PublishedEndTime:   data["published_end_time"].(*string),
		Description:        data["description"].(string),
	}).Error; err != nil {
		return err
	}

	return nil
}

func UpdateContent(data map[string]interface{}) error {
	database := db.GetDB()
	updateTime := time.Now().Format(time.RFC3339)

	if err := database.Model(BlogContent{}).
		Where("id = ?", data["id"]).
		Where("id_branch = ?", data["id_branch"]).
		Updates(BlogContent{
			IdUser:               data["id_user"].(int),
			Title:                data["title"].(string),
			Content:              data["content"].(string),
			Status:               data["status"].(int),
			Thumbnail:            data["thumbnail"].(string),
			PublishedStartTime:   data["published_start_time"].(string),
			PublishedEndTime:     data["published_end_time"].(*string),
			PublishedUpdatedTime: &updateTime,
			Description:          data["description"].(string),
		}).Error; err != nil {
		return err
	}

	return nil
}

/**
 * 記事一覧取得
 */
func GetBlogContentList(param BlogContentParam) ([]map[string]interface{}, error) {
	database := db.GetDB()

	// 記事一覧取得
	query := database.Table("blog_contents").
		Distinct(
			"blog_contents.id",
			"blog_contents.id_branch",
			"blog_contents.id_user",
			"blog_contents.title",
			"blog_contents.content",
			"blog_contents.status",
			"blog_contents.thumbnail",
			"blog_contents.published_start_time",
			"blog_contents.published_end_time",
			"blog_contents.published_updated_time",
			"blog_contents.description",
			"system_users.name as user_name",
			"system_users.description as user_description",
			"system_users.icon_path as user_icon_path",
			"COUNT(blog_contents.*) OVER() as total",
		).
		Joins("LEFT JOIN blog_tags ON blog_contents.id = blog_tags.id_blog_content AND blog_contents.id_branch = blog_tags.id_branch_blog_content").
		Joins("LEFT JOIN system_users ON blog_contents.id_user = system_users.id").
		Where("blog_contents.deleted_at IS NULL").
		Order("blog_contents.published_start_time desc").
		Order("blog_contents.id")

	if param.Id != nil {
		query.Where("blog_contents.id = ?", param.Id)
	}

	if param.IdBranch != nil {
		query.Where("blog_contents.id_branch = ?", param.IdBranch)
	}

	if param.Keyword != "" {
		keywordLike := fmt.Sprintf("%%%s%%", param.Keyword)
		query.Where(database.Where("blog_contents.title LIKE ?", keywordLike).Or("blog_contents.content LIKE ?", keywordLike))
	}

	if len(param.Tags) > 0 {
		query.Where("blog_tags.id_tag IN ?", param.Tags)
	}

	if param.IsOpen {
		// 公開中
		query.Where("blog_contents.status = ?", constant.ARTICLE_OPEN)
		query.Where("blog_contents.published_start_time <= current_timestamp")
		query.Where("(blog_contents.published_end_time >= current_timestamp OR blog_contents.published_end_time IS NULL)")
	}

	if param.Limit != 0 {
		query.Limit(param.Limit)
	}

	if param.Offset != 0 {
		query.Offset(param.Offset)
	}

	var contents []map[string]interface{}
	query.Find(&contents)

	if query.RowsAffected == 0 {
		return nil, nil
	}

	return contents, nil
}

// 記事取得
func GetBlogContent(id int, idBranch int, isOpen bool) (map[string]interface{}, error) {
	database := db.GetDB()

	query := database.Table("blog_contents").
		Select(
			"blog_contents.id",
			"blog_contents.id_branch",
			"blog_contents.id_user",
			"blog_contents.title",
			"blog_contents.content",
			"blog_contents.status",
			"blog_contents.thumbnail",
			"blog_contents.published_start_time",
			"blog_contents.published_end_time",
			"blog_contents.published_updated_time",
			"blog_contents.description",
			"system_users.name as user_name",
			"system_users.description as user_description",
			"system_users.icon_path as user_icon_path",
			"COUNT(blog_contents.*) OVER()",
		).
		Joins("LEFT JOIN system_users ON blog_contents.id_user = system_users.id").
		Where("blog_contents.deleted_at IS NULL")

	// 指定IDで
	query.Where("blog_contents.id = ?", id)

	if idBranch != -1 {
		query.Where("blog_contents.id_branch = ?", idBranch)
	}

	if isOpen {
		// 公開中
		query.Where("blog_contents.status = ?", constant.ARTICLE_OPEN)
		query.Where("blog_contents.published_start_time <= current_timestamp")
		query.Where("(blog_contents.published_end_time >= current_timestamp OR blog_contents.published_end_time IS NULL)")
	}

	content := make(map[string]interface{})
	query.Take(content)

	if query.Error != nil {
		return nil, query.Error
	}

	return content, nil
}

func DeleteContent(id int, idBranch int) error {
	database := db.GetDB()

	var model BlogContent

	result := database.Table("blog_contents").
		Where("id", id).
		Where("id_branch", idBranch).
		Delete(&model)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 指定の枝番より前のデータを削除する
func DeleteContentsBeforeBranch(id int, idBranch int) error {
	database := db.GetDB()

	var model BlogContent

	result := database.Table("blog_contents").
		Where("id", id).
		Where("id_branch < ?", idBranch).
		Delete(&model)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
