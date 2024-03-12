package request

import (
	customValidation "cms/package/validation"

	validation "github.com/go-ozzo/ozzo-validation"
)

type PostCommentRequest struct {
	ArticleId int    `json:"article_id"`
	ReplayId  int    `json:"replay_id"`
	UserName  string `json:"user_name"`
	Comment   string `json:"comment"`
}

func (r PostCommentRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.ArticleId,
			validation.By(customValidation.Numeric),
		),
	)
}
