package request

import validation "github.com/go-ozzo/ozzo-validation"

type RegisterPortfolioRequest struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	DetailUrl   string `json:"detail_url"`
	ReleaseTime string `json:"release_time"`
	Status      int    `json:"status"`
	SortOrder   int    `json:"sort_order"`
}

func (r RegisterPortfolioRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Title,
			validation.Required.Error("サービス名は必須項目です。"),
		),
	)
}
