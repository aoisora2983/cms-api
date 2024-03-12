package models

import (
	"cms/db"

	"gorm.io/gorm"
)

type SystemUser struct {
	gorm.Model
	Id          int    `json:"id"`
	GroupId     int    `json:"group_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Mail        string `json:"mail"`
	IconPath    string `json:"icon_path"`
	Password    string `json:"password"`
	SortOrder   int    `json:"sort_order"`
}

func CheckAuth(mail, password string) (SystemUser, error) {
	var systemUser SystemUser

	database := db.GetDB()
	// TODO: パスワードはhashで持つ
	err := database.Select(
		"id",
		"name",
		"mail",
		"icon_path",
		"password",
	).Where(
		SystemUser{Mail: mail, Password: password},
	).Find(&systemUser).Error

	if err != nil {
		return systemUser, err
	}

	if systemUser.Id > 0 {
		return systemUser, nil
	}

	return systemUser, nil
}

func GetUserList() ([]SystemUser, error) {
	database := db.GetDB()

	var users []SystemUser

	result := database.Select("id", "name", "password", "description", "mail", "icon_path").
		Order("group_id").
		Order("sort_order").
		Order("id").
		Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func SaveUser(data map[string]interface{}) error {
	database := db.GetDB()

	sortOrder := make(map[string]interface{})
	database.Table("system_users").Select("MAX(id) as sort_order").Take(&sortOrder)
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

	content := SystemUser{
		Id:          id,
		GroupId:     data["group_id"].(int),
		Name:        data["name"].(string),
		Description: data["description"].(string),
		Mail:        data["mail"].(string),
		IconPath:    data["filename"].(string),
		Password:    data["password"].(string),
	}

	// updateなら並び順はアップデートしない
	var result *gorm.DB
	if data["id"].(int) == 0 {
		content.SortOrder = _sortOrder
		result = database.Create(&content)
	} else {
		content.Id = data["id"].(int)
		result = database.Updates(&content)
	}

	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(id int) error {
	database := db.GetDB()

	result := database.Where("id = ?", id).
		Unscoped().
		Delete(&SystemUser{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
