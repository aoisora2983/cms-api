package models

import (
	"cms/db"
)

type Opinion struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Content  string `json:"content"`
	IP       string `json:"ip"`
	SendTime string `json:"send_time"`
}

func GetOpinionList() ([]Opinion, error) {
	database := db.GetDB()

	var opinions []Opinion

	result := database.
		Order("id").
		Find(&opinions)

	if err := result.Error; err != nil {
		return nil, err
	}

	return opinions, nil
}

func SaveOpinion(data Opinion) error {
	database := db.GetDB()

	result := database.Create(&data)

	if err := result.Error; err != nil {
		return err
	}

	return nil
}
