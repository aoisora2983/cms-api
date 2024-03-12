package login

import (
	"cms/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", config.AppDomain(), false, true)

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
