package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SuccessResponse 成功レスポンスを返す
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// CreatedResponse 作成成功レスポンスを返す
func CreatedResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}

// OKResponse 成功ステータスのみのレスポンスを返す
func OKResponse(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

