package setting

import (
	"log"

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

	LoggerLevel int
)

func init() {
	var err error
	ConfFile, err = ini.Load("config/config.ini")
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
	MaxOpenConn = sec.Key("MAX_OPEN_CONNS").MustInt(10)
	MaxIdleConn = sec.Key("MAX_IDLE_CONNS").MustInt(2)
}

func LoadLoggerInfo() {
	sec, err := ConfFile.GetSection("log")
	if err != nil {
		log.Fatalf("Fail to get secition 'log': %v", err)
	}
	LoggerLevel = sec.Key("LOGGER_LEVEL").MustInt(0)
}
