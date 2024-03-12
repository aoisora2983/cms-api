package models

import (
	"cms/constant"
	"cms/db"

	"gorm.io/gorm"
)

type CorrectWords struct {
	gorm.Model
	Id       int    `json:"id"`
	WordFrom string `json:"word_from"`
	WordTo   string `json:"word_to"`
	Level    int    `json:"level"`
}

func GetCorrectTargetWordList() ([]CorrectWords, error) {
	database := db.GetDB()

	var correctWords []CorrectWords

	result := database.Select("*").
		Find(&correctWords).
		Where("level <> ?", constant.CORRECT_OK).
		Order("id")

	if result.Error != nil {
		return nil, result.Error
	}

	return correctWords, nil
}
