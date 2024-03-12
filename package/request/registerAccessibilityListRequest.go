package request

import validation "github.com/go-ozzo/ozzo-validation"

type Accessibility struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
	Level   int    `json:"level"`
}

type RegisterAccessibilityListRequest struct {
	AccessibilityList []Accessibility
}

func (r RegisterAccessibilityListRequest) Validate() error {
	var validError error = nil
	for _, accessibility := range r.AccessibilityList {
		validError = validation.ValidateStruct(&accessibility,
			validation.Field(
				&accessibility.Message,
				validation.Required.Error("エラーメッセージは必須項目です。"),
			),
		)

		if validError != nil {
			return validError
		}
	}

	return validError
}
