package models

import (
	"cms/db"
	"cms/package/request"

	"gorm.io/gorm"
)

type SubSite struct {
	Id          int    `json:"id"`
	SubdirName  string `json:"subdir_name"`
	IconText    string `json:"icon_text"`
	Status      int    `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`
	SortOrder   int    `json:"sort_order"`
}

func GetSubSiteById(id int) (SubSite, error) {
	database := db.GetDB()

	var subSite SubSite

	result := database.Find(&subSite).
		Where("id = ?", id).
		Order("sort_order").
		Order("id")

	if result.Error != nil {
		return subSite, result.Error
	}

	return subSite, nil
}

func GetSubSiteByDirname(dirname string) (SubSite, error) {
	database := db.GetDB()

	var subSite SubSite

	result := database.Find(&subSite).
		Where("subdir_name = ?", dirname).
		Order("sort_order").
		Order("id")

	if result.Error != nil {
		return subSite, result.Error
	}

	return subSite, nil
}

func GetSubSiteList(status *int) ([]SubSite, error) {
	database := db.GetDB()

	var subSite []SubSite

	result := database.Find(&subSite).
		Order("sort_order").
		Order("id")

	if status != nil {
		result.Where("status = ?", status)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return subSite, nil
}

func SaveSubSite(data request.RegisterSubSiteRequest) error {
	database := db.GetDB()

	sortOrder := make(map[string]interface{})
	database.Table("sub_sites").Select("MAX(id) as sort_order").Take(&sortOrder)
	var _sortOrder int
	_sortOrder = 1
	if sortOrder["sort_order"] != nil {
		_sortOrder = int(sortOrder["sort_order"].(int32)) + 1
	}
	var id int
	id = _sortOrder
	if data.Id != 0 {
		id = data.Id
	}

	content := SubSite{
		Id:          id,
		SubdirName:  data.SubdirName,
		IconText:    data.IconText,
		Status:      data.Status,
		Title:       data.Title,
		Description: data.Description,
	}

	// updateなら並び順はアップデートしない
	var result *gorm.DB
	if data.Id == 0 {
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

func DeleteSubSite(id int) error {
	var subSite SubSite
	database := db.GetDB()

	database.Table("sub_sites").Where("id = ?", id).Unscoped().Delete(&subSite)

	return nil
}
