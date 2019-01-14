package main

import (
	"ZTrunk_Server/logger"
	"ZTrunk_Server/redispool"
	"ZTrunk_Server/setting"

	"fmt"
	"net/http"
	"strings"

	"github.com/gomodule/redigo/redis"
)

// test 测试代码
func HttpTestTask(w http.ResponseWriter, req *http.Request) {
	h_str := strings.Split(req.URL.RawQuery, "?")
	if len(h_str) == 1 {
		id_str := strings.Split(h_str[0], "=")
		if id_str[0] == "id" {
			HttpGetTask(w, req)
		} else if id_str[0] == "del" {
			HttpDeleteTask(w, req)
		}
	} else if len(h_str) == 2 {
		HttpPostTask(w, req)
	} else {
		logger.Error("http error！")
		//fmt.Println("http error！")
		w.Write([]byte("http error！"))
	}
}

// get
func HttpGetTask(w http.ResponseWriter, req *http.Request) {
	redis_pool := getRedisHandle()
	h_str := strings.Split(req.URL.RawQuery, "?")
	id_str := strings.Split(h_str[0], "=")
	if len(id_str) != 2 {
		logger.Error("http get error!")
		//fmt.Println("http get error!")
		w.Write([]byte("http get error!"))
		return
	}
	id := id_str[1]
	g_str, e := redis.String(redis_pool.DoCmd("get", id))
	if e != nil {
		logger.Error("get error, not this key！")
		//fmt.Println(e)
		w.Write([]byte("get error, not this key！"))
		return
	}
	g_str1 := "get ok！" + g_str
	fmt.Println(g_str1)
	w.Write([]byte(g_str1))
}

// post
func HttpPostTask(w http.ResponseWriter, req *http.Request) {
	redis_pool := getRedisHandle()
	h_str := strings.Split(req.URL.RawQuery, "?")
	id_str := strings.Split(h_str[0], "=")
	v_str := strings.Split(h_str[1], "=")
	if len(id_str) != 2 || len(v_str) != 2 {
		logger.Error("http gost/put error!")
		//fmt.Println("http gost/put error!")
		w.Write([]byte("http post/put error!"))
		return
	}
	_, e := redis_pool.DoCmd("set", id_str[1], v_str[1])
	if e != nil {
		logger.Error("%s", e)
		//fmt.Println(e)
		p_str := "post error " + h_str[0] + "" + h_str[1]
		w.Write([]byte(p_str))
		return
	}
	p_str := "post ok! " + h_str[0] + " " + h_str[1]
	fmt.Println(p_str)
	w.Write([]byte(p_str))
}

// put
func HttpPutTask(w http.ResponseWriter, req *http.Request) {
}

// delete
func HttpDeleteTask(w http.ResponseWriter, req *http.Request) {
	redis_pool := getRedisHandle()
	h_str := strings.Split(req.URL.RawQuery, "?")
	id_str := strings.Split(h_str[0], "=")
	if len(id_str) != 2 {
		logger.Error("http delete error!")
		//fmt.Println("http delete error!")
		w.Write([]byte("http delete error!"))
		return
	}
	id := id_str[1]
	_, e := redis_pool.DoCmd("del", id)
	if e != nil {
		fmt.Println(e)
		return
	}
	p_str := "delete ok! " + id
	fmt.Println(p_str)
	w.Write([]byte(p_str))
}

func HandleMsg(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		HttpTestTask(w, req)
		//HttpGetTask(w, req)
	case "POST":
		HttpPostTask(w, req)
	case "PUT":
		HttpPutTask(w, req)
	case "DELETE":
		HttpDeleteTask(w, req)
	}
}

func getRedisHandle() *redispool.ConnPool {
	return redispool.GetRedis()
}

func HttpStartServer() bool {
	http.HandleFunc("/", HandleMsg)
	http.Handle("/hcg/", http.HandlerFunc(HandleMsg))

	httpAddr := fmt.Sprintf("%s:%d", setting.HTTPIp, setting.HTTPPort)
	e := http.ListenAndServe(httpAddr, nil)
	if e != nil {
		logger.Fatal("%s", e)
		return false
	}
	logger.Info("[启动] Http监听端口 [%d]", setting.HTTPPort)
	return true
}
