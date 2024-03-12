package user

import (
	"cms/db/models"
	code "cms/package/error"
	"cms/package/request"
	"cms/package/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var req request.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		if validErr, ok := err.(response.ValidationError); ok {
			c.JSON(validErr.GetStatus(), validErr.GetResponse())
			return
		}

		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
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

	// タグ追加
	models.SaveUser(user)

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
