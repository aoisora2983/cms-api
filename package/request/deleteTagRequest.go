package request

import validation "github.com/go-ozzo/ozzo-validation"

type DeleteTagRequest struct {
	Id int `json:"id"`
}

func (r DeleteTagRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Id,
			validation.Required.Error("IDは必須です"),
		),
	)
}
