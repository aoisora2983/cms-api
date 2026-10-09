package request

import validation "github.com/go-ozzo/ozzo-validation"

type RegisterSubSiteRequest struct {
	Id          int    `json:"id"`
	SubdirName  string `json:"sub_dir_name"`
	IconText    string `json:"icon_text"`
	Status      int    `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`
	SortOrder   int    `json:"sort_order"`
}

func (r RegisterSubSiteRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.SubdirName,
			validation.Required.Error("サブディレクトリ名は必須項目です。"),
		),
		validation.Field(
			&r.Title,
			validation.Required.Error("細部サイト名は必須です。"),
		),
	)
}
