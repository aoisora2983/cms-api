package models

import (
	"cms/constant"
	"cms/db"
	"strconv"

	"gorm.io/gorm"
)

type BlogContentMeta struct {
	gorm.Model
	Id            int    `gorm:"autoIncrement:"false";column:id" json:"id"`
	IdBlogContent int    `gorm:"column:id_blog_content" json:"id_blog_content"`
	MetaKey       string `gorm:"column:meta_key" json:"meta_key"`
	MetaValue     string `gorm:"column:meta_value" json:"meta_value"`
}

func CountUpArticleGood(id int) error {
	database := db.GetDB()

	var model BlogContentMeta
	result := database.
		Where("meta_key = ?", constant.BLOG_META_ARTICLE_GOOD_KEY).
		Where("id_blog_content = ?", id).
		Find(&model)

	if result.Error != nil {
		return result.Error
	}

	// 空ならINSERT
	if result.RowsAffected == 0 {
		result := database.Create(&BlogContentMeta{
			IdBlogContent: id,
			MetaKey:       constant.BLOG_META_ARTICLE_GOOD_KEY,
			MetaValue:     "1",
		})

		if result.Error != nil {
			return result.Error
		}
	} else {
		good, _ := strconv.Atoi(model.MetaValue)
		// それ以外ならUPDATE
		result := database.Model(&BlogContentMeta{}).
			Where("meta_key = ?", constant.BLOG_META_ARTICLE_GOOD_KEY).
			Where("id_blog_content = ?", id).
			Update("meta_value", strconv.Itoa(good+1))

		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func GetBlogContentMetaList(id int) ([]BlogContentMeta, error) {
	database := db.GetDB()

	var metaList []BlogContentMeta

	result := database.Where("id_blog_content = ?", id).
		Find(&metaList)

	if result.Error != nil {
		return nil, result.Error
	}

	return metaList, nil
}
