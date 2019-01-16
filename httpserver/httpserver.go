package main

import (
	"ZTrunk_Server/logger"
	"ZTrunk_Server/redispool"
)

func main() {
	// init log
	err := logger.InitLog("HTTPServer")
	if err != nil {
		panic(err)
		return
	}

	// start Redis
	if redispool.InitRedisPool() == false {
		logger.Fatal("Connect Redis Server Failed !!!")
		return
	}
	//defer redispool.FreePool()
	/*
		go func() {
			oneSecTick := time.NewTicker(time.Second)
			for {
				select {
				case <-oneSecTick.C:
					logger.Fatal("oncTick")
				}
			}
		}()
	*/

	// start http_server
	if HttpStartServer() == false {
		logger.Fatal("Http Server Init Failed !!!")
		return
	}
}
