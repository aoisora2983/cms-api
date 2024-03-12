package request

import validation "github.com/go-ozzo/ozzo-validation"

type DeleteUserRequest struct {
	Id int `json:"id"`
}

func (r DeleteUserRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Id,
			validation.Required.Error("削除対象IDは必須項目です。"),
		),
	)
}
