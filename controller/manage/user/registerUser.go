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

	user := map[string]interface{}{
		"id":          req.Id,
		"name":        req.Name,
		"password":    req.Password,
		"mail":        req.Mail,
		"group_id":    req.GroupId,
		"description": req.Description,
		"filename":    req.Filename,
	}

	err := models.SaveUser(user)
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	helper.OKResponse(c)
}
