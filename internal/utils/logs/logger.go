// coding: utf-8
// @Author : lryself
// @Date : 2021/12/28 0:03
// @Software: GoLand

package logs

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"sync"
	"time"
)

var (
	log     *logrus.Logger
	logOnce sync.Once
)

func GetLogger() *logrus.Logger {
	logOnce.Do(func() {
		//log = loggerToCmd()
		log = loggerToFile()
		log.Infoln("日志初始化服务完成!")
	})
	return log
}

// 日志记录到文件
func loggerToFile() *logrus.Logger {
	basePath, _ := os.Getwd()
	//logFilePath := basePath + viper.GetString("log.filepath")
	//logFileName := viper.GetString("log.filename")
	logFilePath := path.Join(basePath, "logs")
	logFileName := "system.log"

	// 日志文件
	fileName := path.Join(logFilePath, logFileName)

	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	// 实例化
	logger := logrus.New()

	// 设置输出
	logger.Out = src

	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	logger.AddHook(lfHook)

	return logger
}

//todo 日志记录到 MongoDB
func loggerToMongo() *logrus.Logger {
	return nil
}

//todo 日志记录到 ES
func loggerToES() *logrus.Logger {
	return nil
}

//todo 日志记录到 MQ
func loggerToMQ() *logrus.Logger {
	return nil
}

//记录日志到控制台
func loggerToCmd() *logrus.Logger {
	logger := logrus.New()
	logger.Out = os.Stdout
	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	return logger
}
