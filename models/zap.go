package models

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 用zap实现日志库

//使用日志
//用log.Printf记录日志 使用%s来记录字符串，%d来记录整数，%f来记录浮点数
//log.Printf("Error fetching url %s : %s " , url, err.Error())

// 先定义logger全局实例
var Logger *zap.Logger
var SugarLogger *zap.SugaredLogger

func InitLogger(path string) {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}

	writeSyncer := zapcore.AddSync(lumberJackLogger)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	SugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

/*
if err != nil {
	SugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
} else {
	SugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
	resp.Body.Close()
}
*/

/*
func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
} */

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
