package request

import (
	customValidation "cms/package/validation"

	validation "github.com/go-ozzo/ozzo-validation"
)

type GetArticleRequest struct {
	Id       int `form:"id"`
	IdBranch int `form:"id_branch"`
}

func (r GetArticleRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Id,
			validation.By(customValidation.Numeric),
		),
		validation.Field(
			&r.IdBranch,
			validation.By(customValidation.Numeric),
		),
	)
}
