package main

import (
	"ZTrunk_Server/logger"
	"ZTrunk_Server/redispool"
	"github.com/gomodule/redigo/redis"
	//"io/ioutil"
	"net/http"
)

// 获取数据库redis句柄
func getRedisHandle() *redispool.ConnPool {
	return redispool.GetRedis()
}

// get
func HttpGetTask(w http.ResponseWriter, req *http.Request) {
	redis_pool := getRedisHandle()
	req.ParseForm()
	if len(req.PostForm["userid"]) != 1 {
		s_str := "获取数据id值错误"
		logger.Error("%s", s_str)
		w.Write([]byte(s_str))
		return
	}
	//logger.Info("id:%s", req.PostForm["id"][0])
	str, err := redis.String(redis_pool.DoCmd("get", req.PostForm["userid"][0]))
	if err != nil {
		logger.Debug("%s", err)
		w.Write([]byte(""))
		return
	}
	logger.Info("获取数据成功! %s：%s", req.PostForm["userid"][0], str)
	w.Write([]byte(str))
}

// set
func HttpSetTask(w http.ResponseWriter, req *http.Request) {
	redis_pool := getRedisHandle()
	//body, _ := ioutil.ReadAll(req.Body)
	//logger.Debug("body:%s", string(body))
	req.ParseForm()
	if len(req.PostForm["userid"]) != 1 {
		s_str := "更新数据id值错误"
		logger.Error("%s", s_str)
		w.Write([]byte(s_str))
		return
	}
	if len(req.PostForm["userdata"]) != 1 {
		s_str := "更新数据str值错误"
		logger.Error("%s", s_str)
		w.Write([]byte(s_str))
		return
	}
	logger.Info("id:%s：str:%s", req.PostForm["userid"][0], req.PostForm["userdata"][0])
	_, err := redis_pool.DoCmd("set", req.PostForm["userid"][0], req.PostForm["userdata"][0])
	if err != nil {
		logger.Debug("%s", err)
		w.Write([]byte("更新数据出错"))
		return
	}
	logger.Info("更新数据成功! %s：%s", req.PostForm["userid"][0], req.PostForm["userdata"][0])
	w.Write([]byte("更新数据成功！"))
}
