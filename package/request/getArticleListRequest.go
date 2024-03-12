package request

import (
	customValidation "cms/package/validation"

	validation "github.com/go-ozzo/ozzo-validation"
)

type GetArticleListRequest struct {
	Keyword  string `form:"keyword"`
	Tags     []int  `form:"tags"`
	Statuses []int  `form:"statuses"`
	Limit    int    `form:"limit"`
	Page     int    `form:"page"`
}

func (r GetArticleListRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Limit,
			validation.By(customValidation.Numeric),
		),
	)
}
