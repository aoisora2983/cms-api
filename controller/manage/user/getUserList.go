package user

import (
	"cms/db/models"
	"cms/package/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	models.SystemUser
	Links []models.SystemUserLink `json:"links"`
}

func GetUserList(c *gin.Context) {
	userList, err := models.GetUserList()
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	response := make([]UserListResponse, len(userList))

	for i := 0; i < len(userList); i++ {
		// ユーザーリンク取得
		userLinks, err := models.GetUserLinks(userList[i].Id)
		if err != nil {
			helper.HandleError(c, err, http.StatusInternalServerError)
			return
		}

		response[i] = UserListResponse{
			SystemUser: userList[i],
			Links:      userLinks,
		}
	}

	helper.CreatedResponse(c, response)
}
