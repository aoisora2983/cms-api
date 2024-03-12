package manage

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
