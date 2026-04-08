package request

import (
	customValidation "cms/package/validation"

	validation "github.com/go-ozzo/ozzo-validation"
)

type GetOpenArticleListRequest struct {
	Keyword        string `json:"keyword"`
	Tags           []int  `json:"tags"`
	Limit          int    `json:"limit"`
	Page           int    `json:"page"`
	ExcludePageIds []int  `json:"exclude_page_ids"`
}

func (r GetOpenArticleListRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Limit,
			validation.By(customValidation.Numeric),
		),
	)
}
