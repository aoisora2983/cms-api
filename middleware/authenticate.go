package middleware

import (
	"cms/package/auth"
	code "cms/package/error"
	"cms/package/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWT JWT認証ミドルウェア
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("authToken")
		if err != nil || token == "" {
			response.CustomErrorResponse(
				c,
				http.StatusUnauthorized,
				map[string]string{code.SERVER_ERROR: "認証に失敗しました。"},
			)
			c.Abort()
			return
		}

		subject, err := auth.Verify(token)
		if err != nil {
			response.CustomErrorResponse(
				c,
				http.StatusUnauthorized,
				map[string]string{code.SERVER_ERROR: "認証に失敗しました。"},
			)
			c.Abort()
			return
		}

		c.Set("userId", subject)
		c.Next()
	}
}
