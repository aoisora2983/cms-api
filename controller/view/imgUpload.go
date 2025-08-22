package view

import (
	code "cms/package/error"
	"cms/package/response"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

var ALLOW_EXTENSION = []string{"jpg", "jpeg", "png"}
var ALLOW_MIME_TYPE = []string{"image/jpeg", "image/png"}

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

	// MIMEタイプ検証
	mimeType := file.Header.Get("Content-Type")
	if err != nil && !ValidMIMEType(mimeType) {
		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.SERVER_ERROR: "許可されたMIMETYPEではありません。"},
		)
		return
	}

	// 拡張子を取得
	extension := strings.ToLower(filepath.Ext(file.Filename))

	// 画像以外の拡張子はNG
	if ValidExtension(extension) {
		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.SERVER_ERROR: "許可された拡張子ではありません。"},
		)
		return
	}

	// 一意なIDを生成
	guid := xid.New().String()

	// ルート外に保存
	filename := guid + extension
	dst, err := os.Create(fmt.Sprintf("%s/user/upload/%s", ProjectRoot(), filename))
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
		"filename": filename,
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

func ValidExtension(extension string) bool {
	for _, allowed := range ALLOW_EXTENSION {
		if extension == allowed {
			return true
		}
	}

	return false
}

func ValidMIMEType(mimeType string) bool {
	for _, allowed := range ALLOW_MIME_TYPE {
		if mimeType == allowed {
			return true
		}
	}

	return false
}
