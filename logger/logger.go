package logger

import (
	"ZTrunk_Server/setting"
)

var log Log

func InitLog(name string) (err error) {
	loggerLevel := setting.LoggerLevel
	if name == "console" {
		log, err = CreateConsoleLog(loggerLevel)
	} else {
		log, err = CreateFileLog(loggerLevel, name)
	}
	return
}

func Debug(format string, args ...interface{}) {
	log.Debug(format, args...)
}

func Trace(format string, args ...interface{}) {
	log.Trace(format, args...)
}

func Info(format string, args ...interface{}) {
	log.Info(format, args...)
}

func Warn(format string, args ...interface{}) {
	log.Warn(format, args...)
}

func Error(format string, args ...interface{}) {
	log.Error(format, args...)
}

func Fatal(format string, args ...interface{}) {
	log.Fatal(format, args...)
}
