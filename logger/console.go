package logger

type ConsoleLog struct {
	level int
}

func CreateConsoleLog(level int) (log Log, err error) {
	log = &ConsoleLog{
		level: level,
	}
	return log, nil
}

func (c *ConsoleLog) Init() {

}

func (c *ConsoleLog) SetLevel(level int) {
	if level < DEBUG || level > FATAL {
		c.level = DEBUG
	}
	c.level = level
}

func (c *ConsoleLog) Debug(format string, args ...interface{}) {
	if c.level > DEBUG {
		return
	}
	logData := FormatLog(DEBUG, format, args...)
	FprintfConsoleLog(logData)
}

func (c *ConsoleLog) Trace(format string, args ...interface{}) {
	if c.level > TRACE {
		return
	}
	logData := FormatLog(TRACE, format, args...)
	FprintfConsoleLog(logData)
}

func (c *ConsoleLog) Warn(format string, args ...interface{}) {
	if c.level > WARN {
		return
	}
	logData := FormatLog(WARN, format, args...)
	FprintfConsoleLog(logData)
}

func (c *ConsoleLog) Error(format string, args ...interface{}) {
	if c.level > ERROR {
		return
	}
	logData := FormatLog(ERROR, format, args...)
	FprintfConsoleLog(logData)
}

func (c *ConsoleLog) Fatal(format string, args ...interface{}) {
	if c.level > FATAL {
		return
	}
	logData := FormatLog(FATAL, format, args...)
	FprintfConsoleLog(logData)
}

func (c *ConsoleLog) Info(format string, args ...interface{}) {
	if c.level > INFO {
		return
	}
	logData := FormatLog(INFO, format, args...)
	FprintfConsoleLog(logData)
}
