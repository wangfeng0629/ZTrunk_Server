package main

import (
	"ZTrunk_Server/redispool"
	"ZTrunk_Server/setting"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
	"reflect"
	"strings"
)

var redisPool = &redispool.ConnPool{}

type Handlers struct {
}

func (h *Handlers) ResAction(w http.ResponseWriter, req *http.Request) {
	fmt.Println("res")
	w.Write([]byte("hcg"))
}

type RedisHander struct {
}

// 获取
func GetHttpTask(w http.ResponseWriter, req *http.Request, c redis.Conn) {
	if c == nil {
		fmt.Println("redis connet failed")
	}
	v, e := c.Do("set", string(111111), req.Body)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(v)
		val, _ := redis.String(c.Do("get", string(111111)))
		fmt.Println(val)
		w.Write([]byte(val))
	}
}

func RedisHttpTask(w http.ResponseWriter, req *http.Request) {
	ret, err := redisPool.SetByString("Redis", "RedisPool")
	if err != nil {
		log.Fatalf("Redis set faial : %v", err)
	} else {
		fmt.Println(ret)
		retVal, _ := redisPool.GetByString("Redis")
		fmt.Println(retVal)
	}
}

// 修改
func PostHttpTask(w http.ResponseWriter, req *http.Request, c redis.Conn) {

}

// 上传
func PutHttpTask(w http.ResponseWriter, req *http.Request, c redis.Conn) {

}

// 删除
func DeleteHttpTask(w http.ResponseWriter, req *http.Request, c redis.Conn) {

}

func say(w http.ResponseWriter, req *http.Request) {
	c := GetRedisHander()
	switch req.Method {
	case "GET":
		GetHttpTask(w, req, c)
		RedisHttpTask(w, req)
	case "POST":
		PostHttpTask(w, req, c)
	case "PUT":
		PutHttpTask(w, req, c)
	case "DELETE":
		DeleteHttpTask(w, req, c)
	}
	defer c.Close()
	return
	fmt.Println("Method:", req.Method)
	pathInfo := strings.Trim(req.URL.Path, "/")
	fmt.Println("pathInfo:", pathInfo)

	parts := strings.Split(pathInfo, "/")
	fmt.Println("parts:", parts)

	var action = "ResAction"
	fmt.Println(strings.Join(parts, "|"))
	if len(parts) > 1 {
		fmt.Println("22222222")
		action = strings.Title(parts[1]) + "Action"
	}
	fmt.Println("action:", action)
	handle := &Handlers{}
	controller := reflect.ValueOf(handle)
	method := controller.MethodByName(action)
	r := reflect.ValueOf(req)
	wr := reflect.ValueOf(w)
	method.Call([]reflect.Value{wr, r})
}

func GetRedisHander() redis.Conn {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("test error")
		return nil
	}
	return c
}

func main() {
	http.HandleFunc("/", say)
	http.Handle("/hcg/", http.HandlerFunc(say))

	redisAddr := fmt.Sprintf("%s:%d", setting.RedisIP, setting.RedisPort)
	fmt.Println(redisAddr)
	redisPool = redispool.InitRedisPool(redisAddr, "", 0, setting.MaxOpenConn, setting.MaxIdleConn)

	httpAddr := fmt.Sprintf("%s:%d", setting.HTTPIp, setting.HTTPPort)
	fmt.Printf("serving on %s:%d \n", setting.HTTPIp, setting.HTTPPort)
	http.ListenAndServe(httpAddr, nil)

	//select {} //阻塞进程
}
