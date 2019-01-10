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

}

func (c *Console) Trace(format string, args ...interface{}) {

}

func (c *Console) Warn(format string, args ...interface{}) {

}

func (c *Console) Error(format string, args ...interface{}) {

}

func (c *Console) Fatal(format string, args ...interface{}) {

}

func (c *Console) Info(format string, args ...interface{}) {

}
