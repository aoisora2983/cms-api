package user

import (
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	var req request.DeleteUserRequest
	if !helper.BindRequest(c, &req) {
		return
	}

	err := models.DeleteUser(req.Id)
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.OKResponse(c)
}
