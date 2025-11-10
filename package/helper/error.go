package helper

import (
	code "cms/package/error"
	"cms/package/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleError エラーを処理してレスポンスを返す
func HandleError(c *gin.Context, err error, statusCode int) {
	response.CustomErrorResponse(
		c,
		statusCode,
		map[string]string{code.SERVER_ERROR: err.Error()},
	)
}

// HandleErrorWithMessage カスタムメッセージでエラーを処理してレスポンスを返す
func HandleErrorWithMessage(c *gin.Context, message string, statusCode int) {
	response.CustomErrorResponse(
		c,
		statusCode,
		map[string]string{code.SERVER_ERROR: message},
	)
}

// HandleNotExistsError 存在しないエラーを処理してレスポンスを返す
func HandleNotExistsError(c *gin.Context, message string) {
	response.CustomErrorResponse(
		c,
		http.StatusUnauthorized,
		map[string]string{code.NOT_EXISTS: message},
	)
}
