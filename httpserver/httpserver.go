package main

import (
	"ZTrunk_Server/logger"
	"ZTrunk_Server/redispool"
)

func main() {
	// init log
	if logger.InitLog("console") != nil {
		logger.Fatal("Init Logger Failed !!!")
	}

	//start Redis
	if redispool.InitRedis() == false {
		//log.Fatal("Redis Server Failed !!!")
		logger.Fatal("Redis Server Failed !!!")
		return
	}
	//defer redispool.FreePool()

	// start http_server
	HttpStartServer()
}
