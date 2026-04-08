package login

import (
	"cms/config"
	"cms/db/models"
	"cms/package/auth"
	"cms/package/helper"
	request "cms/package/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req request.LoginRequest
	if !helper.BindRequest(c, &req) {
		return
	}

	systemUser, err := models.CheckAuth(req.Mail, req.Password)
	if err != nil {
		helper.HandleError(c, err, http.StatusUnauthorized)
		return
	}

	// 認証ユーザーが存在しない場合
	if systemUser.Id == 0 {
		helper.HandleNotExistsError(c, "存在しないユーザーです。")
		return
	}

	token, err := auth.GenerateToken(systemUser.Id, systemUser.Name, systemUser.Mail, systemUser.IconPath)
	if err != nil {
		helper.HandleErrorWithMessage(c, "認証に失敗しました。", http.StatusInternalServerError)
		return
	}

	if config.IsLocal() {
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("authToken", token, 36000, "/", "", false, false)
	} else {
		c.SetSameSite(http.SameSiteStrictMode)
		c.SetCookie("authToken", token, 36000, "/", config.AppDomain(), true, false)
	}

	helper.CreatedResponse(c, gin.H{
		"mail":  systemUser.Mail,
		"token": token,
	})
}
