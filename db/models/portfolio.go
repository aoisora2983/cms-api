package models

import (
	"cms/constant"
	"cms/db"

	"gorm.io/gorm"
)

type Portfolio struct {
	gorm.Model
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	DetailUrl   string `json:"detail_url"`
	ReleaseTime string `json:"release_time"`
	Status      int    `json:"status"`
	SortOrder   int    `json:"sort_order"`
}

func SavePortfolio(data map[string]interface{}) error {
	database := db.GetDB()

	sortOrder := make(map[string]interface{})
	database.Table("portfolios").Select("MAX(id) as sort_order").Take(&sortOrder)
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

	content := Portfolio{
		Id:          id,
		Title:       data["title"].(string),
		Description: data["description"].(string),
		Thumbnail:   data["thumbnail"].(string),
		DetailUrl:   data["detail_url"].(string),
		ReleaseTime: data["release_time"].(string),
		Status:      data["status"].(int),
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

func GetPortfolioList() ([]Portfolio, error) {
	database := db.GetDB()

	var portfolios []Portfolio

	result := database.Find(&portfolios).
		Order("sort_order").
		Order("id")

	if result.Error != nil {
		return nil, result.Error
	}

	return portfolios, nil
}

func GetOpenPortfolioList() ([]Portfolio, error) {
	database := db.GetDB()

	var portfolios []Portfolio

	result := database.Select("id", "title", "description", "thumbnail", "detail_url", "release_time").
		Find(&portfolios).
		Where("status = ?", constant.PORTFOLIO_OPEN).
		Where("release_time >= current_timestamp").
		Order("sort_order").
		Order("id")

	if result.Error != nil {
		return nil, result.Error
	}

	return portfolios, nil
}

func DeletePortfolio(id int) error {
	var tag Tag
	database := db.GetDB()

	database.Table("portfolios").Where("id = ?", id).Unscoped().Delete(&tag)

	return nil
}
