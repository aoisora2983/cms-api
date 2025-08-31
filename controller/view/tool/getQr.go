package tool

import (
	code "cms/package/error"
	"cms/package/request"
	"cms/package/response"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	"golang.org/x/image/draw"
)

const (
	DOT_TYPE_SQUARE = 1
	DOT_TYPE_CIRCLE = 2

	EXTENSION_JPG = 1
	EXTENSION_PNG = 2

	DEFAULT_PADING = 80
)

/**
 * QRコード作成
 */
func GetQr(c *gin.Context) {
	var req request.GetQrRequest
	if err := c.Bind(&req); err != nil {
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

	// --- エンコード
	qrc, err := qrcode.NewWith(req.Content)
	if err != nil {
		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}

	// --- option設定
	options := []standard.ImageOption{}

	// ------ 画像サイズ設定
	// 必要なサイズから指定のサイズに一番近いサイズを設定(解像度が悪くなるのを防ぐため)
	width := req.QrWidth - DEFAULT_PADING
	dimension := qrc.Dimension()
	nearQRWidth := width / dimension

	options = append(options, standard.WithQRWidth(uint8(nearQRWidth)))

	// ------ フォーマット(jpg pr png)
	// ※ フォーマットがjpgモードでも透過設定があったりロゴ画像がPNGならPNGで保存される
	extension := ".jpg"
	switch req.ImageExtension {
	case EXTENSION_PNG:
		extension = ".png"
		break
	case EXTENSION_JPG:
	default:
		extension = ".jpg"
		break
	}

	// ------ Logo画像設定
	if req.LogoImage != "" {
		logopath := req.LogoImage

		// PNG or JPG
		_extension := strings.ToLower(filepath.Ext(logopath))
		filepath := fmt.Sprintf("%s/user/upload/%s", ProjectRoot(), logopath)

		// ロゴ画像に使用できる画像サイズを算出してリサイズする
		realWidth := (nearQRWidth * dimension) + DEFAULT_PADING
		availableLogoWidth := realWidth / 5

		// ロゴ画像のwidth, height取得
		logoWidth, logoHeight, err := getFileSize(filepath)
		if err != nil {
			response.CustomErrorResponse(
				c,
				http.StatusBadRequest,
				map[string]string{code.SERVER_ERROR: err.Error()},
			)
			return
		}

		// 制限を超えていたらリサイズ
		if logoWidth >= availableLogoWidth || logoHeight >= availableLogoWidth {
			resizeWidth := availableLogoWidth
			resizeHeight := availableLogoWidth
			// 縦横の長さを測って長い方に合わせてリサイズ
			if logoWidth > logoHeight {
				resizeHeight = int(
					float32(logoHeight) *
						float32(float32(availableLogoWidth)/float32(logoWidth)),
				)
			} else {
				resizeWidth = int(
					float32(logoWidth) *
						float32(float32(availableLogoWidth)/float32(logoHeight)),
				)
			}

			ResizeImage(filepath, _extension, resizeWidth, resizeHeight)
		}

		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			response.CustomErrorResponse(
				c,
				http.StatusBadRequest,
				map[string]string{code.SERVER_ERROR: err.Error()},
			)
			return
		}

		if _extension == ".png" {
			extension = ".png"
			options = append(options, standard.WithLogoImageFilePNG(filepath))
		} else {
			options = append(options, standard.WithLogoImageFileJPEG(filepath))
		}
	}

	// ------ ハーフトーン画像設定
	if req.HalftoneImage != "" {
		halftonePath := req.HalftoneImage
		filepath := fmt.Sprintf("%s/user/upload/%s", ProjectRoot(), halftonePath)

		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			response.CustomErrorResponse(
				c,
				http.StatusBadRequest,
				map[string]string{code.SERVER_ERROR: err.Error()},
			)
			return
		}

		options = append(options, standard.WithHalftone(filepath))
	}

	// ------ 透過設定
	if req.BgTransparent {
		extension = ".png"
		options = append(options, standard.WithBgTransparent())
	}

	// ------ 背景色設定
	if !req.BgTransparent && req.BgColor != "" {
		options = append(options, standard.WithBgColorRGBHex(req.BgColor))
	}

	// ------ 前景色設定
	lenFgColor := len(req.FgColor)
	if lenFgColor == 1 {
		options = append(options, standard.WithFgColorRGBHex(req.FgColor[0]))
	} else if lenFgColor > 1 {
		// 複数ならグラデーション設定
		stops := []standard.ColorStop{}

		for index, color := range req.FgColor {
			rgba, err := convertHexToRGB(color)
			if err != nil {
				response.CustomErrorResponse(
					c,
					http.StatusBadRequest,
					map[string]string{code.SERVER_ERROR: err.Error()},
				)
			}

			stops = append(stops, standard.ColorStop{
				Color: rgba, T: float64(index) / float64(lenFgColor-1),
			})
		}

		gradient := standard.NewGradient(float64(req.FgAngle), stops...)

		options = append(options, standard.WithFgGradient(gradient))
	}

	// ------ ドット設定
	if req.DotType == DOT_TYPE_CIRCLE {
		options = append(options, standard.WithCircleShape())
	}

	// --- QRコード作成
	// ------ 画像作成
	guid := xid.New().String()
	filename := fmt.Sprintf("/upload/download/%s", guid+extension)
	filepath := fmt.Sprintf("%s%s", ProjectRoot(), filename)

	// PNGフォーマットの場合
	if extension == ".png" {
		options = append(options, standard.WithBuiltinImageEncoder(standard.PNG_FORMAT))
	}

	writer, err := standard.New(filepath, options...)
	if err != nil {
		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}

	// ------ ファイル出力
	if err = qrc.Save(writer); err != nil {
		response.CustomErrorResponse(
			c,
			http.StatusBadRequest,
			map[string]string{code.SERVER_ERROR: err.Error()},
		)
		return
	}

	// ------ 画像サイズ設定があればリサイズ
	if req.QrWidth > 0 {
		err := ResizeImage(filepath, extension, req.QrWidth, req.QrWidth)
		if err != nil {
			response.CustomErrorResponse(
				c,
				http.StatusBadRequest,
				map[string]string{code.SERVER_ERROR: err.Error()},
			)
			return
		}
	}

	// ダウンロードファイル名を返却
	c.JSON(http.StatusOK, gin.H{
		"url": filename,
	})
}

