package models

import (
	"cms/db"

	"gorm.io/gorm"
)

type Accessibility struct {
	gorm.Model
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Message   string `json:"message"`
	Level     int    `json:"level"`
	IsReplace int    `json:"is_replace"`
}

type Tabler interface {
	TableName() string
}

func (Accessibility) TableName() string {
	return "accessibility"
}

func GetAccessibilityList() ([]Accessibility, error) {
	database := db.GetDB()

	var accessibility []Accessibility

	result := database.Select("*").
		Find(&accessibility).
		Order("id")

	if result.Error != nil {
		return nil, result.Error
	}

	return accessibility, nil
}

func SaveAccessibility(accessibility Accessibility) error {
	database := db.GetDB()

	result := database.Updates(&accessibility)

	if err := result.Error; err != nil {
		return err
	}

	return nil
}
