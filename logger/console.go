package logger

type Console struct {
	level int
}

func CreateConsoleLog(level int) (log Log) {
	log = &Console{
		level: level,
	}
	return log
}

func (c *Console) SetLevel(level int) {
	if level < DEBUG || level > FATAL {
		c.level = DEBUG
	}
	c.level = level
}

func (c *Console) Debug(format string, args ...interface{}) {
	if c.level > DEBUG {
		return
	}
	logData := FormatLog(DEBUG, format, args...)
	FprintfLog(logData)
}

func (c *Console) Trace(format string, args ...interface{}) {
	if c.level > TRACE {
		return
	}
	logData := FormatLog(TRACE, format, args...)
	FprintfLog(logData)
}

func (c *Console) Warn(format string, args ...interface{}) {
	if c.level > WARN {
		return
	}
	logData := FormatLog(WARN, format, args...)
	FprintfLog(logData)
}

func (c *Console) Error(format string, args ...interface{}) {
	if c.level > ERROR {
		return
	}
	logData := FormatLog(ERROR, format, args...)
	FprintfLog(logData)
}

func (c *Console) Fatal(format string, args ...interface{}) {
	if c.level > FATAL {
		return
	}
	logData := FormatLog(FATAL, format, args...)
	FprintfLog(logData)
}

func (c *Console) Info(format string, args ...interface{}) {
	if c.level > INFO {
		return
	}
	logData := FormatLog(INFO, format, args...)
	FprintfLog(logData)
}
