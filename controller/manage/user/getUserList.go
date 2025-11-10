package user

import (
	"cms/db/models"
	"cms/package/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	userList, err := models.GetUserList()
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.CreatedResponse(c, userList)
}
