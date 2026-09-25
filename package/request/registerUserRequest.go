package request

import validation "github.com/go-ozzo/ozzo-validation"

type RegisterUserRequest struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Password    string      `json:"password"`
	Mail        string      `json:"mail"`
	GroupId     int         `json:"group_id"`
	Description string      `json:"description"`
	Filename    string      `json:"filename"`
	Links       []UserLinks `json:"links"`
}

type UserLinks struct {
	IconPath string `json:"icon_path"`
	Url      string `json:"url"`
	Alt      string `json:"alt"`
}

func (r RegisterUserRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Name,
			validation.Required.Error("ユーザー名は必須項目です。"),
		),
	)
}
