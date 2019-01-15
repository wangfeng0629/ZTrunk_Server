package logger

import (
	"ZTrunk_Server/util"

	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

// 日志数据结构
type LogData struct {
	Message    string // 消息体
	TimeLayout string // 时间样式
	LevelStr   string // 级别
	FileName   string // 所属文件
	FuncName   string // 所属函数
	LineNumber int    // 所在行号
}

// 获取日志结构信息
func GetLogDataInfo() (fileName, funcName string, lineNumber int) {
	pc, file, line, ok := runtime.Caller(4)
	if ok {
		fileName = path.Base(file)
		funcName = runtime.FuncForPC(pc).Name()
		lineNumber = line
	}
	return
}

// 日志彩色数码
func color(s string) (color uint8) {
	switch s {
	case "DEBUG":
		color = Green
		break
	case "TRACE":
		color = SkyBlue
		break
	case "INFO":
		color = Blue
		break
	case "WARN":
		color = Yellow
		break
	case "ERROR":
		color = Red
		break
	case "FATAL":
		color = Purple
		break
	default:
		break
	}
	return
}

// 格式化日志
func FormatLog(level int, format string, args ...interface{}) *LogData {
	nowTime := time.Now()
	nowTimeLayout := nowTime.Format("2006-01-02 15:04:05")
	levelStr := getLevelStr(level)
	fileName, funcName, lineNumber := GetLogDataInfo()
	message := fmt.Sprintf(format, args...)
	logData := &LogData{
		message,
		nowTimeLayout,
		levelStr,
		fileName,
		funcName,
		lineNumber,
	}
	return logData
}

// 格式化输出普通日志
func FormatNormalLog() string {
	logStr := "%s [%s] %s:%d %s\n"
	return logStr
}

// 格式化输出颜色日志
func FormatColorLog(data *LogData) string {
	color := color(data.LevelStr)
	conStr := util.IntToString((int)(color))
	colorStr := "\x1b[" + conStr + "m%s\x1b[0m"
	logStr := "%s " + "[" + colorStr + "]" + " [%s:%d] " + colorStr
	return logStr
}

// 格式化输出到控制台
func FprintfConsoleLog(data *LogData) {
	logStr := FormatColorLog(data)
	fmt.Fprintf(os.Stdout, logStr,
		data.TimeLayout, data.LevelStr, data.FileName, data.LineNumber, data.Message)
	fmt.Fprintf(os.Stdout, "\n")
}
