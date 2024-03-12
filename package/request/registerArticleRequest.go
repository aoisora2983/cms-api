package request

import validation "github.com/go-ozzo/ozzo-validation"

type RegisterArticleRequest struct {
	Id                 *int   `json:"id"`
	IdBranch           *int   `json:"id_branch"`
	Title              string `json:"title"`
	Content            string `json:"content"`
	Tags               []int  `json:"tags"`
	Status             int    `json:"status"`
	Thumbnail          string `json:"thumbnail"`
	PublishedStartDate string `json:"published_start_date"`
	PublishedStartTime string `json:"published_start_time"`
	PublishedEndDate   string `json:"published_end_date"`
	PublishedEndTime   string `json:"published_end_time"`
	Description        string `json:"description"`
}

func (r RegisterArticleRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Title,
			validation.Required.Error("タイトルは必須項目です。"),
		),
		validation.Field(
			&r.Content,
			validation.Required.Error("本文は必須項目です。"),
		),
		// validation.Field(
		// 	&r.Categories,
		// 	validation.Required.Error("カテゴリは必須項目です。"),
		// ),
		// validation.Field(
		// 	&r.Tags,
		// 	validation.In("NOVELS", "COMIC", "FICTION", "NON_FICTION").Error("値が正しくありません。"),
		// ),
	)
}
