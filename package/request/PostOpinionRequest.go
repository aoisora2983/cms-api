package request

import (
	customValidation "cms/package/validation"

	validation "github.com/go-ozzo/ozzo-validation"
)

type PostOpinionRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Content string `json:"content"`
}

func (r PostOpinionRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Name,
			validation.Required.Error("お名前を入力してください。"),
		),
		validation.Field(
			&r.Email,
			validation.Required.Error("メールアドレスを入力してください。"),
			validation.By(customValidation.EmailFormat),
		),
		validation.Field(
			&r.Content,
			validation.Required.Error("問い合わせ内容を入力してください。"),
		),
	)
}
