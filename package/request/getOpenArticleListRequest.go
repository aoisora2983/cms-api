package request

import (
	customValidation "cms/package/validation"

	validation "github.com/go-ozzo/ozzo-validation"
)

type GetOpenArticleListRequest struct {
	Keyword string `form:"keyword"`
	Tags    []int  `form:"tags"`
	Limit   int    `form:"limit"`
	Page    int    `form:"page"`
}

func (r GetOpenArticleListRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Limit,
			validation.By(customValidation.Numeric),
		),
	)
}
