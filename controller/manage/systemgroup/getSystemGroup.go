package systemgroup

import (
	"cms/db/models"
	"cms/package/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSystemGroup(c *gin.Context) {
	userList, err := models.GetSystemGroupList()
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.CreatedResponse(c, gin.H{
		"list": userList,
	})
}
