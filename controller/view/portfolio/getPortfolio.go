package portfolio

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Portfolio struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	Url         string `json:"url"`
}

func GetPortfolio(c *gin.Context) {
	portfolio := Portfolio{
		Title:       "",
		Description: "",
		Thumbnail:   "",
		Url:         "",
	}
	// if err != nil {
	// 	response.CustomErrorResponse(
	// 		c,
	// 		http.StatusInternalServerError,
	// 		map[string]string{code.SERVER_ERROR: err.Error()},
	// 	)
	// 	return
	// }

	c.JSON(
		http.StatusCreated,
		portfolio,
	)
}
