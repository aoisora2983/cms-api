package request

import (
	customValidation "cms/package/validation"

	validation "github.com/go-ozzo/ozzo-validation"
)

type CountUpCommentGoodRequest struct {
	Id int `form:"id"`
}

func (r CountUpCommentGoodRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Id,
			validation.By(customValidation.Numeric),
		),
	)
}
