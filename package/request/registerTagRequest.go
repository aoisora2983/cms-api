package request

import validation "github.com/go-ozzo/ozzo-validation"

type RegisterTagRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Filename string `json:"filename"`
}

func (r RegisterTagRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Name,
			validation.Required.Error("タグ名は必須項目です。"),
		),
	)
}
