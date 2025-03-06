package utils

import (
	"log"
	"os"
)

const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
)

var (
	logger   = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	logLevel = InfoLevel // 默认日志级别
)

func SetLogLevel(level string) {
	logLevel = level
}

func Debug(v ...interface{}) {
	if logLevel == DebugLevel {
		logger.Println("[DEBUG]", v)
	}
}

func Info(v ...interface{}) {
	logger.Println("[INFO]", v)
}

func Warn(v ...interface{}) {
	logger.Println("[WARN]", v)
}

func Error(v ...interface{}) {
	logger.Println("[ERROR]", v)
}
