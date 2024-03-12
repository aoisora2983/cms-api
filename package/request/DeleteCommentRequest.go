package request

import validation "github.com/go-ozzo/ozzo-validation"

type DeleteCommentRequest struct {
	Ids []int `json:"ids"`
}

func (r DeleteCommentRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Ids,
			validation.Required.Error("削除対象ID一覧は必須項目です。"),
		),
	)
}
