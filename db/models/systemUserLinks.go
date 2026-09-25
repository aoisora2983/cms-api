package models

import (
	"cms/db"
	"cms/package/request"
)

type SystemUserLink struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Url       string `json:"url"`
	IconPath  string `json:"icon_path"`
	Alt       string `json:"alt"`
	SortOrder int    `json:"sort_order"`
}

func GetUserLinks(userId int) ([]SystemUserLink, error) {
	database := db.GetDB()

	var linkList []SystemUserLink

	result := database.
		Where("user_id = ?", userId).
		Order("id ASC").
		Find(&linkList)

	if result.Error != nil {
		return nil, result.Error
	}

	return linkList, nil
}

func SaveUserLink(userId int, data []request.UserLinks) error {
	// 一度既存データを削除し登録
	err := DeleteUserLinkByUserId(userId)
	if err != nil {
		return err
	}

	database := db.GetDB()

	for i := 0; i < len(data); i++ {
		content := SystemUserLink{
			UserId:    userId,
			Url:       data[i].Url,
			IconPath:  data[i].IconPath,
			Alt:       data[i].Alt,
			SortOrder: i,
		}

		result := database.Create(&content)

		if err := result.Error; err != nil {
			return err
		}
	}

	return nil
}

func DeleteUserLinkByUserId(userId int) error {
	database := db.GetDB()

	result := database.Where("user_id = ?", userId).
		Unscoped().
		Delete(&SystemUserLink{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
