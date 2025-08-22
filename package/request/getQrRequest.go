package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type GetQrRequest struct {
	Content        string   `json:"content"`
	LogoImage      string   `json:"logo_image_path"`
	HalftoneImage  string   `json:"halftone_image_path"`
	FgColor        []string `json:"fg_color"`
	BgColor        string   `json:"bg_color"`
	BgTransparent  bool     `json:"bg_transparent"`
	DotType        int      `json:"dot_type"`
	ImageExtension int      `json:"image_extension"`
}

func (r GetQrRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Content,
			validation.Required.Error("QRコードの内容は必ず入力してください。"),
		),
	)
}
