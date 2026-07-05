package models

import (
	"cms/constant"
	"cms/db"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CorrectWord struct {
	gorm.Model
	Id              int    `json:"id"`
	IdAccessibility int    `id_accessibility`
	WordFrom        string `json:"word_from"`
	WordTo          string `json:"word_to"`
	Level           int    `json:"level"`
}

func GetReplaceableCorrectWordListById(accessibilityId int) ([]CorrectWord, error) {
	database := db.GetDB()

	var correctWords []CorrectWord

	result := database.Select("*").
		Where("id_accessibility = ?", accessibilityId).
		Where("level <> ?", constant.CORRECT_NO_CHECK).
		Order("id").
		Find(&correctWords)

	if result.Error != nil {
		return nil, result.Error
	}

	return correctWords, nil
}

func GetCorrectWordListById(accessibilityId int, keyword string, alertLevel []int) ([]CorrectWord, error) {
	database := db.GetDB()

	var correctWords []CorrectWord

	query := database.Select("*").
		Where("id_accessibility = ?", accessibilityId).
		Order("id")

	if keyword != "" {
		likeWord := "%" + keyword + "%"
		query.Where(database.Where("word_from LIKE ?", likeWord).Or("word_to LIKE ?", likeWord))
	}

	if len(alertLevel) > 0 {
		query.Where("level IN ?", alertLevel)
	}

	result := query.Find(&correctWords)
	if result.Error != nil {
		return nil, result.Error
	}

	return correctWords, nil
}

func SaveCorrectWord(correctWord CorrectWord) error {
	database := db.GetDB()

	result := database.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"word_from", "word_to", "level"}),
	}).Create(&correctWord)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
