package helper

import (
	code "cms/package/error"
	"cms/package/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BindRequest リクエストをバインディングし、エラーがあれば処理する
func BindRequest(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		if validErr, ok := err.(response.ValidationError); ok {
			c.JSON(validErr.GetStatus(), validErr.GetResponse())
			return false
		}

		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return false
	}
	return true
}

// BindQuery クエリパラメータをバインディングし、エラーがあれば処理する
func BindQuery(c *gin.Context, req interface{}) bool {
	if err := c.Bind(req); err != nil {
		if validErr, ok := err.(response.ValidationError); ok {
			c.JSON(validErr.GetStatus(), validErr.GetResponse())
			return false
		}

		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return false
	}
	return true
}

// GetUserIDFromContext コンテキストからユーザーIDを取得する
func GetUserIDFromContext(c *gin.Context) (int, bool) {
	cUserId, exists := c.Get("userId")
	if !exists {
		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.NOT_EXISTS: "存在しないユーザーです。"},
		)
		return 0, false
	}

	userIdStr, ok := cUserId.(string)
	if !ok {
		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.SERVER_ERROR: "ユーザーIDの形式が不正です。"},
		)
		return 0, false
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.SERVER_ERROR: "ユーザーIDの変換に失敗しました。"},
		)
		return 0, false
	}

	return userId, true
}


