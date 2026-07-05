package request

import (
	customValidation "cms/package/validation"

	validation "github.com/go-ozzo/ozzo-validation"
)

type GetUniqueArticleRequest struct {
	PageType int `form:"page_type"`
}

func (r GetUniqueArticleRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.PageType,
			validation.By(customValidation.Numeric),
		),
	)
}
