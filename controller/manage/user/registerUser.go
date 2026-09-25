package user

import (
	"cms/db/models"
	"cms/package/helper"
	"cms/package/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var req request.RegisterUserRequest
	if !helper.BindRequest(c, &req) {
		return
	}

	err := models.SaveUser(req)
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.OKResponse(c)
}
