package request

import (
	customValidation "cms/package/validation"

	validation "github.com/go-ozzo/ozzo-validation"
)

type LoginRequest struct {
	Mail     string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r LoginRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Mail,
			validation.Required.Error("メールアドレスを入力してください。"),
			validation.By(customValidation.EmailFormat),
		),
		validation.Field(
			&r.Password,
			validation.Required.Error("パスワードを入力してください。"),
		),
	)
}