// #rrggbb(16進数)を分解してRGBAに変換する
func convertHexToRGB(hex string) (color.RGBA, error) {
	rgba := color.RGBA{0, 0, 0, 255}
	switch len(hex) {
	case 7:
		_, err := fmt.Sscanf(hex, "#%02x%02x%02x", &rgba.R, &rgba.G, &rgba.B)
		if err != nil {
			return rgba, err
		}
		break
	case 4:
		_, err := fmt.Sscanf(hex, "#%1x%1x%1x", &rgba.R, &rgba.G, &rgba.B)
		if err != nil {
			return rgba, err
		}
		rgba.R *= 17
		rgba.G *= 17
		rgba.B *= 17
		break
	}
	return rgba, nil
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

func getFileSize(filepath string) (int, int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return 0, 0, err
	}

	bound := img.Bounds()

	return bound.Dx(), bound.Dy(), nil
}

func ResizeImage(filepath string, extension string, width, height int) error {
	originFile, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer originFile.Close()

	var img image.Image
	if extension == ".png" {
		img, err = png.Decode(originFile)
	} else {
		img, err = jpeg.Decode(originFile)
	}

	if err != nil {
		return err
	}

	newImage := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.BiLinear.Scale(newImage, newImage.Bounds(), img, img.Bounds(), draw.Over, nil)

	newFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	if extension == ".png" {
		err = png.Encode(newFile, newImage)
	} else {
		err = jpeg.Encode(newFile, newImage, &jpeg.Options{
			Quality: 80,
		})
	}

	return err
}
