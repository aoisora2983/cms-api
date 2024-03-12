package request

import validation "github.com/go-ozzo/ozzo-validation"

type DeleteArticleRequest struct {
	Targets []string `json:"id_branch"`
}

func (r DeleteArticleRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Targets,
			validation.Required.Error("削除対象のid, 枝番は必須項目です。"),
		),
	)
}
