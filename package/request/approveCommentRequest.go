package request

import validation "github.com/go-ozzo/ozzo-validation"

type ApproveCommentRequest struct {
	Ids []int `json:"ids"`
}

func (r ApproveCommentRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Ids,
			validation.Required.Error("承認対象ID一覧は必須項目です。"),
		),
	)
}
