package middleware

import (
	"cms/package/auth"
	code "cms/package/error"
	"cms/package/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var isOk bool

		token, err := c.Cookie("authToken")
		if err != nil {
			isOk = false
		} else {
			if token == "" {
				isOk = false
			} else {
				subject, err := auth.Verify(token)
				if err != nil {
					isOk = false

				} else {
					isOk = true
					c.Set("userId", subject)
				}
			}
		}

		if !isOk {

			response.CustomErrorResponse(
				c,
				http.StatusUnauthorized,
				map[string]string{code.SERVER_ERROR: "認証に失敗しました。"},
			)

			c.Abort()
			return
		}

		c.Next()
	}
}
