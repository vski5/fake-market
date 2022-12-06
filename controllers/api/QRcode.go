package api

import (
	"fake-market/models"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

type QRcode struct {
	url  string
	size int //size是图像的宽度和高度，单位是像素。

}
type QRinterface interface {
	URL2Qrcode()
}

func (qr QRcode) URL2Qrcode(c *gin.Context) []byte {
	var png []byte
	png, err := qrcode.Encode(qr.url, qrcode.Medium, qr.size)
	if err != nil {
		models.SugarLogger.Errorf("返回二维码失败 : Error = %s", err.Error())
	}
	return png
}
