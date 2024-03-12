package response

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type ErrorMessages struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

type ResponseError struct {
	ErrorMessages []ErrorMessages `json:"errorMessages"`
}

type ValidationError struct {
	Response ResponseError
	Status   int
	Err      error
}

func (v ValidationError) GetResponse() ResponseError {
	return v.Response
}

func (v ValidationError) GetStatus() int {
	return v.Status
}

func (v ValidationError) Error() string {
	return v.Err.Error()
}

func CustomErrorResponse(
	c *gin.Context,
	errCode int,
	errorMessages map[string]string,
) {
	var params []ErrorMessages
	for key, value := range errorMessages {
		params = append(
			params,
			ErrorMessages{Name: key, Reason: value},
		)
	}

	c.JSON(errCode, ValidationError{
		Response: ResponseError{ErrorMessages: params},
		Status:   errCode,
	})
}
