package request

import (
	customValidation "cms/package/validation"

	validation "github.com/go-ozzo/ozzo-validation"
)

type GetArticleListRequest struct {
	Keyword  string `json:"keyword"`
	PageNo   *int   `json:"page_no"`
	Tags     []int  `json:"tags"`
	Statuses []int  `json:"statuses"`
	SubSites []int  `json:"sub_sites"`
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
}

func (r GetArticleListRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Limit,
			validation.By(customValidation.Numeric),
		),
	)
}
