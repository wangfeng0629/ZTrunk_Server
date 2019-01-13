package setting

import (
	"flag"
	"log"
	"runtime"

	"github.com/go-ini/ini"
)

var (
	ConfFile *ini.File

	RunMode string

	HTTPIp   string
	HTTPPort int

	RedisIP   string
	RedisPort int

	MaxOpenConn int
	MaxIdleConn int

	LoggerLevel      int
	FileDir          string
	LogDataChanSize  int
	SplitFileLogSize int64
	SplitFileType    int
)

func init() {
	var err error
	var dir string
	goOS := runtime.GOOS
	if goOS == "windows" {
		dir = "config"
		if flag.Lookup("test.v") != nil {
			dir = "../config"
		}
	} else {
		dir = "../config"
	}
	ConfFile, err = ini.Load(dir + "/config.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'config/config.ini': %v", err)
	}

	LoadRunMode()
	LoadServerInfo()
	LoadRedisInfo()
	LoadLoggerInfo()
}

func LoadRunMode() {
	RunMode = ConfFile.Section("").Key("Run_Mode").MustString("debug")
}

func LoadServerInfo() {
	sec, err := ConfFile.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get secition 'server': %v", err)
	}

	HTTPIp = sec.Key("HTTP_IP").MustString("120.92.189.115")
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
}

func LoadRedisInfo() {
	sec, err := ConfFile.GetSection("redis")
	if err != nil {
		log.Fatalf("Fail to get secition 'redis': %v", err)
	}

	RedisIP = sec.Key("HTTP_IP").MustString("127.0.0.1")
	RedisPort = sec.Key("HTTP_PORT").MustInt(6379)
	MaxOpenConn = sec.Key("MAX_OPEN_CONN").MustInt(10)
	MaxIdleConn = sec.Key("MAX_IDLE_CONN").MustInt(2)
}

func LoadLoggerInfo() {
	sec, err := ConfFile.GetSection("log")
	if err != nil {
		log.Fatalf("Fail to get secition 'log': %v", err)
	}
	LoggerLevel = sec.Key("LOGGER_LEVEL").MustInt(0)
	FileDir = sec.Key("FILE_DIR").MustString("/log")
	LogDataChanSize = sec.Key("LOG_DATA_CHAN_SIZE").MustInt(50000)
	SplitFileLogSize = sec.Key("SPLIT_FILE_SIZE").MustInt64(104857600) // 默认100M
	SplitFileType = sec.Key("SPLIT_FILE_TYPE").MustInt(0)
}
