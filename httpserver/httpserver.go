package main

import (
	"ZTrunk_Server/redispool"
	"log"
)

func main() {

	//start Redis
	if redispool.InitRedis() == false {

		log.Fatal("Redis Server Failed !!!")
		return
	}
	defer redispool.FreePool()
	// start http_server
	HttpStartServer()
}
