package models

import (
	"log"
	"os"

	"go.uber.org/zap"
)

// 用zap实现日志库

// 设置输出位置，文件属性
/* func SetupLogger() {
	logFileLocation, _ := os.OpenFile("/home/ubuntu/fake-market/models/zap.go", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	//log.SetOutput()正式的把文件（路径） 设置为日志
	log.SetOutput(logFileLocation) //把日志文件地址传进去
} */

//使用日志
//用log.Printf记录日志 使用%s来记录字符串，%d来记录整数，%f来记录浮点数
//log.Printf("Error fetching url %s : %s " , url, err.Error())

// 先定义logger全局实例
var Logger *zap.Logger

// 设置logger的初始化
func InitLogger(loggerDir string) {
	// 设置输出位置，文件属性
	logFileLocation, _ := os.OpenFile(loggerDir, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	os.Chmod(loggerDir, 0777)
	//log.SetOutput()正式的把文件（路径） 设置为日志
	log.SetOutput(logFileLocation) //把日志文件地址传进去
	Logger, _ = zap.NewProduction()
	//对于SugaredLogger还有一步：
	//SugarLogger = Logger.Sugar() 被称为加点糖
	//后面直接用SugarLogger.Infof("Error", zap.String("url", "http://www.baidu.com"), zap.Error(err) ) 之类的就可以了

	defer Logger.Sync() // flushes buffer, if any

}

/*现在的使用方法
//首先 sugar := models.Logger.Sugar()声明一下
if err != nil {

		sugar.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))

	} else {

		sugar.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))

	}

*/
