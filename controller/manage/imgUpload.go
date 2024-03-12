package manage

import (
	code "cms/package/error"
	"cms/package/response"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func ImgUpload(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		_file, _err := c.FormFile("upload")

		err = _err
		file = _file
	}

	if err != nil {
		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}

	src, err := file.Open()
	if err != nil {
		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}

	defer src.Close()

	// 拡張子を退避し、一意なIDを生成して保存する
	extension := filepath.Ext(file.Filename)
	guid := xid.New().String()

	path := fmt.Sprintf("/upload/img/%s", guid+extension)
	dst, err := os.Create(fmt.Sprintf("%s/upload/img/%s", ProjectRoot(), guid+extension))
	if err != nil {
		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {

		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"url": path,
	})
}

func ProjectRoot() string {
	currentDir, err := os.Getwd()
	if err != nil {
		return ""
	}

	for {
		_, err := os.ReadFile(filepath.Join(currentDir, "go.mod"))
		if os.IsNotExist(err) {
			if currentDir == filepath.Dir(currentDir) {
				return ""
			}
			currentDir = filepath.Dir(currentDir)
			continue
		} else if err != nil {
			return ""
		}
		break
	}
	return currentDir
}
