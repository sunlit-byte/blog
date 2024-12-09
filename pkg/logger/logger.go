package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sunlit-byte/blog/pkg/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *logrus.Logger

func InitLogger() error {
	// 创建新的实例
	log = logrus.New()

	// setup log output format
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: time.DateTime,
	})

	// set log level
	switch config.LogLevel {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}

	// set log split by week
	log.SetOutput(&lumberjack.Logger{
		Filename:   fmt.Sprintf("./logs/blog-%s.log", time.Now().Format(time.DateOnly)),
		MaxSize:    50,   // max 50 MB
		MaxBackups: 26,   // max 26 week
		MaxAge:     180,  // 最近180天的日志文件
		Compress:   true, // 压缩旧日志文件

	})

	if _, err := os.Stat(config.LogUrl); os.IsNotExist(err) {
		return os.Mkdir(config.LogUrl, os.ModePerm)
	}
	return nil
}

func GetLogger() *logrus.Logger {
	return log
}

// 直接暴露日志方法

func Info(args ...interface{}) {
	log.Info(args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}
