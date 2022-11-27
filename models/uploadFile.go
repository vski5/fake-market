package models

import (
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//获取上传文件的集合

// 获取上传的图片

// 获取上传的单张图片
// userFilmSrc应该为  "./static/focus/" 这种，只包含前置的文件夹不包括名字。。
// 返回 完整的 文件保存路径和文件名
func UploadOneImg(c *gin.Context, picName string, userFilmSrc string) (string, error) {
	userFilm, FormFileError := c.FormFile(picName)
	userFilmExt := path.Ext(userFilm.Filename)
	allowExt := map[string]bool{
		".jpg": true,
		".png": true,
	}
	//allowExt[userFilmExt] 会返回value（也就是对应的布尔类型）
	if _, ok := allowExt[userFilmExt]; !ok {
		//c.String(200, "文件后缀不合法")   ,  判断留给外面
		return "", FormFileError
	} else {
		//获取现在的unix时间戳
		timeUnix := time.Now().Unix()
		//用本日时间戳组成文件名
		userFilmName := strconv.FormatInt(timeUnix, 10) + userFilmExt
		//获取本日时间
		date := time.Now().Format("20060102")

		//拼接文件保存路径
		dateDir := userFilmSrc + "/" + date
		//创造文件保存路径
		MkdirAllErr := os.MkdirAll(dateDir, 0666)
		//拼接文件保存路径和文件名
		dateFileDir := path.Join(dateDir + userFilmName)
		//最重要的，最后一步，保存文件。
		c.SaveUploadedFile(userFilm, dateFileDir)
		return dateFileDir, MkdirAllErr
	}
}
