package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type RegisterReplaceWordRequest struct {
	Id              int    `json:"id"`
	IdAccessibility int    `json:"id_accessibility"`
	WordFrom        string `json:"word_from"`
	WordTo          string `json:"word_to"`
	Level           int    `json:"level"`
}

func (r RegisterReplaceWordRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.WordFrom,
			validation.Required.Error("置換元の単語を入力してください。"),
		),
		validation.Field(
			&r.WordTo,
			validation.Required.Error("置換先の単語を入力してください。"),
		),
	)
}
