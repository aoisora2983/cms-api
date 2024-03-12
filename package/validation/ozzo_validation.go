package validation

import (
	"cms/package/response"
	"net/http"

	"github.com/gin-gonic/gin/binding"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Validator struct{}

func (v Validator) Validate(obj any) error {
	if obj == nil {
		return nil
	}

	val, ok := obj.(validation.Validatable)
	if !ok {
		return nil
	}
	if err := val.Validate(); err != nil {
		if verr, ok := err.(validation.Errors); ok {
			var params []response.ErrorMessages
			for key, val := range verr {
				params = append(params, response.ErrorMessages{Name: key, Reason: val.Error()})
			}
			return v.newValidationError(params, err)
		}
		return v.newServerError(err)
	}
	return nil
}

func (v Validator) newServerError(err error) error {
	return response.ValidationError{
		Response: response.ResponseError{},
		Status:   http.StatusInternalServerError,
		Err:      err,
	}
}

func (v Validator) newValidationError(params []response.ErrorMessages, err error) error {
	return response.ValidationError{
		Response: response.ResponseError{ErrorMessages: params},
		Status:   http.StatusBadRequest,
		Err:      err,
	}
}

type ozzoValidator struct {
	validator *Validator
}

func NewOzzoValidator() binding.StructValidator {
	return &ozzoValidator{validator: &Validator{}}
}

func (v *ozzoValidator) ValidateStruct(obj any) error {
	return v.validator.Validate(obj)
}

func (v *ozzoValidator) Engine() any {
	return v.validator
}
