package request

import (
	customValidation "cms/package/validation"

	validation "github.com/go-ozzo/ozzo-validation"
)

type GetOpenTagListRequest struct {
	IdSubSite int `form:"id_sub_site"`
}

func (r GetOpenTagListRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.IdSubSite,
			validation.By(customValidation.Numeric),
		),
	)
}
