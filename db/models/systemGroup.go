package models

import (
	"cms/db"

	"gorm.io/gorm"
)

type SystemGroup struct {
	gorm.Model
	Id           int32  `json:"id"`
	Name         string `json:"name"`
	EditBlog     int    `json:"edit_blog"`
	EditCategory int    `json:"edit_category"`
	EditTag      int    `json:"edit_tag"`
	EditUser     int    `json:"edit_user"`
	Admin        int16  `json:"admin"`
}

func GetSystemGroupList() ([]SystemGroup, error) {
	database := db.GetDB()

	var systemGroups []SystemGroup

	result := database.Select("id", "name", "edit_blog", "edit_category", "edit_tag", "edit_user").
		Find(&systemGroups).
		Order("sort_order").
		Order("id")

	if result.Error != nil {
		return nil, result.Error
	}

	return systemGroups, nil
}
