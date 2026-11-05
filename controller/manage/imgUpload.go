package manage

import (
	"cms/package/helper"
	"fmt"
	"image"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/gen2brain/webp"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func ImgUpload(c *gin.Context) {
	var webpQuality = 75
	var webpMethod = 6
	var webpExtension = ".webp"
	var imageExtensions = []string{
		".png", ".jpg", ".jpeg",
	}

	file, err := c.FormFile("file")

	if err != nil {
		_file, _err := c.FormFile("upload")

		err = _err
		file = _file
	}

	if err != nil {
		helper.HandleError(c, err, http.StatusBadRequest)
		return
	}

	src, err := file.Open()
	if err != nil {
		helper.HandleError(c, err, http.StatusBadRequest)
		return
	}

	defer src.Close()

	// 拡張子を退避し、一意なIDを生成して保存する
	extension := strings.ToLower(filepath.Ext(file.Filename))

	if slices.Contains(imageExtensions, extension) {
		extension = webpExtension
	}

	guid := xid.New().String()
	path := fmt.Sprintf("/upload/img/%s", guid+extension)
	absPath := fmt.Sprintf("%s%s", ProjectRoot(), path)

	// ※ SVG等他のファイルの場合は圧縮・変換せずに、そのまま保存
	dst, err := os.Create(absPath)
	if err != nil {
		helper.HandleError(c, err, http.StatusBadRequest)
		return
	}

	// JPEGやPNGは、デコードしてWebPに変換
	if extension == webpExtension {
		img, _, err := image.Decode(src)
		if err != nil {
			helper.HandleError(c, err, http.StatusBadRequest)
			return
		}

		// Quality: 80 でWebP変換
		options := webp.Options{
			Quality:  webpQuality,
			Method:   webpMethod,
			Lossless: false,
		}
		if err = webp.Encode(dst, img, options); err != nil {
			helper.HandleError(c, err, http.StatusInternalServerError)
			return
		}
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		helper.HandleError(c, err, http.StatusBadRequest)
		return
	}

	helper.CreatedResponse(c, gin.H{
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
