package login

import (
	"cms/config"
	"cms/package/helper"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", config.AppDomain(), false, true)
	helper.OKResponse(c)
}
