package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

type QRcode struct {
	url  string
	size int //size是图像的宽度和高度，单位是像素。
}

func (qr QRcode) Qrcode1(c *gin.Context) {
	var png []byte
	png, err := qrcode.Encode(qr.url, qrcode.Medium, qr.size)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		c.String(200, "返回二维码失败")
	}
	c.String(200, string(png))
}
