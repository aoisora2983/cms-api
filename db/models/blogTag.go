package models

import (
	"cms/db"
	"time"

	"gorm.io/gorm"
)

type BlogTag struct {
	IdBlogContent       int `json:"id_blog_content"`
	IdBranchBlogContent int `json:"id_branch_blog_content"`
	IdTag               int `json:"id_tag"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}

// 記事のID,枝番からタグのID一覧取得
func GetBlogTags(id int32, idBranch int32) ([]map[string]interface{}, error) {
	database := db.GetDB()

	query := database.Table("blog_tags").
		Select("blog_tags.id_tag, tags.name, tags.icon_path").
		Joins("LEFT JOIN tags ON blog_tags.id_tag = tags.id").
		Where("id_blog_content = ?", id).
		Where("id_branch_blog_content", idBranch).
		Order("tags.sort_order")

	var tags []map[string]interface{}
	query.Find(&tags)

	if query.Error != nil {
		return nil, query.Error
	}

	if query.RowsAffected == 0 {
		return nil, nil
	}

	return tags, nil
}

func InsertBlogTag(data map[string]interface{}) error {
	content := BlogTag{
		IdBlogContent:       data["id_content"].(int),
		IdBranchBlogContent: data["id_branch_content"].(int),
		IdTag:               data["id_tag"].(int),
	}

	database := db.GetDB()
	if err := database.Create(&content).Error; err != nil {
		return err
	}

	return nil
}

func DeleteBlogTag(id int, idBranch int) error {
	var model BlogTag

	database := db.GetDB()

	result := database.Model(BlogTag{}).Unscoped().
		Where("id_blog_content = ?", id).
		Where("id_branch_blog_content = ?", idBranch).
		Delete(&model)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
