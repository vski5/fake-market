package models

import (
	"fmt"
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
// picName是前端表格某个单元（包含图片的部分）的名字
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
		//timeUnix := time.Now().Unix()
		timeUnix := time.Now().UnixNano() //纳秒级
		//用本日时间戳组成文件名
		userFilmName := strconv.FormatInt(timeUnix, 10) + userFilmExt
		//获取本日时间
		today := time.Now().Format("20060102")

		//拼接文件保存路径
		dateDir := userFilmSrc + today + "/"
		fmt.Println("dateDir-------------", dateDir)
		//创造文件保存路径
		//os.Mkdir(dateDir, 0666) 生成的文件夹有问题（无法操作）。
		//os.MkdirAll(dateFileDir, 0666)生成的文件夹【一样】有问题（无法操作）
		//只能暂时存在focusUpload文件夹这种自己手工创建的问价夹下面，暂时不按日分类了
		//MkdirAllErr := os.Mkdir(dateDir, 0666)  //按当前日期生成文件夹
		//拼接文件保存路径和文件名
		dateFileDir := path.Join(dateDir + userFilmName)
		dateFileDirAndPoint := dateDir + userFilmName
		os.MkdirAll(dateDir, 0777)
		os.Chmod(dateDir, 0777)
		fmt.Println("dateFileDirAndPoint------------------", dateFileDirAndPoint)
		//最重要的，最后一步，保存文件。
		c.SaveUploadedFile(userFilm, dateFileDir)
		return dateFileDir, nil

	}
}

// 因为gin很快，所以只需要获取纳秒级的时间戳，就能分辨同时间上传的多张图片
func UploadManyImg(c *gin.Context, picName string, userFilmSrc string) (string, error) {
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
		today := time.Now().Format("20060102")

		//拼接文件保存路径
		dateDir := userFilmSrc + today + "/"
		fmt.Println("dateDir-------------", dateDir)
		//创造文件保存路径
		//os.Mkdir(dateDir, 0666) 生成的文件夹有问题（无法操作）。
		//os.MkdirAll(dateFileDir, 0666)生成的文件夹【一样】有问题（无法操作）
		//只能暂时存在focusUpload文件夹这种自己手工创建的问价夹下面，暂时不按日分类了
		//MkdirAllErr := os.Mkdir(dateDir, 0666)  //按当前日期生成文件夹
		//拼接文件保存路径和文件名
		dateFileDir := path.Join(dateDir + userFilmName)
		dateFileDirAndPoint := dateDir + userFilmName
		os.MkdirAll(dateDir, 0777)
		os.Chmod(dateDir, 0777)
		fmt.Println("dateFileDirAndPoint------------------", dateFileDirAndPoint)
		//最重要的，最后一步，保存文件。
		c.SaveUploadedFile(userFilm, dateFileDir)
		return dateFileDir, nil

	}
}
