package models

import (
	"cms/constant"
	"cms/db"
)

type BlogComment struct {
	Id            int    `gorm:"column:id" json:"id"`
	IdBlogContent int    `gorm:"column:id_blog_content" json:"id_blog_content"`
	IdReplay      int    `gorm:"column:id_replay" json:"id_replay"`
	UserName      string `gorm:"column:user_name" json:"user_name"`
	Comment       string `gorm:"column:comment" json:"comment"`
	Ip            string `gorm:"column:ip" json:"ip"`
	Good          int    `gorm:"column:good" json:"good"`
	Status        int    `gorm:"column:status" json:"status"`
	CommentTime   string `gorm:"column:comment_time" json:"comment_time"`
}

func GetCommentList() ([]BlogComment, error) {
	database := db.GetDB()

	var commentList []BlogComment

	result := database.Order("status = 0 ASC").
		Order("id_blog_content").
		Order("id_replay").
		Find(&commentList)

	if result.Error != nil {
		return nil, result.Error
	}

	return commentList, nil
}

func CountUpComment(id int) error {
	database := db.GetDB()

	var model BlogComment
	database.Where("id = ?", id).
		First(&model)

	// UPDATE
	result := database.Model(&BlogComment{}).
		Where("id = ?", id).
		Update("good", model.Good+1)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func InsertComment(data BlogComment) error {
	database := db.GetDB()

	if err := database.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func GetCommentListById(id int) ([]BlogComment, error) {
	database := db.GetDB()

	var commentList []BlogComment

	result := database.Where("id_blog_content = ?", id).
		Order("id_replay ASC").
		Order("id ASC").
		Find(&commentList)

	if result.Error != nil {
		return nil, result.Error
	}

	return commentList, nil
}

func ApproveComment(ids []int) error {
	database := db.GetDB()

	result := database.Model(&BlogComment{}).
		Where("id IN ?", ids).
		Update("status", constant.COMMENT_OPEN)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteComment(ids []int) error {
	database := db.GetDB()

	result := database.Where("id IN ?", ids).
		Delete(&BlogComment{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
