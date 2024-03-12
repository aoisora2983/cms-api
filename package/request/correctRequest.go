package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type CorrectRequest struct {
	Sentence string `json:"sentence"`
	SkipIds  []int  `json:"skip_ids"`
}

func (r CorrectRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Sentence,
			validation.Required.Error("校正内容は必須項目です。"),
		),
	)
}
