package request

import (
	customValidation "cms/package/validation"

	validation "github.com/go-ozzo/ozzo-validation"
)

type GetSubSiteRequest struct {
	Id         *int    `form:"id"`
	SubDirName *string `form:"sub_dir_name"`
}

func (r GetSubSiteRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Id,
			validation.By(customValidation.Numeric),
		),
	)
}
