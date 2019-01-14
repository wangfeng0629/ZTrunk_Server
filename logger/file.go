package logger

import (
	"ZTrunk_Server/setting"
	"ZTrunk_Server/util"
	"strings"

	"fmt"
	"time"
)

// 文件日志结构
type FileLog struct {
	level         int           // 日志级别
	path          string        // 文件路径
	name          string        // 文件名
	file          *util.File    // 文件句柄
	dataChan      chan *LogData // 数据通道
	splitType     int           // 分割类型
	splitSize     int64         // 分割文件大小
	lastSplitTime int           // 上一次分割时间
}

func CreateFileLog(level int, fileName string) (log Log, err error) {
	var logPath string
	logPath = util.GetSystemGoPATH()
	resultStr := strings.Split(logPath, ";")
	for i := range resultStr {
		fmt.Printf("%s\n", resultStr[i])
		logPath = resultStr[i] + "/src/" + setting.ProjectName
		if isExist, err := util.CheckPathExist(logPath); isExist {
			if err != nil {
				fmt.Printf("get dir error !!!!")
				continue
			}
			if isExist {
				continue
			}
			break
		}
	}
	/*
		logPath, err := util.GetCurrentDir()
		if err != nil {
			fmt.Printf("get current dir failed !!!")
			return log, err
		}
	*/
	chanSize := setting.LogDataChanSize
	if chanSize == 0 {
		fmt.Printf("chanSize error ！！！")
		return log, err
	}
	splitFileSize := setting.SplitFileLogSize
	if splitFileSize == 0 {
		fmt.Printf("splitFileSize error ！！！")
		return log, err
	}
	splitFileType := setting.SplitFileType
	log = &FileLog{
		level:         level,
		path:          logPath + setting.FileDir,
		name:          fileName,
		dataChan:      make(chan *LogData, chanSize),
		splitSize:     splitFileSize,
		splitType:     splitFileType,
		lastSplitTime: time.Now().Hour(),
	}
	log.Init()
	return log, err
}

func (f *FileLog) SetLevel(level int) {
	if level < DEBUG || level > FATAL {
		f.level = DEBUG
	}
	f.level = level
}

func (f *FileLog) Debug(format string, args ...interface{}) {
	if f.level > DEBUG {
		return
	}
	logData := FormatLog(DEBUG, format, args...)
	select {
	case f.dataChan <- logData:
	default:
	}
	FprintfConsoleLog(logData)
}

func (f *FileLog) Trace(format string, args ...interface{}) {
	if f.level > TRACE {
		return
	}
	logData := FormatLog(TRACE, format, args...)
	select {
	case f.dataChan <- logData:
	default:
	}
	FprintfConsoleLog(logData)
}

func (f *FileLog) Warn(format string, args ...interface{}) {
	if f.level > WARN {
		return
	}
	logData := FormatLog(WARN, format, args...)
	select {
	case f.dataChan <- logData:
	default:
	}
	FprintfConsoleLog(logData)
}

func (f *FileLog) Error(format string, args ...interface{}) {
	if f.level > ERROR {
		return
	}
	logData := FormatLog(ERROR, format, args...)
	select {
	case f.dataChan <- logData:
	default:
	}
	FprintfConsoleLog(logData)
}

func (f *FileLog) Fatal(format string, args ...interface{}) {
	if f.level > FATAL {
		return
	}
	logData := FormatLog(FATAL, format, args...)
	select {
	case f.dataChan <- logData:
	default:
	}
	FprintfConsoleLog(logData)
}

func (f *FileLog) Info(format string, args ...interface{}) {
	if f.level > INFO {
		return
	}
	logData := FormatLog(INFO, format, args...)
	select {
	case f.dataChan <- logData:
	default:
	}
	FprintfConsoleLog(logData)
}

func (f *FileLog) Init() {
	filename := fmt.Sprintf("%s/%s.log", f.path, f.name)
	exist, err := util.CheckPathExist(f.path)
	if err != nil {
		fmt.Printf("get dir error !!!!")
		panic(err)
	}
	if exist {
		goto WriteLog
	} else {
		fmt.Printf("dir not exist\n")
		err := util.CreateDir(f.path)
		if err != nil {
			fmt.Printf("mkdir failed !!!, %v", err)
			panic(err)
		}
		goto WriteLog
	}
WriteLog:
	file, err := util.OpenFile(filename)
	if err != nil {
		fmt.Printf("open file %s failed, %v", filename, err)
		panic(err)
	}
	f.file = file
	go f.writeLogBackGround()
}

func (f *FileLog) writeLogBackGround() {
	for logData := range f.dataChan {
		f.fprintfLog(logData)
		f.checkSplitFileByType(f.splitType)
	}
}

func (f *FileLog) fprintfLog(data *LogData) {
	if f.file != nil {
		logStr := FormatNormalLog()
		_, err := fmt.Fprintf(f.file, logStr,
			data.TimeLayout, data.LevelStr, data.FileName, data.LineNumber, data.FuncName, data.Message)
		if err != nil {
			return
		}
	}
}

func (f *FileLog) splitFileByTime() {
	nowTime := time.Now()
	nowHour := nowTime.Hour()

	if nowHour == f.lastSplitTime {
		return
	}
	file := f.file
	f.lastSplitTime = nowHour
	backupFileName := fmt.Sprintf("%s/%s.log_%04d%02d%2d%02d",
		f.path, f.name, nowTime.Year(), nowTime.Month(), nowTime.Day(), f.lastSplitTime)
	fileName := fmt.Sprintf("%s/%s.log", f.path, f.name)
	file.Close()

	util.RenameFile(fileName, backupFileName)
	file, err := util.OpenFile(fileName)
	if err != nil {
		Fatal("open file %s failed, %v", fileName, err)
		panic(err)
	}
}

func (f *FileLog) splitFileBySize() {
	nowTime := time.Now()
	file := f.file

	fileStat, err := file.Stat()
	if err != nil {
		Error("file stat error %v", err)
		return
	}
	fileSize := fileStat.Size()
	if fileSize <= f.splitSize {
		return
	}

	backupFileName := fmt.Sprintf("%s/%s.log_%04d%02d%2d%02d%02d%0d",
		f.path, f.name, nowTime.Year(), nowTime.Month(), nowTime.Day(), nowTime.Hour(), nowTime.Minute(), nowTime.Second())
	fileName := fmt.Sprintf("%s/%s.log", f.path, f.name)
	file.Close()

	util.RenameFile(fileName, backupFileName)
	file, errFile := util.OpenFile(fileName)
	if errFile != nil {
		Fatal("open file %s failed, %v", fileName, errFile)
		panic(errFile)
	}
}

func (f *FileLog) checkSplitFileByType(splitType int) {
	if f.splitType == SplitByTime {
		f.splitFileByTime()
	} else if f.splitType == SplitBySize {
		f.splitFileBySize()
	}
}
