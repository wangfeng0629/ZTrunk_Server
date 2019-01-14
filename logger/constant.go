package logger

// 日志级别
const (
	DEBUG = iota
	TRACE
	INFO
	WARN
	ERROR
	FATAL
)

// 日志颜色
const (
	Red = uint8(iota + 91)
	Green
	Yellow
	Blue
	Purple
	SkyBlue
)

// 文件日志分割类型
const (
	SplitByTime = iota
	SplitBySize
)

func getLevelStr(level int) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "UNKNOWN"
}
