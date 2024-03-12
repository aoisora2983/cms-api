package login

import (
	"cms/config"
	"cms/db/models"
	"cms/package/auth"
	code "cms/package/error"
	request "cms/package/request"
	"cms/package/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var json request.LoginRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		if validErr, ok := err.(response.ValidationError); ok {
			c.JSON(validErr.GetStatus(), validErr.GetResponse())
			return
		}
		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}

	systemUser, err := models.CheckAuth(json.Mail, json.Password)

	// エラー有
	if err != nil {
		response.CustomErrorResponse(
			c,
			http.StatusUnauthorized,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}

	// 認証ユーザーが存在しない場合
	if systemUser.Id == 0 {
		response.CustomErrorResponse(
			c,
			http.StatusUnauthorized,
			map[string]string{code.NOT_EXISTS: "存在しないユーザーです。"},
		)
		return
	}

	token, err := auth.GenerateToken(systemUser.Id, systemUser.Name, systemUser.Mail, systemUser.IconPath)
	if err != nil {
		response.CustomErrorResponse(
			c,
			http.StatusInternalServerError,
			map[string]string{code.NOT_EXISTS: "認証に失敗しました。"},
		)
		return
	}

	if config.IsLocal() {
		c.SetCookie("authToken", token, 36000, "/", config.AppDomain(), false, false)
	} else {
		c.SetCookie("authToken", token, 36000, "/", config.AppDomain(), true, false)
		c.SetSameSite(http.SameSiteStrictMode)
	}
	c.JSON(http.StatusCreated, gin.H{
		"mail":  systemUser.Mail,
		"token": token,
	})
}
