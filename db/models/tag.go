package models

import (
	"cms/constant"
	"cms/db"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Id        int    `json:"id"`
	Name      string `json:"label"`
	SortOrder int    `json:"sort_order"`
	IconPath  string `json:"icon_path"`
}

func SaveTag(data map[string]interface{}) error {
	database := db.GetDB()

	sortOrder := make(map[string]interface{})
	database.Table("tags").Select("MAX(id) as sort_order").Take(&sortOrder)
	var _sortOrder int
	_sortOrder = 1
	if sortOrder["sort_order"] != nil {
		_sortOrder = int(sortOrder["sort_order"].(int32)) + 1
	}
	var id int
	id = _sortOrder
	if data["id"] != 0 {
		id = data["id"].(int)
	}

	content := Tag{
		Id:       id,
		Name:     data["name"].(string),
		IconPath: data["filename"].(string),
	}

	// updateなら並び順はアップデートしない
	var result *gorm.DB
	if data["id"] == 0 {
		content.SortOrder = _sortOrder
		result = database.Create(&content)
	} else {
		result = database.Updates(&content)
	}

	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func GetTagList() ([]Tag, error) {
	database := db.GetDB()

	var tags []Tag

	result := database.Select("id", "name", "icon_path").
		Find(&tags).
		Order("sort_order").
		Order("id")

	if result.Error != nil {
		return nil, result.Error
	}

	return tags, nil
}

func GetOpenTagList(idSubSite int) ([]Tag, error) {
	database := db.GetDB()

	var tags []Tag

	subQuery := database.Table("blog_tags").
		Select("1").
		Joins("INNER JOIN blog_contents ON blog_contents.id = blog_tags.id_blog_content AND blog_contents.id_branch = blog_tags.id_branch_blog_content").
		Where("blog_tags.id_tag = tags.id").
		Where("blog_tags.deleted_at IS NULL").
		Where("blog_contents.deleted_at IS NULL").
		Where("blog_contents.status = ?", constant.ARTICLE_OPEN).
		Where("blog_contents.published_start_time <= current_timestamp").
		Where("(blog_contents.published_end_time >= current_timestamp OR blog_contents.published_end_time IS NULL)").
		Where("id_sub_site = ?", idSubSite)

	result := database.Table("tags").
		Select("tags.id, tags.name, tags.icon_path, tags.sort_order").
		Where("tags.deleted_at IS NULL").
		Where("EXISTS (?)", subQuery).
		Order("tags.sort_order, tags.id").
		Find(&tags)

	if result.Error != nil {
		return nil, result.Error
	}

	return tags, nil
}

// タグ削除
func DeleteTag(id int) error {
	var tag Tag
	database := db.GetDB()

	database.Table("tags").Where("id = ?", id).Unscoped().Delete(&tag)

	return nil
}

func GetTagListByIds(ids []int32) ([]Tag, error) {
	database := db.GetDB()

	var tags []Tag

	result := database.Select("id", "name", "icon_path").
		Where("id IN ?", ids).
		Order("sort_order").
		Order("id").
		Find(&tags)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return tags, nil
}

func GetTag(id int) (Tag, error) {
	database := db.GetDB()

	var tag Tag

	result := database.Select("id", "name", "icon_path").
		Where("id = ?", id).
		Order("sort_order").
		Order("id").
		Find(&tag)

	if result.RowsAffected == 0 {
		return tag, result.Error
	}

	return tag, nil
}
