package manage

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Menu(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
