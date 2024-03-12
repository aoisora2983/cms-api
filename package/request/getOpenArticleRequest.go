package request

import (
	customValidation "cms/package/validation"

	validation "github.com/go-ozzo/ozzo-validation"
)

type GetOpenArticleRequest struct {
	Id       int `form:"id"`
	IdBranch int `form:"id_branch"`
}

func (r GetOpenArticleRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Id,
			validation.By(customValidation.Numeric),
		),
	)
}
