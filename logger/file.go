package logger

type File struct {
	level int
}

func CreateFileLog(level int) (log Log) {
	log = &Console{
		level: level,
	}
	return log
}

func (c *File) SetLevel(level int) {
	if level < DEBUG || level > FATAL {
		c.level = DEBUG
	}
	c.level = level
}

func (c *File) Debug(format string, args ...interface{}) {

}

func (c *File) Trace(format string, args ...interface{}) {

}

func (c *File) Warn(format string, args ...interface{}) {

}

func (c *File) Error(format string, args ...interface{}) {

}

func (c *File) Fatal(format string, args ...interface{}) {

}

func (c *File) Info(format string, args ...interface{}) {

}
